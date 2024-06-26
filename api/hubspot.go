package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/antchfx/jsonquery"
	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/query"
	"github.com/dirkarnez/stemexapi/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/sessions"
	"gorm.io/gorm"
)

// /api/deals/search
func SearchDealIDList(httpClient *http.Client, dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		userName := sessions.Get(ctx).GetString("user_name")
		studentID := ctx.URLParam("student-id")

		studentIDUUID, err := model.ValidUUIDExFromIDString(studentID)
		if err != nil {
			ctx.StopWithStatus(iris.StatusNotFound)
			return
		}

		log.Println(userName, studentID)

		var q = query.Use(dbInstance)

		var studentsToUser *model.StudentToUser
		q.Transaction(func(tx *query.Query) error {
			studentsToUser, err = tx.StudentToUser.Where(tx.StudentToUser.ID.Eq(studentIDUUID)).First()
			return err
		})

		response, err := services.SearchDealIDList(httpClient, studentsToUser.GoogleSheetUserName)
		if err != nil {
			ctx.StopWithStatus(iris.StatusNotFound)
		} else {
			ctx.JSON(response)
		}
	}
}

func GetStudentsToUser(httpClient *http.Client, dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		userName := sessions.Get(ctx).GetString("user_name")

		log.Println(userName)

		var q = query.Use(dbInstance)

		var err error
		var studentsToUser []*model.StudentToUser
		err = q.Transaction(func(tx *query.Query) error {
			studentsToUser, err = tx.StudentToUser.LeftJoin(tx.User, tx.StudentToUser.UserID.EqCol(tx.User.ID)).
				Where(tx.User.Email.Eq(userName)).
				Find()
			return err
		})

		if err != nil {
			ctx.StopWithStatus(iris.StatusForbidden)
		} else {
			// ctx.JSON(bytes)
			ctx.JSON(studentsToUser)
		}
	}
}

func GetDeals(httpClient *http.Client) context.Handler {
	return func(ctx iris.Context) {
		dealId := ctx.URLParam("dealId")
		log.Println("dealId", dealId)

		if len(dealId) < 1 {
			ctx.StopWithStatus(iris.StatusForbidden)
			return
		}

		checkErr := func(err error) {
			if err != nil {
				log.Fatal(err)
			}
		}

		req, err := http.NewRequest("GET", fmt.Sprintf(`https://api.hubapi.com/deals/v1/deal/%s`, dealId), nil)
		checkErr(err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer pat-na1-20d567d6-1d88-4e04-bf49-5c6d78c53c4d")

		resp, err := httpClient.Do(req)
		checkErr(err)
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		checkErr(err)

		ctx.JSON(iris.Map{
			"status": 200,
			"data":   json.RawMessage(body),
		})
	}
}

func GetAttachment(httpClient *http.Client) context.Handler {
	return func(ctx iris.Context) {
		dealId := ctx.URLParam("deal-id")
		log.Println("deal-id", dealId)

		if len(dealId) < 1 {
			ctx.StopWithStatus(iris.StatusForbidden)
			return
		}

		checkErr := func(err error) {
			if err != nil {
				log.Fatal(err)
			}
		}

		req, err := http.NewRequest("GET", fmt.Sprintf(`https://api.hubapi.com/crm/v4/objects/deal/%s/associations/note`, dealId), nil)
		checkErr(err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer pat-na1-20d567d6-1d88-4e04-bf49-5c6d78c53c4d")

		resp, err := httpClient.Do(req)
		checkErr(err)
		defer resp.Body.Close()

		doc, err := jsonquery.Parse(resp.Body)
		checkErr(err)

		objectIDNodes := jsonquery.Find(doc, "/results/*/toObjectId")

		type Attachment struct {
			// ID                string    `json:"id"`
			// CreatedAt         time.Time `json:"createdAt"`
			// UpdatedAt         time.Time `json:"updatedAt"`
			// Archived          bool      `json:"archived"`
			Name string `json:"name"`
			// Path              string    `json:"path"`
			// Size              int       `json:"size"`
			// Type              string    `json:"type"`
			Extension string `json:"extension"`
			// DefaultHostingURL string    `json:"defaultHostingUrl"`
			URL string `json:"url"`
			// IsUsableInContent bool      `json:"isUsableInContent"`
			// Access            string    `json:"access"`
		}

		attachmentList := []Attachment{}

		for _, objectIDNode := range objectIDNodes {
			objectID := fmt.Sprintf("%.0f", objectIDNode.Value())

			reqB, err := http.NewRequest("GET", fmt.Sprintf(`https://api.hubapi.com/crm/v3/objects/notes/%s?properties=hs_attachment_ids`, objectID), nil)
			checkErr(err)
			reqB.Header.Set("Content-Type", "application/json")
			reqB.Header.Set("Authorization", "Bearer pat-na1-20d567d6-1d88-4e04-bf49-5c6d78c53c4d")

			respB, err := httpClient.Do(reqB)
			checkErr(err)
			defer respB.Body.Close()

			docB, err := jsonquery.Parse(respB.Body)
			checkErr(err)

			attachmentIDNode := jsonquery.FindOne(docB, "/properties/hs_attachment_ids")

			if attachmentIDNode != nil {

				attachmentID := fmt.Sprintf("%s", attachmentIDNode.Value())

				reqC, err := http.NewRequest("GET", fmt.Sprintf(`https://api.hubapi.com/files/v3/files/%s`, attachmentID), nil)
				checkErr(err)
				reqC.Header.Set("Content-Type", "application/json")
				reqC.Header.Set("Authorization", "Bearer pat-na1-20d567d6-1d88-4e04-bf49-5c6d78c53c4d")

				respC, err := httpClient.Do(reqC)
				checkErr(err)
				defer respC.Body.Close()

				respCBody, err := io.ReadAll(respC.Body)
				checkErr(err)

				var attachment Attachment
				err = json.Unmarshal(respCBody, &attachment)
				checkErr(err)

				attachmentList = append(attachmentList, attachment)
			}
		}

		ctx.JSON(iris.Map{
			"status": 200,
			"data":   attachmentList,
		})
	}
}
