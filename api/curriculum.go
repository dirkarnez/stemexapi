package api

import (
	"fmt"
	"net/http"

	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/query"
	"github.com/dirkarnez/stemexapi/utils"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetCurriculumTree(dbInstance *gorm.DB) context.Handler {
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

		var q = query.Use(dbInstance)

		var curriculumEntryList []*model.CurriculumEntry
		err := q.Transaction(func(tx *query.Query) error {
			var err error
			curriculumEntryList, err = tx.CurriculumEntry.
				Select(q.CurriculumEntry.ALL, q.CurriculumCourse.ID).
				LeftJoin(q.CurriculumCourse, q.CurriculumEntry.ID.EqCol(q.CurriculumCourse.ID)).
				Where(q.CurriculumEntry.ID.Eq(model.NewUUIDEx())).
				Group(q.CurriculumEntry.ID).
				Find()
			return err
		})

		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		} else {
			ctx.JSON(curriculumEntryList)
		}

		// var err = initSession.
		// 	Select("`ce`.*, CASE WHEN count(`ccytve`.`entry_id`) > 0 OR count(`ccbe`.`entry_id`) > 0 THEN true ELSE false END AS `is_course`").
		// 	Joins("LEFT JOIN `curriculum_course_youtube_video_entries` `ccytve` ON `ccytve`.`entry_id` = `ce`.`id`").
		// 	Joins("LEFT JOIN `curriculum_course_blog_entries` `ccbe` ON `ccbe`.`entry_id` = `ce`.`id`").
		// 	Group("`ce`.`id`").
		// 	Scan(&curriculumEntryList).Error
		// if err != nil {
		// 	ctx.StatusCode(iris.StatusInternalServerError)
		// 	return
		// } else {
		// 	ctx.JSON(curriculumEntryList)
		// }

		// initSession := dbInstance.Table("`curriculum_entries` `ce`")

		// if topLevel {
		// 	initSession = initSession.Where("`parent_id` IS NULL")
		// }

		// var IDUUID model.UUIDEx
		// var err error

		// var q = query.Use(dbInstance)

		// if len(id) != 0 {
		// 	IDUUID, err = model.ValidUUIDExFromIDString(id)
		// 	if err != nil {
		// 		ctx.StopWithStatus(http.StatusNotFound)
		// 		return
		// 	}

		// 	curriculumCourseBlogEntries := []dto.CurriculumCourseBlogEntries{}
		// 	curriculumCourseYoutubeVideoEntries := []dto.CurriculumCourseYoutubeVideoEntries{}

		// 	// err = initSession.Where("`curriculum_entries`.`id` = ?", IDUUID).
		// 	// Joins("left join `curriculum_course_blog_entries` on `curriculum_course_blog_entries`.`entry_id` = `curriculum_entries`.`id`").
		// 	// Joins("left join `curriculum_course_information_entries` on `curriculum_course_information_entries`.`entry_id` = `curriculum_entries`.`id`").
		// 	// Joins("left join `curriculum_course_youtube_video_entries` on `curriculum_course_youtube_video_entries`.`entry_id` = `curriculum_entries`.`id`").
		// 	// First(&details).Error
		// 	// Select("`ce`.*, CASE WHEN count(`entry_id`) > 0 THEN true ELSE false END AS `is_course`").
		// 	//CurriculumCourse
		// 	// err = initSession.
		// 	// 	Select("`ce`.*,  IF(`cc`.`entry_id` IS NOT NULL, true, false) AS `is_course`").
		// 	// 	Joins("LEFT JOIN `curriculum_courses` `cc` ON `cc`.`entry_id` = `ce`.`id`").
		// 	// 	Where("`ce`.`id` = ?", IDUUID).
		// 	// 	Group("`ce`.`id`").
		// 	// 	Limit(1).
		// 	// 	Scan(&curriculumEntry).Error
		// 	// if err != nil {
		// 	// 	ctx.StatusCode(iris.StatusInternalServerError)
		// 	// 	return
		// 	// }

		// 	var curriculumEntry *model.CurriculumEntry = nil
		// 	err := q.Transaction(func(tx *query.Query) error {
		// 		var err error
		// 		curriculumEntry, err = tx.CurriculumEntry.
		// 			Select(q.CurriculumEntry.ALL, q.CurriculumCourse.ID).
		// 			LeftJoin(q.CurriculumCourse, q.CurriculumEntry.ID.EqCol(q.CurriculumCourse.ID)).
		// 			Where(q.CurriculumEntry.ID.Eq(model.NewUUIDEx())).
		// 			Group(q.CurriculumEntry.ID).
		// 			First()

		// 		// .Where(u.Name.Eq("modi")).First()

		// 		// u.WithContext(ctx).Select(u.Name, e.Email).LeftJoin(e, e.UserID.EqCol(u.ID)).Scan(&result)

		// 		// curriculumEntry, err = tx.CurriculumEntry

		// 		// err := u.WithContext(ctx)
		// 		// .Select(u.Name, u.Age.Sum().As("total")).Group(u.Name).Having(u.Name.Eq("group")).Scan(&users)
		// 		// .Where().Find()
		// 		if err != nil {
		// 			if errors.Is(err, gorm.ErrRecordNotFound) {
		// 				curriculumEntry = nil
		// 				return nil
		// 			}
		// 			return err
		// 		}
		// 		return nil
		// 	})

		// 	err = dbInstance.
		// 		Model(&model.CurriculumCourseBlogEntries{}).
		// 		Where(&model.CurriculumCourseBlogEntries{EntryID: &curriculumEntry.ID}).
		// 		Find(&curriculumCourseBlogEntries).Error
		// 	if err != nil {
		// 		ctx.StatusCode(iris.StatusInternalServerError)
		// 		return
		// 	}
		// 	// _ = dbInstance.
		// 	// 	Model(&model.CurriculumCourseInformationEntries{}).
		// 	// 	Where(&model.CurriculumCourseInformationEntries{EntryID: &curriculumEntry.ID}).
		// 	// 	Find(&curriculumCourseInformationEntries).Error

		// 	err = dbInstance.
		// 		Model(&model.CurriculumCourseYoutubeVideoEntries{}).
		// 		Where(&model.CurriculumCourseYoutubeVideoEntries{EntryID: &curriculumEntry.ID}).
		// 		Find(&curriculumCourseYoutubeVideoEntries).Error

		// 	if err != nil {
		// 		ctx.StatusCode(iris.StatusInternalServerError)
		// 		return
		// 	} else {
		// 		ctx.JSON(dto.CurriculumCourseDetails{
		// 			ID:          curriculumEntry.ID,
		// 			Description: curriculumEntry.Description,
		// 			IconID:      curriculumEntry.IconID,
		// 			ParentID:    curriculumEntry.ParentID,
		// 			//Prerequisites: []string
		// 			YoutubeVideoURLs: curriculumCourseYoutubeVideoEntries,
		// 			// InformationEntries: curriculumCourseInformationEntries,
		// 			BlogEntries: curriculumCourseBlogEntries,
		// 		})
		// 	}
		// } else {

		// }
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
			return tx.Table("`curriculum_entries` `ce`").
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

func CreateOrUpdateCurriculumType(s3 *utils.StemexS3Client, dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		var entryToSave = model.CurriculumEntry{}
		type Form struct {
			ID     string `form:"id"  json:"id"`
			IconID string `form:"icon_id"  json:"icon_id"`
			//IconFile/**multipart.FileHeader */ []byte                                           `form:"icon_file"`
			Description string `form:"description"  json:"description"`
			ParentID    string `form:"parent_id"  json:"parent_id"`
		}

		err := dbInstance.Transaction(func(tx *gorm.DB) error {
			// type InformationEntry struct {
			// 	IconID string `form:"icon_id"`
			// 	//IconFile []byte/**multipart.FileHeader*/ `form:"icon_file"`
			// 	Title   string `form:"title"`
			// 	Content string `form:"content"`
			// }

			var form Form
			err := ctx.ReadForm(&form)
			if err != nil {
				return err
			}

			if len(form.ID) > 1 {
				IDUUID, err := model.ValidUUIDExFromIDString(form.ID)
				if err != nil {
					return err
				}
				tx.First(&entryToSave, "`id` = ?", IDUUID)
			}

			entryToSave.Description = form.Description

			if len(form.IconID) > 1 {
				IconIDUUID, err := model.ValidUUIDExFromIDString(form.IconID)
				entryToSave.IconID = &IconIDUUID
				if err != nil {
					return err
				}
			}

			// // Get the max post value size passed via iris.WithPostMaxMemory.
			maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()

			err = ctx.Request().ParseMultipartForm(maxSize)
			if err != nil {
				return err
			}

			_, iconFileHeader, err := ctx.Request().FormFile("icon_file")
			if err == nil {
				file, err := utils.SaveUpload(iconFileHeader, []string{utils.PrefixCourseResourses, entryToSave.Description}, s3, tx, ctx)
				if err != nil {
					return err
				}
				entryToSave.IconID = &file.ID
			}

			if entryToSave.IconID == nil {
				return fmt.Errorf("no icon id")
			}

			if len(form.ParentID) > 1 && form.ParentID != "null" {
				parentIDUUID, err := model.ValidUUIDExFromIDString(form.ParentID)
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

			if err := tx.Save(&entryToSave).Error; err != nil {
				return err
			}

			// // return nil will commit the whole transaction
			return nil
		})

		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
		} else {
			var returnForm Form
			returnForm.Description = entryToSave.Description
			returnForm.ID = entryToSave.ID.ToString()

			if entryToSave.IconID != nil {
				returnForm.IconID = (*entryToSave.IconID).ToString()
			}

			if entryToSave.ParentID != nil {
				returnForm.ParentID = (*entryToSave.ParentID).ToString()
			}

			ctx.JSON(returnForm)
		}
	}
}

