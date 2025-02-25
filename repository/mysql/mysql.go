package mysql

import "wire_practice_demo/conf"

// Driver 假装这里是一个mysql的驱动，比如gorm
type Driver struct {
}

var Db *Driver

func Init(conf *conf.DbConf) {
	//	使用conf初始化Driver
	Db = &Driver{}
}
