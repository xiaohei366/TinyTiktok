// 用于保存一些全局的变量和一些配置
package shared

const (
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
	VideoServiceAddr    = ":9021"
	RelationServiceAddr = ":9002"

	ApiServiceName      = "api"
	UserServiceName     = "user"
	VideoServiceName    = "video"
	RelationServiceName = "relation"

	DefaultLimit = 10

	//minio
	MinioUrl       = "localhost"
	MinioPort      = "9000"
	MinioAccessKey = "minio"
	MinioSecretKey = "minio123"

	MinioLocation = "cn-northwest-1"

	//日志的输出路径
	HlogFilePath = "./tmp/hlog/logs/"
	KlogFilePath = "./tmp/klog/logs/"
)
