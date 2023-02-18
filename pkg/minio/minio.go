package minio

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
	"io"
	"net/url"
	"os"
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
	return nil
}

// 上传对象
func UploadObject(minioClient *minio.Client, filetype string,
	bulkName, objectName string, reader io.Reader, size int64, expires time.Duration) (*url.URL, error) {
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
		return nil, err
	}
	klog.Infof("upload %s success", objectName)
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

// localFileName是下载下来的目标地址
func DownObject(minioClient *minio.Client, bulkName, objectName, localFileName string) (*minio.Object, error) {
	object, err := minioClient.GetObject(context.Background(), bulkName, objectName, minio.GetObjectOptions{})
	if err != nil {
		klog.Fatalf("download object error " + err.Error())
		return nil, err
	}
	localFile, err := os.Create(localFileName)
	if err != nil {
		klog.Fatalf("create local object error " + err.Error())
		return nil, err
	}
	_, err = io.Copy(localFile, object)
	if err != nil {
		klog.Fatalf("write object from object to local file error " + err.Error())
		return nil, err
	}
	return object, nil
}

// 删除对象
func RemoveObject(minioClient *minio.Client, bulkName, objectName string) {
	err := minioClient.RemoveObject(context.Background(), bulkName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		klog.Fatalf("remove object eror " + err.Error())
		return
	}
}

// GetFileUrl 从 minio 获取文件Url
func GetFileUrl(minioClient *minio.Client, bucketName string, fileName string, expires time.Duration) (*url.URL, error) {
	ctx := context.Background()
	reqParams := make(url.Values)
	if expires <= 0 {
		expires = time.Second * 60 * 60 * 24
	}
	presignedUrl, err := minioClient.PresignedGetObject(ctx, bucketName, fileName, expires, reqParams)
	if err != nil {
		klog.Errorf("get url of file %s from bucket %s failed, %s", fileName, bucketName, err)
		return nil, err
	}
	// TODO: url可能要做截取
	return presignedUrl, nil
}
