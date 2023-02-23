// 用于保存一些全局的变量和一些配置
package shared

import "time"

const (
	//host_ip
	host_ip = "172.19.134.195"
	//数据库的预留格式
	DBName   = "TinyTiktok"
	MySqlDSN = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	//jwt设置
	SecretKey   = "tiktok6363" //密钥
	IdentityKey = "id"         //jwt携带的识别信息

	//各个服务器的设置
	ExportEndpoint      = ":4317" //链路追踪地址
	ETCDAddress         = "localhost:2379"
	ApiServiceAddr      = ":8888"
	UserServiceAddr     = ":13123"
	VideoServiceAddr    = ":13124"
	RelationServiceAddr = ":13125"
	CommentServiceAddr  = ":13126"
	FavoriteServiceAddr = ":13127"

	ApiServiceName      = "api"
	UserServiceName     = "user"
	VideoServiceName    = "video"
	RelationServiceName = "relation"
	CommentServiceName  = "comment"
	FavoriteServiceName = "favorite"

	DefaultLimit = 10

	//minio
	MinioUrl       = host_ip //这个地方需要用虚拟机ip。不能用localhost，否则客户端不播放视频
	MinioPort      = "9000"
	MinioAccessKey = "minio"
	MinioSecretKey = "minio123"

	MinioLocation = "cn-northwest-1"

	//日志的输出路径
	HlogFilePath = "./tmp/hlog/logs/"
	KlogFilePath = "./tmp/klog/logs/"

	//消息队列
	RabbitMqURI = "amqp://%s:%s@%s:%d/"
	//消息队列设置
	MQUser     = "admin"
	MQPassword = "admin"
	MQHost     = host_ip
	MQPort     = 5672

	//Redis
	RedisExpireTime = time.Hour * 48
	RedisAddr       = host_ip + ":6379"
	RedisPassword   = "123"
	RedisFollower   = 0
	RedisFollow     = 1
	RedisRelation1  = 2
	RedisRelation2  = 3
	RedisName       = 4
	RedisCount1     = 5
	RedisCount2     = 6
)
