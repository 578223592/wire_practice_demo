//go:build wireinject

package factory

import (
	"github.com/google/wire"
	"wire_practice_demo/conf"
	"wire_practice_demo/repository"
	"wire_practice_demo/service"
)

var SuperSet = wire.NewSet(
	conf.NewDbConf,
	repository.NewBookRep,
	repository.NewUserRep,
	repository.NewUserRepPtr,
	//service.NewBookShopServiceErr,
	service.NewBookShopService,
)

func NewBookShopService() *service.BookShop {
	wire.Build(SuperSet)
	return nil
}

//func NewBookShopServiceErr() (*service.BookShop, error) {
//	wire.Build(SuperSet)
//	return nil, nil
//}
