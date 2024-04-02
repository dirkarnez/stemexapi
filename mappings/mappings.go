package mappings

import (
	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/model"
)

func MapCurriculumCourseFormToCurriculumEntry(form *dto.CurriculumCourseForm, entry *model.CurriculumEntry) error {
	var err error
	curriculumEntry.ID, err = model.ValidUUIDExFromIDString(form.ID)
	if err != nil {
		return err
	}

	curriculumEntry.ParentID, err = model.ValidUUIDExPointerFromIDString(form.ParentID)
	if err != nil {
		return err
	}

	curriculumEntry.IconID, err = model.ValidUUIDExFromIDString(form.IconID)
	if err != nil {
		return err
	}

}
