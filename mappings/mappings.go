package mappings

import (
	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/query"
	"github.com/dirkarnez/stemexapi/utils"
)

func MapCurriculumCourseTypeFormToCurriculumEntry(form *dto.CurriculumCourseTypeForm, curriculumEntry *model.CurriculumEntry, s3 *utils.StemexS3Client, txOrQ *query.Query) error {
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

	return nil
}

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

	return nil
}
