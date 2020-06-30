package dao

import "github.com/jinzhu/gorm"

type MySqlDB struct {
	db *gorm.DB
}

func NewMysqlDB() *MySqlDB {
	return &MySqlDB{}
}
