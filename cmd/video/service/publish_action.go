package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/disintegration/imaging"
	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/xiaohei366/TinyTiktok/cmd/video/config"
	"github.com/xiaohei366/TinyTiktok/cmd/video/initialize/db"
	dal2 "github.com/xiaohei366/TinyTiktok/cmd/video/service/dal"
	"github.com/xiaohei366/TinyTiktok/kitex_gen/VideoServer"
	"github.com/xiaohei366/TinyTiktok/pkg/minio"
	"image/jpeg"
	"os"
	"strings"
)

type PublishActionService struct { //还是没太明白为什么要new一个videoPost
	ctx context.Context
}

// NewPublishActionService new PublishAction
func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

// PublishAction post video into the minio buckets and database.
func (s *PublishActionService) PublishAction(req *VideoServer.DouyinPublishActionRequest) error {
	klog.Info("Publish action start:")
	// link minio
	minioCli := minio.GetMinioClient()
	klog.Info("minioCli:", minioCli)

	// prepare videos data
	videoData := []byte(req.Data)
	u2, err := uuid.NewV4() //给视频文件加编号
	if err != nil {
		return err
	}

	// prepare videoName
	fileName := u2.String() + "." + "mp4"
	klog.Info("filename:", fileName)
	videoReader := bytes.NewReader(videoData)

	// upload video into minio video bucket and get video playUrl
	url, err := minio.UploadObject(minioCli.Client, "video", config.PublishVideosBucket, fileName, videoReader, int64(len(videoData)), 0) //
	if err != nil {
		return err
	}
	playUrl := strings.Split(url.String(), "?")[0] //做截取
	klog.Info("playUrl:", playUrl)

	// prepare cover name
	u3, err := uuid.NewV4()
	if err != nil {
		return err
	}

	// prepare cover image data
	coverName := u3.String() + "." + "jpg"
	coverImages, err := readFrameAsJpeg(playUrl) //封面数据
	coverReader := bytes.NewReader(coverImages)
	if err != nil {
		return err
	}

	// upload cover image and get coverUrl
	url2, err := minio.UploadObject(minioCli.Client, "image", config.PublishImagesBucket, coverName, coverReader, int64(len(coverImages)), 0)
	coverUrl := strings.Split(url2.String(), "?")[0] // 做截取
	if err != nil {
		return err
	}
	klog.Info("coverUrl:", coverUrl)

	// publish action response model prepare
	videoModel := &db.Video{
		AuthorID: req.UserId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		FavCount: 0,
		ComCount: 0,
		Title:    req.Title,
	}

	// store video info into mysql video model.
	err = dal2.PublishVideo(s.ctx, videoModel)
	if err != nil {
		return err
	}
	return nil
}

// 从视频流中截取一帧并返
// readFrameAsJpeg get a frame as cover image of video
// todo test.
func readFrameAsJpeg(url string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(url).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	klog.Info("ffmpeg input success:")
	img, err := imaging.Decode(reader)
	if err != nil {
		return nil, err
	}
	klog.Info("ffmpeg image Decode success:")
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
