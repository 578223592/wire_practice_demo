package repository

import "wire_practice_demo/repository/mysql"

type BookRep struct {
	db *mysql.Driver
}

func NewBookRep() BookRep {
	return BookRep{db: mysql.Db}
}

func NewBookRepPtr() *BookRep {
	return &BookRep{db: mysql.Db}
}

func (b BookRep) GetBooks() (bookInfo string) {
	bookInfo = "book info"
	return bookInfo
}
