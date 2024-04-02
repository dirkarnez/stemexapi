package mappings

import (
	"mime/multipart"
	"strings"

	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/utils"
)

func MapCurriculumCourseFormToCurriculumEntry(form *dto.CurriculumCourseForm, curriculumEntry *model.CurriculumEntry) error {
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

	OverrideFileID(form.IconFile)

	if err == nil {
		file, err := utils.SaveUploadV2(iconFileHeader, &curriculumEntry.IconID, []string{utils.PrefixCourseResourses, curriculumEntry.Description}, s3, tx, ctx)
		if err != nil {
			return err
		}
		curriculumEntry.IconID = file.ID
	}

	curriculumEntry.IconID, err = model.ValidUUIDExFromIDString(form.IconID)
	if err != nil {
		return err
	}
	return nil
}

func OverrideFileID[V any](file *multipart.FileHeader, pkGetter func(item V) model.UUIDEx) {
	//if file ok, then save the file, override the id
	if file.Size > 0 && len(strings.TrimSpace(file.Filename)) > 0 {

	}
}
