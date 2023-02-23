package config

import "time"

const (
	//数据库
	SqlName         = "admin"
	SqlPassword     = "admin"
	SqlHost         = "localhost"
	SqlPort         = 3306
	FollowTableName = "relation"
	//延迟双删的时间
	SleepTime = time.Millisecond * 500
	//粉丝人数（大V的定义）
	FansNum = 1000
)
