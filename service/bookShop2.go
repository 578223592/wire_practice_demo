package service

import (
	"wire_practice_demo/repository"
)

type BookShop2 struct {
	BookRep repository.BookRep
	UserRep repository.UserRep
}
