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
		id := ctx.URLParam("id")
		topLevel := ctx.URLParamBoolDefault("top-level", false)

		initSession := dbInstance.Model(&model.CurriculumEntry{})

		if topLevel {
			initSession = initSession.Where("parent_id IS NULL")
		}

		var IDUUID model.UUIDEx
		var err error

		if len(id) != 0 {
			IDUUID, err = model.UUIDExFromIDString(id)
			if err != nil {
				ctx.StopWithStatus(http.StatusNotFound)
				return
			}
			curriculumEntry := model.CurriculumEntry{}
			curriculumCourseBlogEntries := []dto.CurriculumCourseBlogEntries{}
			curriculumCourseInformationEntries := []dto.CurriculumCourseInformationEntries{}
			curriculumCourseYoutubeVideoEntries := []dto.CurriculumCourseYoutubeVideoEntries{}

			// err = initSession.Where("`curriculum_entries`.`id` = ?", IDUUID).
			// Joins("left join `curriculum_course_blog_entries` on `curriculum_course_blog_entries`.`entry_id` = `curriculum_entries`.`id`").
			// Joins("left join `curriculum_course_information_entries` on `curriculum_course_information_entries`.`entry_id` = `curriculum_entries`.`id`").
			// Joins("left join `curriculum_course_youtube_video_entries` on `curriculum_course_youtube_video_entries`.`entry_id` = `curriculum_entries`.`id`").
			// First(&details).Error

			err = initSession.Where("`id` = ?", IDUUID).First(&curriculumEntry).Error

			_ = dbInstance.
				Model(&model.CurriculumCourseBlogEntries{}).
				Where(&model.CurriculumCourseBlogEntries{EntryID: &curriculumEntry.ID}).
				Find(&curriculumCourseBlogEntries).Error

			_ = dbInstance.
				Model(&model.CurriculumCourseInformationEntries{}).
				Where(&model.CurriculumCourseInformationEntries{EntryID: &curriculumEntry.ID}).
				Find(&curriculumCourseInformationEntries).Error

			_ = dbInstance.
				Model(&model.CurriculumCourseYoutubeVideoEntries{}).
				Where(&model.CurriculumCourseYoutubeVideoEntries{EntryID: &curriculumEntry.ID}).
				Find(&curriculumCourseYoutubeVideoEntries).Error

			if err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				return
			} else {
				ctx.JSON(dto.CurriculumCourseDetails{
					ID:          curriculumEntry.ID,
					Description: curriculumEntry.Description,
					IconID:      curriculumEntry.IconID,
					//Prerequisites: []string
					YoutubeVideoURLs:   curriculumCourseYoutubeVideoEntries,
					InformationEntries: curriculumCourseInformationEntries,
					BlogEntries:        curriculumCourseBlogEntries,
				})
			}
		} else {
			var curriculumEntryList []dto.CurriculumEntry
			err = initSession.Find(&curriculumEntryList).Error
			if err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				return
			} else {
				ctx.JSON(curriculumEntryList)
			}
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
