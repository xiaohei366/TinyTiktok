//用于保存一些全局的变量和一些配置
package configs

const (

	//jwt设置
	SecretKey       = "tiktok6363" //密钥
	IdentityKey     = "id" //jwt携带的识别信息

	//各个服务器的设置
	ETCDAddress     = ":6363"
	ApiServiceAddr  = ":8888"
	UserServiceAddr = ":9000"

	ApiServiceName      = "api"
	UserServiceName     = "user"

	//数据库
	UserTableName   = "user"
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP             = "tcp"

	
	DefaultLimit    = 10
)