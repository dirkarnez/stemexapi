package migration

import (
	"fmt"
	"log"

	"github.com/dirkarnez/stemexapi/bo"
	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/query"
	"github.com/dirkarnez/stemexapi/utils"
)

func AddCourseType(qOrTx *query.Query, s3 *utils.StemexS3Client,
	prefix, rootDir, description, iconFilePath string,
) (*dto.CurriculumCourseTypeForm, error) {
	iconFile, err := utils.CreateMultipartFileHeader(fmt.Sprintf(`%s\%s\%s`, prefix, rootDir, iconFilePath))
	if err != nil {
		log.Println("?????????????????????????????")
		log.Fatalln(err)
	}

	dtoInput := dto.CurriculumCourseTypeForm{
		IconFile:    iconFile,
		Description: description,
	}

	return bo.CreateOrUpdateCurriculumCourseType(&dtoInput, s3, qOrTx)
}
