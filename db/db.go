package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConntection() (db *gorm.DB, err error) {
	dsn := "webadmin:password@tcp(ec2-43-198-151-195.ap-east-1.compute.amazonaws.com:3306)/testing?charset=utf8mb4&parseTime=True"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
