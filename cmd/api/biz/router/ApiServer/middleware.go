// Code generated by hertz generator.

package ApiServer

import (
	"github.com/cloudwego/hertz/pkg/app"
	mw "github.com/xiaohei366/TinyTiktok/cmd/api/biz/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _actionMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _comment_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.MiddlewareFuncOptional(),
	}
}

func _commentlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoriteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _action0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favorite_ctionMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _list0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoritelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _feedMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.MiddlewareFuncOptional(),
	}
}

func _feed0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.MiddlewareFuncOptional(),
	}
}

func _action1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publish_ctionMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _list1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishlistMw() []app.HandlerFunc {
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getuserinfoMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		// use jwt mw
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _login0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _register0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relationMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		// use jwt mw
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _action2Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relation_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list2Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relationfollowlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _list3Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relationfollowerlistMw() []app.HandlerFunc {
	return nil
}
