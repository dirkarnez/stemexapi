package main

import (
	"fmt"
	"log"

	"github.com/antchfx/jsonquery"
	"github.com/dirkarnez/stemexapi/db"
	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/query"
	"gorm.io/gen/field"
)

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

	// var user []*model.User
	// q.Transaction(func(tx *query.Query) error {
	// 	var err error
	// 	user, err = tx.User.Where(q.User.Password.Eq("stemex")).Find()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
	// })

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

	jsonquery.LoadURL("https://sheets.googleapis.com/v4/spreadsheets/1mRMBmxKuReGqp9MvcTiv-Z-QcxSDsHUHKwnxPORcj2Y/values/Form?key=AIzaSyBAuyTYKGijZn3jkwoMDlw0ZsR8JR5iOno")

	/*
		{
			"range": "Form!A1:Z1000",
			"majorDimension": "ROWS",
			"values": [
			  [
				"Username",
				"Password",
				"SID",
				"Student Name",
				"Parent Contact Number",
				"Parent Email",
				"Green Highlight = Working Accs",
				"Yellow Highlight = Having Class with us now or in the future",
				"White Highlight = Not Ready Accs Yet"
			  ],
	*/

	var curriculumEntryList []*model.CurriculumEntry
	err := q.Transaction(func(tx *query.Query) error {

		// create a new generic field map to `generic_a`
		f := field.NewField("curriculum_courses", "id")
		// `table_name`.`generic` IS NULL
		//f.IsNotNull()

		var err error
		curriculumEntryList, err = tx.CurriculumEntry.
			Select(q.CurriculumEntry.ALL, f.IsNotNull().As("is_course")).
			LeftJoin(q.CurriculumCourse, q.CurriculumEntry.ID.EqCol(q.CurriculumCourse.ID)).
			Where(q.CurriculumEntry.ID.Eq(model.NewUUIDEx())).
			Group(q.CurriculumEntry.ID).
			Find()
		return err
	})
	fmt.Printf("curriculumEntryList %+v, err = %+v", len(curriculumEntryList), err)

}
