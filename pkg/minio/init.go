package minio

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/xiaohei366/TinyTiktok/cmd/video/config"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
)

type MinioClient struct {
	Client      *minio.Client
	endPoint    string
	VideoBucket string
	ImageBucket string
}

var client MinioClient

func GetMinioClient() MinioClient {
	return client
}

func Init_minio() {
	//初始化配置
	minioEndpoint := shared.MinioUrl + ":" + shared.MinioPort

	minioClient, err := minio.New(minioEndpoint,
		&minio.Options{Creds: credentials.NewStaticV4(shared.MinioAccessKey, shared.MinioSecretKey, "")})
	if err != nil {
		klog.Fatalf("connect minio server fail %s url %s ", err.Error(), minioEndpoint)
		return
	}

	klog.Debug("minio client init successfully")
	client = MinioClient{
		Client:      minioClient,
		endPoint:    minioEndpoint,
		VideoBucket: config.PublishVideosBucket,
		ImageBucket: config.PublishImagesBucket,
	}
	if err := CreateBucket(client.Client, client.VideoBucket); err != nil {
		klog.Errorf("minio client init video bucket failed: %v", err)
	}
	if err := CreateBucket(client.Client, client.ImageBucket); err != nil {
		klog.Errorf("minio client init image bucket failed: %v", err)
	}
}
