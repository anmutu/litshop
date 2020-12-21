package model

import (
	"gorm.io/gorm"
	"litshop/src/config/mysql"
)

type BaseRepository struct {
}

func (*BaseRepository) GetDb(conn string) *gorm.DB {
	db := mysql.GormClientByConn(conn)
	return db
}
