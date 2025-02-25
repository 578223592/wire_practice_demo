package repository

import (
	"wire_practice_demo/repository/mysql"
)

type UserRep struct {
	db *mysql.Driver
}

func NewUserRep() UserRep {
	return UserRep{db: mysql.Db}
}

func NewUserRepPtr() *UserRep {
	return &UserRep{db: mysql.Db}
}

func (u UserRep) GetUserInfo() string {
	return "user Info"
}
