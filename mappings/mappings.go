package mappings

import (
	"mime/multipart"
	"strings"

	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/model"
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

	OverrideFileID(curriculumEntry, form.IconFile, func(fileID *model.UUIDEx, entity *model.CurriculumEntry) {
		return entry.IconID = *fileID
	})

	// if err == nil {
	// 	file, err := utils.SaveUploadV2(iconFileHeader, &curriculumEntry.IconID, []string{utils.PrefixCourseResourses, curriculumEntry.Description}, s3, tx, ctx)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	curriculumEntry.IconID = file.ID
	// }

	curriculumEntry.IconID, err = model.ValidUUIDExFromIDString(form.IconID)
	if err != nil {
		return err
	}
	return nil
}

func OverrideFileID[V any](entity *V, file *multipart.FileHeader, onComplete func(*model.File, *V)) {
	//if file ok, then save the file, override the id
	if file.Size > 0 && len(strings.TrimSpace(file.Filename)) > 0 {
		file, err := utils.SaveUploadV2(iconFileHeader, &curriculumEntry.IconID, []string{utils.PrefixCourseResourses, curriculumEntry.Description}, s3, tx, ctx)
		if err != nil {
			return err
		}
		onComplete(&model.File{}, entity)
	}
}
