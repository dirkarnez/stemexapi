package api

import (
	"fmt"
	"net/http"
	"time"

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

		initSession := dbInstance.Table("curriculum_entries `ce`")

		if topLevel {
			initSession = initSession.Where("`parent_id` IS NULL")
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

			err = initSession.
				Select("`ce`.*, CASE WHEN count(`entry_id`) > 0 THEN true ELSE false END AS `is_course`").
				Joins("LEFT JOIN `curriculum_course_information_entries` `ccie` ON `ccie`.`entry_id` = `ce`.`id`").
				Where("`ce`.`id` = ?", IDUUID).
				Group("`ce`.`id`").
				First(&curriculumEntry).Error

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
					ParentID:    curriculumEntry.ParentID,
					//Prerequisites: []string
					YoutubeVideoURLs:   curriculumCourseYoutubeVideoEntries,
					InformationEntries: curriculumCourseInformationEntries,
					BlogEntries:        curriculumCourseBlogEntries,
				})
			}
		} else {
			var curriculumEntryList []dto.CurriculumEntry
			err = initSession.
				Select("`ce`.*, CASE WHEN count(`entry_id`) > 0 THEN true ELSE false END AS `is_course`").
				Joins("LEFT JOIN `curriculum_course_information_entries` `ccie` ON `ccie`.`entry_id` = `ce`.`id`").
				Group("`ce`.`id`").
				Scan(&curriculumEntryList).Error
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

		err := dbInstance.Transaction(func(tx *gorm.DB) error {
			return tx.Table("curriculum_entries `ce`").
				Select("`ce`.*, CASE WHEN count(`entry_id`) > 0 THEN true ELSE false END AS `is_course`").
				Joins("LEFT JOIN `curriculum_course_information_entries` `ccie` ON `ccie`.`entry_id` = `ce`.`id`").
				Where("`ce`.`parent_id` = ?", &parentIDUUIDEx).
				Group("`ce`.`id`").
				Scan(&curriculumEntryList).Error
		})

		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
		} else {
			ctx.JSON(curriculumEntryList)
		}
	}
}

func CreateOrUpdateCurriculumEntry(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		// type CreateOrUpdateCurriculumEntryForm struct {
		// 	ID          string `json:"id`
		// 	Description string `json:"description"`
		// 	IconID      string `json:"icon_id"`
		// 	IconFile    string `json:"icon_file`
		// }

		// type CurriculumEntry struct {
		// 	BaseModel
		// 	IconID         *UUIDEx `gorm:"column:icon_id;type:binary(16)"`
		// 	Icon           *File   `gorm:"foreignKey:IconID"` //constraint:OnDelete:SET NULL
		// 	Description    string  `gorm:"column:description;type:varchar(255);unique;not null"`
		// 	ParentID       *UUIDEx `gorm:"column:parent_id;type:binary(16);uniqueIndex:idx_seq_no_same_level"`
		// 	SeqNoSameLevel uint64  `gorm:"column:seq_no_same_level;not null;default:0;uniqueIndex:idx_seq_no_same_level"`
		// }

		err := dbInstance.Transaction(func(tx *gorm.DB) error {
			var entryToSave = model.CurriculumEntry{}
			IDString := ctx.Request().FormValue("id")
			if len(IDString) > 1 {
				IDUUID, err := model.UUIDExFromIDString(IDString)
				if err != nil {
					return err
				}
				tx.First(&entryToSave, "`id` = ?", IDUUID)
			}

			iconIDString := ctx.Request().FormValue("icon_id")
			if len(iconIDString) > 1 {
				IconIDUUID, err := model.UUIDExFromIDString(iconIDString)
				entryToSave.IconID = &IconIDUUID
				if err != nil {
					return err
				}
			}

			// Get the max post value size passed via iris.WithPostMaxMemory.
			maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()

			err := ctx.Request().ParseMultipartForm(maxSize)
			if err != nil {
				return err
			}

			// Access the uploaded file
			_, header, err := ctx.Request().FormFile("icon_file")
			if err == nil {
				serverPhysicalFileName := fmt.Sprintf("%d", time.Now().UnixNano())
				_, err = ctx.SaveFormFile(header, fmt.Sprintf("./%s/%s", "uploads", serverPhysicalFileName))
				if err != nil {
					ctx.WriteString("Failed to save file: " + header.Filename)
					return err
				}

				file := model.File{OriginalPhysicalFileName: header.Filename, ServerPhysicalFileName: serverPhysicalFileName}
				if err := tx.
					Create(&file).Error; err != nil {
					return err
				}

				entryToSave.IconID = &file.ID
			}

			if entryToSave.IconID == nil {
				ctx.WriteString("No icon uploaded")
				return err
			}

			entryToSave.Description = ctx.Request().FormValue("description")

			parentIDString := ctx.Request().FormValue("parent_id")
			if len(parentIDString) > 1 && parentIDString != "null" {
				parentIDUUID, err := model.UUIDExFromIDString(parentIDString)
				if err != nil {
					return err
				}
				entryToSave.ParentID = &parentIDUUID

				tx.Model(&model.CurriculumEntry{}).
					Select("MAX(`seq_no_same_level`)").
					Where("`parent_id` = ?", *entryToSave.ParentID).
					Group("`parent_id`").
					Scan(&entryToSave.SeqNoSameLevel)
				entryToSave.SeqNoSameLevel = entryToSave.SeqNoSameLevel + 1
			}

			// // return nil will commit the whole transaction
			return tx.Save(&entryToSave).Error
		})

		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
		} else {
			ctx.JSON(iris.Map{
				"status": 200,
			})
		}
	}
}

// func isCurrentLayerForClasses() {
// 	/*
// 		SELECT count(id) from curriculum_course_information_entries ccie where entry_id in
// 		(select id from curriculum_entries ce WHERE parent_id = 0x2B0FE76E764111EE9AA006C3BC34E27E)
// 	*/

// 	db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
// }
