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
	"github.com/xiaohei366/TinyTiktok/pkg/errno"
	"github.com/xiaohei366/TinyTiktok/pkg/minio"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
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
func (s *PublishActionService) PublishAction(req *VideoServer.DouyinPublishActionRequest) error {
	// link minio
	minioCli := minio.GetMinioClient()

	// prepare videos data
	videoData := []byte(req.Data)
	u2, err := uuid.NewV4() //给视频文件加编号
	if err != nil {
		return errno.PublishActionErr
	}

	// prepare videoName
	videoName := u2.String() + "." + "mp4"
	videoReader := bytes.NewReader(videoData)

	// upload video into minio video bucket and get video playUrl
	err = minio.UploadObject(minioCli.Client, "video", config.PublishVideosBucket, videoName, videoReader, int64(len(videoData))) //
	if err != nil {
		return errno.PublishActionErr
	}
	playUrl := "http://" + shared.MinioUrl + ":9000/videos/" + videoName

	// prepare cover name
	u3, err := uuid.NewV4()
	if err != nil {
		return errno.PublishActionErr
	}
	coverName := u3.String() + "." + "jpg"
	coverUrl := "http://" + shared.MinioUrl + ":9000/images/" + coverName
	//开启协程去做将封面上传的事情
	go func() {
		coverImages, err := readFrameAsJpeg(playUrl) //封面数据
		coverReader := bytes.NewReader(coverImages)
		if err != nil {
			klog.Info("Publish action read Frame as jpeg failed")
		}
		err = minio.UploadObject(minioCli.Client, "image", config.PublishImagesBucket, coverName, coverReader, int64(len(coverImages)))
		if err != nil {
			klog.Info("Publish action read Frame as jpeg failed")
		}
	}()

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
		return errno.PublishActionErr
	}
	//wg.Wait()
	return nil
}

// 从视频流中截取一帧并返
// readFrameAsJpeg get a frame as cover image of video
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

	img, err := imaging.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
