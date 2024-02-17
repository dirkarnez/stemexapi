package main

import (
	"fmt"
	"log"

	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/query"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	dsn := "webadmin:password@tcp(ec2-43-198-151-195.ap-east-1.compute.amazonaws.com:3306)/testing?charset=utf8mb4&parseTime=True"
	dbInstance, dbInstanceErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbInstanceErr != nil {
		log.Fatal(dbInstanceErr.Error())
	}
	if dbInstance != nil {
		fmt.Println("Connected!")
	}

	dbInstance = dbInstance.Debug()
	var q = query.Use(dbInstance)

	var user []*model.User
	q.Transaction(func(tx *query.Query) error {
		var err error
		user, err = tx.User.Where(q.User.Password.Eq("stemex")).Find()
		if err != nil {
			return err
		}
		return nil
	})

	fmt.Printf("Users %d", len(user))

	// var roles []model.Role
	// if err := dbInstance.Model(&model.Role{}).Find(&roles).Error; err != nil {
	// 	ctx.StatusCode(iris.StatusInternalServerError)
	// 	return
	// } else {
	// 	ctx.JSON(roles)
	// }

	// // gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	// g.UseDB(dbInstance) // reuse your gorm db

	// // Generate basic type-safe DAO API for struct `model.User` following conventions
	// g.ApplyBasic(model.AllTables...)

	// // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	// g.ApplyInterface(func(Querier) {}, model.AllTables...)

	// // Generate the code
	// g.Execute()
}
