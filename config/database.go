package config

import (
	"fmt"
	"go_backend/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "222.20.126.128"
	port     = 52600
	user     = "root"
	password = "pkg-passwd"
	dbName   = "aclfs"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)
	return db
}