func CreateOrUpdateCurriculumEntry(s3 *utils.StemexS3Client, dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		err := dbInstance.Transaction(func(tx *gorm.DB) error {
			// type InformationEntry struct {
			// 	IconID string `form:"icon_id"`
			// 	//IconFile []byte/**multipart.FileHeader*/ `form:"icon_file"`
			// 	Title   string `form:"title"`
			// 	Content string `form:"content"`
			// }

			type Form struct {
				ID     string `form:"id"`
				IconID string `form:"icon_id"`
				//IconFile/**multipart.FileHeader */ []byte                                           `form:"icon_file"`
				Description string `form:"description"`
				ParentID    string `form:"parent_id"`
				// InformationEntries  []InformationEntry                        `form:"information_entries"`
				BlogEntries         []dto.CurriculumCourseBlogEntries         `form:"blog_entries"`
				YoutubeVideoEntries []dto.CurriculumCourseYoutubeVideoEntries `form:"youtube_video_entries"`
			}

			var form Form
			err := ctx.ReadForm(&form)
			if err != nil {
				return err
			}

			var entryToSave = model.CurriculumEntry{}
			entryToSave.Description = form.Description

			if len(form.ID) > 1 {
				IDUUID, err := model.ValidUUIDExFromIDString(form.ID)
				if err != nil {
					return err
				}
				tx.First(&entryToSave, "`id` = ?", IDUUID)
			}

			if len(form.IconID) > 1 {
				IconIDUUID, err := model.ValidUUIDExFromIDString(form.IconID)
				entryToSave.IconID = &IconIDUUID
				if err != nil {
					return err
				}
			}

			// // Get the max post value size passed via iris.WithPostMaxMemory.
			maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()

			err = ctx.Request().ParseMultipartForm(maxSize)
			if err != nil {
				return err
			}

			_, iconFileHeader, err := ctx.Request().FormFile("icon_file")
			if err == nil {
				file, err := utils.SaveUpload(iconFileHeader, []string{utils.PrefixCourseResourses, entryToSave.Description}, s3, tx, ctx)
				if err != nil {
					return err
				}
				entryToSave.IconID = &file.ID
			}

			if entryToSave.IconID == nil {
				return fmt.Errorf("no icon id")
			}

			if len(form.ParentID) > 1 && form.ParentID != "null" {
				parentIDUUID, err := model.ValidUUIDExFromIDString(form.ParentID)
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

			if err := tx.Save(&entryToSave).Error; err != nil {
				return err
			}

			if err := tx.Delete(&model.CurriculumCourseBlogEntries{}, "`entry_id` = ?", entryToSave.ID).Error; err != nil {
				return err
			}

			// if err := tx.Delete(&model.CurriculumCourseInformationEntries{}, "`entry_id` = ?", entryToSave.ID).Error; err != nil {
			// 	return err
			// }

			if err := tx.Delete(&model.CurriculumCourseYoutubeVideoEntries{}, "`entry_id` = ?", entryToSave.ID).Error; err != nil {
				return err
			}

			if form.BlogEntries != nil {
				for _, blogEntry := range form.BlogEntries {
					blogEntryModel := model.CurriculumCourseBlogEntries{}
					blogEntryModel.ID = blogEntry.ID
					blogEntryModel.ExternalURL = blogEntry.ExternalURL
					blogEntryModel.Title = blogEntry.Title
					blogEntryModel.EntryID = &entryToSave.ID

					if err := tx.Clauses(clause.OnConflict{
						Columns:   []clause.Column{{Name: "id"}},
						DoUpdates: clause.AssignmentColumns([]string{"external_url", "title", "entry_id"}),
					}).Create(&blogEntryModel).Error; err != nil {
						return err
					}
				}
			}

			// if form.InformationEntries != nil {
			// 	for i, informationEntry := range form.InformationEntries {
			// 		informationEntryModel := model.CurriculumCourseInformationEntries{}
			// 		informationEntryModel.Title = informationEntry.Title
			// 		informationEntryModel.Content = informationEntry.Content
			// 		// informationEntryModel.EntryID = &entryToSave.ID

			// 		if len(informationEntry.IconID) > 1 {
			// 			IconIDUUID, err := model.ValidUUIDExFromIDString(informationEntry.IconID)
			// 			informationEntryModel.IconID = &IconIDUUID
			// 			if err != nil {
			// 				return err
			// 			}
			// 		}

			// 		_, iconFileHeader, err := ctx.Request().FormFile(fmt.Sprintf("information_entries.%d.icon_file", i))
			// 		if err == nil {
			// 			file, err := utils.SaveUpload(iconFileHeader, []string{utils.PrefixCourseResourses, entryToSave.Description}, s3, tx, ctx)
			// 			if err != nil {
			// 				return err
			// 			}
			// 			informationEntryModel.IconID = &file.ID
			// 		}

			// 		if informationEntryModel.IconID == nil {
			// 			return fmt.Errorf("no icon id")
			// 		}

			// 		if err := tx.Clauses(clause.OnConflict{
			// 			Columns:   []clause.Column{{Name: "id"}},
			// 			DoUpdates: clause.AssignmentColumns([]string{"icon_id", "title", "content", "entry_id"}),
			// 		}).Create(&informationEntryModel).Error; err != nil {
			// 			return err
			// 		}
			// 	}
			// }

			if form.YoutubeVideoEntries != nil {
				for _, youtubeVideoEntry := range form.YoutubeVideoEntries {
					youtubeVideoEntryModel := model.CurriculumCourseYoutubeVideoEntries{}
					youtubeVideoEntryModel.ID = youtubeVideoEntry.ID
					youtubeVideoEntryModel.URL = youtubeVideoEntry.URL
					youtubeVideoEntryModel.EntryID = &entryToSave.ID

					if err := tx.Clauses(clause.OnConflict{
						Columns:   []clause.Column{{Name: "id"}},
						DoUpdates: clause.AssignmentColumns([]string{"url", "title", "entry_id"}),
					}).Create(&youtubeVideoEntryModel).Error; err != nil {
						return err
					}
				}
			}

			// // return nil will commit the whole transaction
			return nil
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

// func ShouldBeACourse(dbInstance *gorm.DB) context.Handler {
// 	return func(ctx iris.Context) {
// 		// SELECT CASE WHEN count(id) > 0 THEN true ELSE false END AS `should_be_a_course` from curriculum_course_information_entries ccie where entry_id in (
// 		// 	select id from curriculum_entries WHERE parent_id = (
// 		// 		SELECT parent_id from curriculum_entries where id = 0x7ba94959764011ee9aa006c3bc34e27e
// 		// 	)
// 		// )

// 		parentID := ctx.URLParam("parent-id")

// 		if len(parentID) < 1 {
// 			ctx.StopWithStatus(http.StatusForbidden)
// 			return
// 		}

// 		parentIDUUIDEx, err := model.ValidUUIDExFromIDString(parentID)
// 		if err != nil {
// 			ctx.StopWithStatus(http.StatusNotFound)
// 			return
// 		}

// 		itShouldBeACourse := false

// 		dbInstance.Table("`curriculum_course_information_entries`").
// 			Select("CASE WHEN count(`id`) > 0 THEN true ELSE false END AS `should_be_a_course`").
// 			Where("`entry_id` IN (?)", dbInstance.Table("`curriculum_entries`").
// 				Select("`id`").
// 				Where("`parent_id` = ?", parentIDUUIDEx)).
// 			Pluck("`should_be_a_course`", &itShouldBeACourse)

// 		ctx.JSON(iris.Map{
// 			"it_should_be_a_course": itShouldBeACourse,
// 		})
// 	}
// }
