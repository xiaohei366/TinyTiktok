package minio

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/xiaohei366/TinyTiktok/pkg/shared"
	"io"
	"os"
)

// 创建桶
func CreateBucket(ctx context.Context, minioClient *minio.Client, bucketName string) error {
	if len(bucketName) <= 0 {
		klog.Error("Oss bucket name invalid")
	}
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: shared.MinioLocation})
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
func UploadObject(minioClient *minio.Client, filetype string, bulkName, objectName string, reader io.Reader, size int64) (string, error) {
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
		return "", err
	}
	klog.Infof("upload %s success", objectName)
	url := "http://" + client.endPoint + "/" + bulkName + "/" + objectName
	return url, nil
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

// 从OSS获取文件Url//用流式传输的方法?
func GetFileUrl(ctx context.Context, minioClient *minio.Client, bucketName string, fileName string, localFilePath string) ([]byte, error) {
	//下载文件
	body, err := DownObject(minioClient, bucketName, fileName, localFilePath)
	if err != nil {
		klog.Errorf("get obj file %s failed:%v", fileName, err)
		return nil, err
	}
	defer body.Close()
	data, err := io.ReadAll(body)
	if err != nil {
		klog.Errorf("read obj file %s failed:%v", fileName, err)
		return nil, err
	}
	return data, nil
}
