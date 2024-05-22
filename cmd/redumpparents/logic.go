package redumpparents

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

var getString = func(nodeList []*jsonquery.Node, index int) string {
	length := len(nodeList)
	if length > index {
		return strings.TrimSpace(nodeList[index].Value().(string))
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
	parentRole, err := tx.Role.Where(tx.Role.Name.Eq("parent")).First()
	if err != nil {
		return fmt.Errorf("parent role not exist")
	}

	for _, n := range list[1:] {
		n := jsonquery.Find(n, "/*")

		userName := strings.ToLower(getString(n, 0))
		password := strings.ToLower(getString(n, 1))
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

		parentEmail := strings.ToLower(removeAllSpaces(getString(n, 5)))

		if userName == "stemex.demo2023" || userName == "20220872.stemex" {
			continue
		}

		if len(studentName) < 1 || len(parentContactNumber) < 1 || len(parentEmail) < 1 {
			continue
		}

		user, _ := tx.User.Where(tx.User.Email.Eq(parentEmail)).FirstOrInit()
		if user.ID.IsEmpty() {
			user.UserName = parentEmail
			user.Password = "stemex"
			user.ContactNumberAreaCode = areaCode
			user.ContactNumber = parentContactNumber
			user.Email = parentEmail
			user.IsActivated = true //TODO: should be false
			user.IsDummy = false
			user.RoleID = &parentRole.ID

			err = tx.User.Create(user)
			if err != nil {
				return err
			}
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