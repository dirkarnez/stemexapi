package mappings

import (
	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/query"
	"github.com/dirkarnez/stemexapi/utils"
)

func MapCurriculumCourseFormToCurriculumEntry(form *dto.CurriculumCourseForm, curriculumEntry *model.CurriculumEntry, s3 *utils.StemexS3Client, txOrQ *query.Query) error {
	curriculumEntry.Description = form.Description

	var err error
	curriculumEntry.ID, err = model.ValidUUIDExFromIDString(form.ID)
	if err != nil {
		return err
	}

	curriculumEntry.ParentID, err = model.ValidUUIDExPointerFromIDString(form.ParentID)
	if err != nil {
		return err
	}

	iconIDNilablePtr, err := model.ValidUUIDExPointerFromIDString(form.IconID)
	if err != nil {
		return err
	}

	file, err := utils.SaveUploadV2(form.IconFile, iconIDNilablePtr, []string{utils.PrefixCourseResourses, curriculumEntry.Description}, s3, txOrQ)
	if err != nil {
		return err
	}

	curriculumEntry.IconID = file.ID

	return nil
}
