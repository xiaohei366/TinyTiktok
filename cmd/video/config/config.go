package config

const (
	//数据库
	SqlName        = "root"
	SqlPassword    = "123456"
	SqlHost        = "192.168.52.128"
	SqlPort        = 3306
	VideoTableName = "Video"
	Limit          = 30 //限制视频条数

	PublishVideosBucket = "PublishVideosBucket" // minio video bucket name
	PublishImagesBucket = "PublishImagesBucket" // minio image bucket name
)
