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

	var getString = func(nodeList []*jsonquery.Node, index int) string {
		length := len(nodeList)
		if length > index {
			return nodeList[index].Value().(string)
		} else {
			return ""
		}
	}

	var removeAllSpaces = func(str string) string {
		reg, _ := regexp.Compile(`\s+`)
		return reg.ReplaceAllString(str, "")
	}

	// globalObj[0].content.values.slice(1).filter(a => a[0] != "stemex.demo2023" && a[0] != "20220872.stemex")
	var list []*jsonquery.Node
	doc, _ := jsonquery.LoadURL("https://sheets.googleapis.com/v4/spreadsheets/1mRMBmxKuReGqp9MvcTiv-Z-QcxSDsHUHKwnxPORcj2Y/values/Form?key=AIzaSyBAuyTYKGijZn3jkwoMDlw0ZsR8JR5iOno")
	list, _ = jsonquery.QueryAll(doc, "values/*")

	// var user []*model.User
	err := q.Transaction(func(tx *query.Query) error {
		for _, n := range list[1:] {
			n := jsonquery.Find(n, "/*")

			userName := getString(n, 0)
			password := getString(n, 1)
			sid := getString(n, 2)
			studentName := getString(n, 3)
			parentContactNumber := removeAllSpaces(getString(n, 4))

			var areaCode string
			if strings.HasPrefix(parentContactNumber, "+852") {
				parentContactNumber = strings.TrimPrefix(parentContactNumber, "+852")
				areaCode = "852"
			} else if strings.HasPrefix(parentContactNumber, "852") {
				parentContactNumber = strings.TrimPrefix(parentContactNumber, "852")
				areaCode = "852"
			} else {
				areaCode = ""
			}

			parentEmail := removeAllSpaces(getString(n, 5))

			// GreenHigh := n[6].Value().(string)
			// userName := n[7].Value().(string)
			// userName := n[8].Value().(string)

			// := len(n)
			//v := n.Value().(string)
			if userName == "stemex.demo2023" || userName == "20220872.stemex" {
				continue
			}

			parentRole, err := tx.Role.Where(tx.Role.Name.Eq("parent")).First()
			if err != nil {
				return fmt.Errorf("parent role not exist")
			}

			user, _ := tx.User.Where(tx.User.Email.Eq(parentEmail)).FirstOrInit()
			if user.ID.IsEmpty() {
				user.UserName = ""
				user.Password = ""
				user.ContactNumberAreaCode = areaCode
				user.ContactNumber = parentContactNumber
				user.Email = parentEmail
				user.IsActivated = true
				user.IsDummy = false
				user.RoleID = &parentRole.ID
			}

			// create students per record, with associate with new or existing parent user matched by email
			newStudentToUser := model.StudentToUser{GoogleSheetUserName: userName, GoogleSheetPassword: password, GoogleSheetSID: sid, Name: studentName, UserID: user.ID}
			err = tx.StudentToUser.Create(&newStudentToUser)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("err!")
	}
}
