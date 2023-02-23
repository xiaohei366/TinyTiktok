package initialize
import (

	mw "github.com/xiaohei366/TinyTiktok/cmd/api/biz/middleware"
	"github.com/xiaohei366/TinyTiktok/cmd/api/biz/rpc"
	
)

func Init() {
	//RPC框架初始化
	rpc.Init()
	//中间件jwt鉴权
	mw.InitJWT()
	// 日志 初始化
	InitHLogger()
	InitKLogger()
}