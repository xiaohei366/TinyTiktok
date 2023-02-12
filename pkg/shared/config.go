// 用于保存一些全局的变量和一些配置
package shared

const (
	//数据库的预留格式
	DBName        = "TinyTiktok"
	MySqlDSN      = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	MySqlDSN_TEST = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	//jwt设置
	SecretKey   = "tiktok6363" //密钥
	IdentityKey = "id"         //jwt携带的识别信息

	//各个服务器的设置
	ExportEndpoint     = ":4317" //链路追踪地址
	ETCDAddress        = "192.168.52.128:2379"
	ApiServiceAddr     = ":8888"
	UserServiceAddr    = ":9000"
	FeedServiceAddr    = ":9001"
	PublishServiceAddr = ":9002"

	ApiServiceName     = "api"
	ApiFeedServiceName = "api-feed"
	UserServiceName    = "user"
	FeedServiceName    = "feed"
	PublishServiceName = "publish"

	DefaultLimit  = 10
	MinioLocation = "cn-northwest-1"

	//minio
	MinioUrl       = "192.168.52.128"
	MinioPort      = "9000"
	MinioAccessKey = "jacob"
	MinioSecretKey = "jacobminio"
	minioConsolPwd = "jacob"
)
