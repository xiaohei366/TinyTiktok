package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/xiaohei366/TinyTiktok/cmd/feed/service/dal"
	"github.com/xiaohei366/TinyTiktok/cmd/publish/config"
	"github.com/xiaohei366/TinyTiktok/cmd/publish/kitex_gen/PublishServer"
	dal2 "github.com/xiaohei366/TinyTiktok/cmd/publish/service/dal"
	"github.com/xiaohei366/TinyTiktok/pkg/minio"
	"image"
	"image/jpeg"
	"os"
)

type PublishActionService struct { //还是没太明白为什么要new一个videoPost
	ctx context.Context
}

func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

func (s *PublishActionService) PublishAction(req *PublishServer.DouyinPublishActionRequest) (*dal.Video, error) {
	//链接minio
	minioCli := minio.GetMinioClient()
	//创建桶Vdieo的bucket
	err := minio.CreateBucket(s.ctx, minioCli.Client, config.PublishVideosBucket)
	if err != nil {
		return nil, err
	}
	//组合视频
	videoData := []byte(req.Data)
	u2, err := uuid.NewV4() //给视频文件加编号
	if err != nil {
		return nil, err
	}
	fileName := u2.String() + "." + "mp4"
	videoReader := bytes.NewReader(videoData)
	// upload file
	playUrl, err := minio.UploadObject(minioCli.Client, "video", config.PublishVideosBucket, fileName, videoReader, int64(len(videoData))) //
	if err != nil {
		return nil, err
	}
	//获取视频链接
	u3, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	//获取封面
	coverName := u3.String() + "." + "jpg"
	coverImages, err := readFrameAsJpeg(playUrl, coverName) //封面数据
	coverReader := bytes.NewReader(coverImages)
	if err != nil {
		return nil, err
	}
	//上传封面
	coverUrl, err := minio.UploadObject(minioCli.Client, "image", config.PublishImagesBucket, coverName, coverReader, int64(len(coverImages)))
	if err != nil {
		return nil, err
	}
	//存储这个文件
	videoModel := &dal.Video{
		AuthorID: req.User.Id,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		FavCount: 0,
		ComCount: 0,
		Title:    req.Title,
	}
	err = dal2.PublishVideo(s.ctx, videoModel)
	if err != nil {
		return nil, err
	}
	return videoModel, nil
}

// 从视频流中截取一帧并返回
func readFrameAsJpeg(url string, fileName string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(url).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)
	return buf.Bytes(), nil
}
