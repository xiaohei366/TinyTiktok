package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/xiaohei366/TinyTiktok/cmd/video/config"
	"github.com/xiaohei366/TinyTiktok/cmd/video/kitex_gen/VideoServer"
	dal2 "github.com/xiaohei366/TinyTiktok/cmd/video/service/dal"
	"github.com/xiaohei366/TinyTiktok/pkg/minio"
	"image"
	"image/jpeg"
	"os"
)

type PublishActionService struct { //还是没太明白为什么要new一个videoPost
	ctx context.Context
}

// NewPublishActionService new PublishAction
func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

// PublishAction post video into the minio buckets and database.
func (s *PublishActionService) PublishAction(req *VideoServer.DouyinPublishActionRequest) (*dal2.Video, error) {
	// link minio
	minioCli := minio.GetMinioClient()
	// crate video bucket
	err := minio.CreateBucket(s.ctx, minioCli.Client, config.PublishVideosBucket)
	if err != nil {
		return nil, err
	}
	// prepare videos data
	videoData := []byte(req.Data)
	u2, err := uuid.NewV4() //给视频文件加编号
	if err != nil {
		return nil, err
	}
	// prepare videoName
	fileName := u2.String() + "." + "mp4"
	videoReader := bytes.NewReader(videoData)
	// upload video into minio video bucket and get video playUrl
	playUrl, err := minio.UploadObject(minioCli.Client, "video", config.PublishVideosBucket, fileName, videoReader, int64(len(videoData))) //
	if err != nil {
		return nil, err
	}
	// prepare cover name
	u3, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	// prepare cover image data
	coverName := u3.String() + "." + "jpg"
	coverImages, err := readFrameAsJpeg(playUrl) //封面数据
	coverReader := bytes.NewReader(coverImages)
	if err != nil {
		return nil, err
	}
	// upload cover image and get coverUrl
	coverUrl, err := minio.UploadObject(minioCli.Client, "image", config.PublishImagesBucket, coverName, coverReader, int64(len(coverImages)))
	if err != nil {
		return nil, err
	}
	// publish action response model prepare
	videoModel := &dal2.Video{
		AuthorID: req.User.Id,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		FavCount: 0,
		ComCount: 0,
		Title:    req.Title,
	}
	// store video info into mysql video model.
	err = dal2.PublishVideo(s.ctx, videoModel)
	if err != nil {
		return nil, err
	}
	return videoModel, nil
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
