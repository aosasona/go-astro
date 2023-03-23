package repositories

import "xorm.io/xorm"

type user struct {
	db *xorm.Engine
}

func NewUserRepository(db *xorm.Engine) *user {
	return &user{
		db: db,
	}
}
