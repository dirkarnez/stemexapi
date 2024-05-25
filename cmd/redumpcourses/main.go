package main

import (
	"fmt"
	"log"

	"github.com/dirkarnez/stemexapi/bo"
	"github.com/dirkarnez/stemexapi/db"
	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/query"
	"github.com/dirkarnez/stemexapi/utils"
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
	dto, err := bo.CreateOrUpdateCurriculumCourse(&dto.CurriculumCourseForm{BlogEntries: []dto.CurriculumCourseBlogEntries{{Title: "你"}}}, utils.NewStemexS3Client(), query.Use(dbInstance))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%+v", dto)
	}
}
