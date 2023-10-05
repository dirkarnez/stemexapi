package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/antchfx/jsonquery"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

// /api/deals/search
func SearchDeal(httpClient *http.Client) context.Handler {
	return func(ctx iris.Context) {
		studentId := ctx.URLParam("studentId")
		log.Println("studentId", studentId)

		if len(studentId) < 1 {
			ctx.StopWithStatus(http.StatusForbidden)
			return
		}

		type Filter struct {
			PropertyName string `json:"propertyName"`
			Operator     string `json:"operator"`
			Value        string `json:"value"`
		}

		type FilterGroup struct {
			Filters []Filter `json:"filters"`
		}

		type Sort struct {
			PropertyName string `json:"propertyName"`
			Direction    string `json:"direction"`
		}

		type AutoGenerated struct {
			FilterGroups []FilterGroup `json:"filterGroups"`
			Properties   []string      `json:"properties"`
			Sorts        []Sort        `json:"sorts"`
		}

		checkErr := func(err error) {
			if err != nil {
				log.Fatal(err)
			}
		}

		data := AutoGenerated{
			FilterGroups: []FilterGroup{{Filters: []Filter{{PropertyName: "student_id", Operator: "EQ", Value: studentId}}}},
			Properties:   []string{"dealname", "student_id", "new_course_name", "course_dates", "zoom_link"},
			Sorts:        []Sort{{PropertyName: "createdate", Direction: "DESCENDING"}},
		}

		jsonValue, _ := json.Marshal(data)

		req, err := http.NewRequest("POST", `https://api.hubapi.com/crm/v3/objects/deals/search`, bytes.NewBuffer(jsonValue))
		checkErr(err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer pat-na1-20d567d6-1d88-4e04-bf49-5c6d78c53c4d")

		resp, err := httpClient.Do(req)
		checkErr(err)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		checkErr(err)

		ctx.JSON(iris.Map{
			"status": 200,
			"data":   json.RawMessage(body),
		})
	}
}

func GetDeals(httpClient *http.Client) context.Handler {
	return func(ctx iris.Context) {
		dealId := ctx.URLParam("dealId")
		log.Println("dealId", dealId)

		if len(dealId) < 1 {
			ctx.StopWithStatus(http.StatusForbidden)
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

		body, err := ioutil.ReadAll(resp.Body)
		checkErr(err)

		ctx.JSON(iris.Map{
			"status": 200,
			"data":   json.RawMessage(body),
		})
	}
}

func GetAttachment(httpClient *http.Client) context.Handler {
	return func(ctx iris.Context) {
		dealId := ctx.URLParam("dealId")
		log.Println("dealId", dealId)

		if len(dealId) < 1 {
			ctx.StopWithStatus(http.StatusForbidden)
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

				respCBody, err := ioutil.ReadAll(respC.Body)
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
