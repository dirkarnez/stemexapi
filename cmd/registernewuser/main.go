package main

import (
	"database/sql/driver"
	"fmt"
	"log"

	"github.com/dirkarnez/stemexapi/db"
	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/query"
	"github.com/dirkarnez/stemexapi/utils"
	"github.com/samber/lo"
	"gorm.io/gen"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	dbInstance, dbInstanceErr := db.InitConntection()
	if dbInstanceErr != nil {
		log.Fatal(dbInstanceErr.Error())
	}
	if dbInstance != nil {
		fmt.Println("Connected!")
	}

	dbInstance = dbInstance.Debug()
	var q = query.Use(dbInstance)

	q.Transaction(func(tx *query.Query) error {
		var err error

		userName := "stemex"
		password := "password"
		email := "porosil664@artgulin.com"

		newUser := model.User{UserName: userName, Password: password, Email: email, IsActivated: false}
		err = tx.User.Create(&newUser)
		if err != nil {
			return err
		}

		activationKey := model.NewUUIDEx().ToString()
		err = tx.ParentUserActivating.Create(&model.ParentUserActivating{UserID: &newUser.ID, ActivationKey: activationKey})
		if err != nil {
			return err
		}

		err = utils.SendActivationHTMLEmail(newUser.Email, activationKey)
		if err != nil {
			return err
		}

		users, err := tx.User.
			LeftJoin(tx.ParentUserActivating, tx.User.ID.EqCol(tx.ParentUserActivating.UserID)).
			Where(tx.ParentUserActivating.ActivationKey.Eq(activationKey)).Find()
		if err != nil {
			// invalid key
			return err
		}

		tx.User.
			Where(tx.User.ID.In(lo.Map(users, func(user *model.User, index int) driver.Valuer {
				return user.ID
			})...)).
			Update(tx.User.IsActivated, true)

		// tx.User.UpdateFrom(tx.Select(c.ID, c.Address, c.Phone).Where(c.ID.Gt(100))).
		// 	Where(ua.CompanyID.EqCol(ca.ID)).

		//

		return nil
	})

	// var curriculumEntry *model.CurriculumEntry = nil
	// err := q.Transaction(func(tx *query.Query) error {
	// 	var err error
	// 	curriculumEntry, err = tx.CurriculumEntry.
	// 		Select(q.CurriculumEntry.ALL, q.CurriculumCourse.ID).
	// 		LeftJoin(q.CurriculumCourse, q.CurriculumEntry.ID.EqCol(q.CurriculumCourse.ID)).
	// 		Where(q.CurriculumEntry.ID.Eq(model.NewUUIDEx())).
	// 		Group(q.CurriculumEntry.ID).
	// 		First()

	// 	// .Where(u.Name.Eq("modi")).First()

	// 	// u.WithContext(ctx).Select(u.Name, e.Email).LeftJoin(e, e.UserID.EqCol(u.ID)).Scan(&result)

	// 	// curriculumEntry, err = tx.CurriculumEntry

	// 	// err := u.WithContext(ctx)
	// 	// .Select(u.Name, u.Age.Sum().As("total")).Group(u.Name).Having(u.Name.Eq("group")).Scan(&users)
	// 	// .Where().Find()
	// 	if err != nil {
	// 		if errors.Is(err, gorm.ErrRecordNotFound) {
	// 			curriculumEntry = nil
	// 			return nil
	// 		}
	// 		return err
	// 	}
	// 	return nil
	// })

	// fmt.Printf("curriculumEntry %+v, err = %+v", curriculumEntry, err)

	// 	// create a new generic field map to `generic_a`
	// 	f := field.NewField("curriculum_courses", "id")
	// 	// `table_name`.`generic` IS NULL
	// 	//f.IsNotNull()

	// 	var err error
	// 	curriculumEntryList, err = tx.CurriculumEntry.
	// 		Select(q.CurriculumEntry.ALL, f.IsNotNull().As("is_course")).
	// 		LeftJoin(q.CurriculumCourse, q.CurriculumEntry.ID.EqCol(q.CurriculumCourse.ID)).
	// 		Where(q.CurriculumEntry.ID.Eq(model.NewUUIDEx())).
	// 		Group(q.CurriculumEntry.ID).
	// 		Find()
	// 	return err
	// })
	// fmt.Printf("curriculumEntryList %+v, err = %+v", len(curriculumEntryList), err)
	//utils.SendActivationHTMLEmail("noyip90061@aersm.com", "fdgd", "https://stackoverflow.com/")
}

// Select("`ce`.*,  IF(`cc`.`entry_id` IS NOT NULL, true, false) AS `is_course`").
// Joins("LEFT JOIN `curriculum_courses` `cc` ON `cc`.`entry_id` = `ce`.`id`").
// Where("`ce`.`id` = ?", IDUUID).
// Group("`ce`.`id`").
// Limit(1).
