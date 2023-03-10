package minio

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/prometheus/common/log"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
	"io"
	"net/url"
	"time"
)

// 创建桶
func CreateBucket(minioClient *minio.Client, bucketName string) error {
	if len(bucketName) <= 0 {
		klog.Error("Oss bucket name invalid")
	}
	ctx := context.Background()
	err := minioClient.MakeBucket(ctx, bucketName,
		minio.MakeBucketOptions{Region: shared.MinioLocation})
	if err != nil {
		exist, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exist {
			klog.Errorf("bucket %v already exsits\n", bucketName)
			return nil
		} else {
			return err
		}
	} else {
		klog.Infof("bucket create successfully %s:\n", bucketName)
	}
	//设置桶策略
	policy := `{"Version": "2012-10-17",
				"Statement": 
					[{
						"Action":["s3:GetObject"],
						"Effect": "Allow",
						"Principal": {"AWS": ["*"]},
						"Resource": ["arn:aws:s3:::` + bucketName + `/*"],
						"Sid": ""
					}]
				}`
	err = minioClient.SetBucketPolicy(ctx, bucketName, policy)
	if err != nil {
		log.Errorf("SetBucketPolicy %s  err:%s", bucketName, err.Error())
	}
	return nil
}

// 上传对象
func UploadObject(minioClient *minio.Client, filetype string,
	bulkName, objectName string, reader io.Reader, size int64) error {
	var contentType string
	if filetype == "video" {
		contentType = "video/mp4"
	} else if filetype == "image" {
		contentType = "image/jpeg"
	}
	_, err := minioClient.PutObject(context.Background(),
		bulkName,
		objectName,
		reader,
		size,
		minio.PutObjectOptions{
			ContentType: contentType,
		})
	if err != nil {
		klog.Fatalf("upload object error " + err.Error())
		return err
	}
	klog.Infof("upload %s success", objectName)

	return nil
}

func GetMinioUrl(minioClient *minio.Client,
	bulkName, objectName string, expires time.Duration) (*url.URL, error) {
	//返回url
	ctx := context.Background()
	reqParams := make(url.Values)
	if expires <= 0 {
		expires = time.Second * 60 * 60 * 24
	}
	presignedUrl, err := minioClient.PresignedGetObject(ctx, bulkName, objectName, expires, reqParams)
	if err != nil {
		klog.Errorf("get url of file %s from bucket %s failed, %s", objectName, bulkName, err)
		return nil, err
	}
	klog.Info("presignedUrl:", presignedUrl)
	return presignedUrl, nil
}

// 删除对象
func RemoveObject(minioClient *minio.Client, bulkName, objectName string) {
	err := minioClient.RemoveObject(context.Background(), bulkName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		klog.Fatalf("remove object eror " + err.Error())
		return
	}
}
