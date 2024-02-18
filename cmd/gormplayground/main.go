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

	var user *model.CurriculumEntry = nil

	fmt.Printf("Users %d", len(user))
}


Select("`ce`.*,  IF(`cc`.`entry_id` IS NOT NULL, true, false) AS `is_course`").
Joins("LEFT JOIN `curriculum_courses` `cc` ON `cc`.`entry_id` = `ce`.`id`").
Where("`ce`.`id` = ?", IDUUID).
Group("`ce`.`id`").
Limit(1).