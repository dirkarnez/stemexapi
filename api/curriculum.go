package api

import (
	"net/http"

	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/model"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
)

func GetCurriculum(dbInstance *gorm.DB) context.Handler {
	// return func(ctx iris.Context) {
	// 	var curriculumEntryList []model.CurriculumEntry
	// 	if err := dbInstance.Raw(`
	// 		WITH RECURSIVE curriculum_entries_nested AS (
	// 			SELECT id, description, parent_id FROM curriculum_entries WHERE description = 'Micro:bits'
	// 			UNION
	// 			SELECT curriculum_entries.id, curriculum_entries.description, curriculum_entries.parent_id FROM curriculum_entries JOIN curriculum_entries_nested ON curriculum_entries_nested.id = curriculum_entries.parent_id
	// 		)
	// 		SELECT * FROM curriculum_entries_nested
	// 	`).
	// 		Scan(&curriculumEntryList).Error; err != nil {
	// 		ctx.StatusCode(iris.StatusInternalServerError)
	// 		return
	// 	} else {
	// 		ctx.JSON(curriculumEntryList)
	// 	}
	// }

	return func(ctx iris.Context) {
		var curriculumEntryList []dto.CurriculumEntry
		if err := dbInstance.
			Model(&model.CurriculumEntry{}).
			Where("parent_id IS NULL").
			Find(&curriculumEntryList).Error; err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		} else {
			ctx.JSON(curriculumEntryList)
		}
	}
}

func GetCurriculumCourses(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		parentID := ctx.URLParam("parent-id")

		if len(parentID) < 1 {
			ctx.StopWithStatus(http.StatusForbidden)
			return
		}

		parentIDUUID, _ := uuid.Parse(parentID)
		parentIDUUIDEx := model.UUIDEx(parentIDUUID)

		var curriculumEntryList []dto.CurriculumEntry
		if err := dbInstance.
			Model(&model.CurriculumEntry{}).
			Where(&model.CurriculumEntry{ParentID: &parentIDUUIDEx}).
			Find(&curriculumEntryList).Error; err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		} else {
			ctx.JSON(curriculumEntryList)
		}
	}
}

func GetCurriculumCourseDetails(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		id := ctx.URLParam("id")

		if len(id) < 1 {
			ctx.StopWithStatus(http.StatusForbidden)
			return
		}

		idUUID, _ := uuid.Parse(id)

		param := model.CurriculumEntry{}
		param.ID = model.UUIDEx(idUUID)

		entry := model.CurriculumEntry{}
		curriculumCourseBlogEntries := []dto.CurriculumCourseBlogEntries{}

		err := dbInstance.Transaction(func(tx *gorm.DB) error {
			// do some database operations in the transaction (use 'tx' from this point, not 'db')
			if err := tx.
				Model(&model.CurriculumEntry{}).
				Where(&param).
				First(&entry).Error; err != nil {
				// return any error will rollback
				return err
			}

			if err := tx.
				Model(&model.CurriculumCourseBlogEntries{}).
				Where(&model.CurriculumCourseBlogEntries{EntryID: &entry.ID}).
				Find(&curriculumCourseBlogEntries).Error; err != nil {
				// return any error will rollback
				return err
			}

			// return nil will commit the whole transaction
			return nil
		})

		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
		} else {
			ctx.JSON(dto.CurriculumCourseDetails{
				ID:          entry.ID,
				Description: entry.Description,
				//Prerequisites: []string
				// YoutubeVideoURLs: []string
				// InformationEntries []CurriculumCourseInformationEntries
				BlogEntries: curriculumCourseBlogEntries,
			})
		}
	}
}
