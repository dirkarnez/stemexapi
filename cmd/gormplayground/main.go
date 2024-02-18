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

	var curriculumEntry *model.CurriculumEntry = nil
	q.Transaction(func(tx *query.Query) error {
		var err error
		curriculumEntry, err = tx.CurriculumEntry.
			Select(q.CurriculumEntry.ALL, q.CurriculumCourse.ID).
			LeftJoin(q.CurriculumEntry, q.CurriculumEntry.ID.EqCol(q.CurriculumCourse.ID)).
			Where(q.CurriculumEntry.ID.Eq(model.NewUUIDEx())).
			Group(q.CurriculumEntry.ID).
			First()

		// .Where(u.Name.Eq("modi")).First()

		// u.WithContext(ctx).Select(u.Name, e.Email).LeftJoin(e, e.UserID.EqCol(u.ID)).Scan(&result)

		// curriculumEntry, err = tx.CurriculumEntry

		// err := u.WithContext(ctx)
		// .Select(u.Name, u.Age.Sum().As("total")).Group(u.Name).Having(u.Name.Eq("group")).Scan(&users)
		// .Where().Find()
		if err != nil {
			return err
		}
		return nil
	})

	fmt.Printf("Users %d", curriculumEntry)
}

// Select("`ce`.*,  IF(`cc`.`entry_id` IS NOT NULL, true, false) AS `is_course`").
// Joins("LEFT JOIN `curriculum_courses` `cc` ON `cc`.`entry_id` = `ce`.`id`").
// Where("`ce`.`id` = ?", IDUUID).
// Group("`ce`.`id`").
// Limit(1).
