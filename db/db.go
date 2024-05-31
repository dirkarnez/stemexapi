package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConntection() (db *gorm.DB, err error) {
	dsn := "webadmin:password@tcp(18.163.71.246:3306)/testing?charset=utf8mb4&parseTime=True"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
