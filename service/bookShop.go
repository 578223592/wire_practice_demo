package service

import (
	"fmt"
	"wire_practice_demo/repository"
)

type BookShopService interface {
	BuyBook()
}

type BookShop struct {
	BookRep repository.BookRep
	UserRep repository.UserRep
}

func (b BookShop) BuyBook() {
	fmt.Println(b.UserRep.GetUserInfo(), "bug book:", b.BookRep.GetBooks())

}

func NewBookShopService(br repository.BookRep, ur repository.UserRep) *BookShop {
	return &BookShop{
		BookRep: br,
		UserRep: ur,
	}
}

func NewBookShopServiceErr(br *repository.BookRep, ur *repository.UserRep) (*BookShop, error) {
	if br == nil || ur == nil {
		return nil, fmt.Errorf("BookShopService init error")
	}
	return &BookShop{
		BookRep: *br,
		UserRep: *ur,
	}, nil
}
