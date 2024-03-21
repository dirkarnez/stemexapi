package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/antchfx/jsonquery"
	"github.com/dirkarnez/stemexapi/db"
	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/query"
	"gorm.io/gen"
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

	var user []*model.User
	q.Transaction(func(tx *query.Query) error {
		var err error

		userName := "stemex"
		password := "password"
		email := "porosil664@artgulin.com"

		newUser := model.User{UserName: userName, Password: password, Email: email, IsActivated: true}
		err = tx.User.Not(gen.Exists(tx.User.Where(tx.User.UserName.Eq(userName)))).Create(&newUser)
		if err != nil {
			return err
		}
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

	var getString = func(nodeList []*jsonquery.Node, index int) string {
		length := len(nodeList)
		if length > index {
			return nodeList[index].Value().(string)
		} else {
			return ""
		}
	}

	// globalObj[0].content.values.slice(1).filter(a => a[0] != "stemex.demo2023" && a[0] != "20220872.stemex")
	var list []*jsonquery.Node
	doc, _ := jsonquery.LoadURL("https://sheets.googleapis.com/v4/spreadsheets/1mRMBmxKuReGqp9MvcTiv-Z-QcxSDsHUHKwnxPORcj2Y/values/Form?key=AIzaSyBAuyTYKGijZn3jkwoMDlw0ZsR8JR5iOno")
	list, _ = jsonquery.QueryAll(doc, "values/*")
	for _, n := range list[1:] {
		n := jsonquery.Find(n, "/*")

		reg, _ := regexp.Compile("s+")

		userName := getString(n, 0)
		password := getString(n, 1)
		sid := getString(n, 2)
		studentName := getString(n, 3)
		parentContactNumber := reg.ReplaceAllString(getString(n, 4), "")

		var areaCode string
		if strings.HasPrefix(parentContactNumber, "+852") || strings.HasPrefix(parentContactNumber, "852") {
			areaCode = "852"
		} else {
			areaCode = ""
		}

		parentEmail := reg.ReplaceAllString(getString(n, 5), "")

		// GreenHigh := n[6].Value().(string)
		// userName := n[7].Value().(string)
		// userName := n[8].Value().(string)

		// := len(n)
		//v := n.Value().(string)
		if userName != "stemex.demo2023" && userName != "20220872.stemex" {
			fmt.Println(userName, password, sid, studentName, parentContactNumber, areaCode, parentEmail)
		}

		// create students per record, with associate with new or existing parent user matched by email
	}

	// var curriculumEntryList []*model.CurriculumEntry
	// err := q.Transaction(func(tx *query.Query) error {

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

}
