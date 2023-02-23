package config

import "time"

const (
	//数据库
	SqlName       = "admin"
	SqlPassword   = "admin"
	SqlHost       = "localhost"
	SqlPort       = 3306
	UserTableName = "user"
	//延迟双删的时间
	SleepTime = time.Millisecond * 600
)
