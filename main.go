package main

import (
	"fmt"
	"wire_practice_demo/conf"
	"wire_practice_demo/factory"
	"wire_practice_demo/repository/mysql"
)

func main() {
	Init()

	service := factory.NewBookShopService()
	fmt.Println(service)
}

func Init() {
	dbConf := conf.NewDbConf()
	mysql.Init(dbConf) // 这样使用单例模式的本质上是init对象，而不是new一个新对象，因此可能wire框架不是那么适用

}
