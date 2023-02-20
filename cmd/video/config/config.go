package config

const (
	//数据库
	SqlName        = "admin"
	SqlPassword    = "admin"
	SqlHost        = "localhost" // 这个是我虚拟机的端口号
	SqlPort        = 3306
	VideoTableName = "Video"

	Limit = 30 //限制视频条数

	PublishVideosBucket = "videos" // minio video bucket name
	PublishImagesBucket = "images" // minio image bucket name
)
