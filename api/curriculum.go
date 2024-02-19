package api

import (
	"fmt"
	"net/http"

	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/query"
	"github.com/dirkarnez/stemexapi/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gen/field"
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
		parentID := ctx.URLParamDefault("parent-id", "")
		// topLevel := ctx.URLParamBoolDefault("top-level", false)

		var err error
		var q = query.Use(dbInstance)

		var parentUUIDPtr *model.UUIDEx = nil
		if len(parentID) > 0 {
			parentUUID, err := model.ValidUUIDExFromIDString(parentID)
			if err != nil {
				ctx.StopWithError(http.StatusNotFound, fmt.Errorf("invalid id"))
				return
			}
			parentUUIDPtr = &parentUUID
		} else {
			parentUUIDPtr = nil
		}

		// type CurriculumEntry struct {
		// 	ID          model.UUIDEx  `json:"id"`
		// 	Description string        `json:"description"`
		// 	ParentID    *model.UUIDEx `json:"parent_id"`
		// 	IconID      *model.UUIDEx `json:"icon_id"`
		// 	IsCourse    bool          `json:"is_course"`
		// }

		// err := u.WithContext(ctx).Select(u.Name, u.ID.Count().As("total")).Group(u.Name).Scan(&users)

		var curriculumEntryList []dto.CurriculumEntry
		err = q.Transaction(func(tx *query.Query) error {
			err := tx.CurriculumEntry.
				Select(q.CurriculumEntry.ALL, field.NewField(q.CurriculumCourse.TableName(), q.CurriculumCourse.ID.ColumnName().String()).IsNotNull().As("is_course")).
				LeftJoin(q.CurriculumCourse, q.CurriculumEntry.ID.EqCol(q.CurriculumCourse.EntryID)).
				Where(func() field.Expr {
					if parentUUIDPtr == nil {
						return q.CurriculumEntry.ParentID.IsNull()
					} else {
						return q.CurriculumEntry.ParentID.Eq(*parentUUIDPtr)
					}
				}()).
				Scan(&curriculumEntryList)
			return err
		})

		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		} else {
			if curriculumEntryList == nil {
				curriculumEntryList = []dto.CurriculumEntry{}
			}
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
func GetCurriculumCourseType(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		ID := ctx.URLParamDefault("id", "")
		// topLevel := ctx.URLParamBoolDefault("top-level", false)

		var err error
		var q = query.Use(dbInstance)

		var idUUIDPtr *model.UUIDEx = nil
		if len(ID) != 0 {
			idUUID, err := model.ValidUUIDExFromIDString(ID)
			if err != nil {
				ctx.StopWithError(http.StatusNotFound, fmt.Errorf("invalid id"))
				return
			}
			idUUIDPtr = &idUUID
		} else {
			idUUIDPtr = nil
		}

		var curriculumEntry *dto.CurriculumEntry
		err = q.Transaction(func(tx *query.Query) error {
			err := tx.CurriculumEntry.
				Select(q.CurriculumEntry.ALL, field.NewField(q.CurriculumCourse.TableName(), q.CurriculumCourse.ID.ColumnName().String()).IsNotNull().As("is_course")).
				LeftJoin(q.CurriculumCourse, q.CurriculumEntry.ID.EqCol(q.CurriculumCourse.ID)).
				Where(q.CurriculumEntry.ID.Eq(*idUUIDPtr)).
				Scan(&curriculumEntry)
			return err
		})

		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		} else {
			ctx.JSON(curriculumEntry)
		}
	}
}

// func GetCurriculumCourses(dbInstance *gorm.DB) context.Handler {
// 	return func(ctx iris.Context) {
// 		parentID := ctx.URLParam("parent-id")

// 		if len(parentID) < 1 {
// 			ctx.StopWithStatus(http.StatusForbidden)
// 			return
// 		}

// 		parentIDUUID, _ := uuid.Parse(parentID)
// 		parentIDUUIDEx := model.UUIDEx(parentIDUUID)

// 		var curriculumEntryList []dto.CurriculumEntry

// 		err := dbInstance.Transaction(func(tx *gorm.DB) error {
// 			return tx.Table("`curriculum_entries` `ce`").
// 				Select("`ce`.*, CASE WHEN count(`entry_id`) > 0 THEN true ELSE false END AS `is_course`").
// 				Joins("LEFT JOIN `curriculum_course_information_entries` `ccie` ON `ccie`.`entry_id` = `ce`.`id`").
// 				Where("`ce`.`parent_id` = ?", &parentIDUUIDEx).
// 				Group("`ce`.`id`").
// 				Scan(&curriculumEntryList).Error
// 		})

// 		if err != nil {
// 			ctx.StatusCode(iris.StatusInternalServerError)
// 		} else {
// 			ctx.JSON(curriculumEntryList)
// 		}
// 	}
// }

func CreateOrUpdateCurriculumCourseType(s3 *utils.StemexS3Client, dbInstance *gorm.DB) context.Handler {
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
				entryToSave.IconID = IconIDUUID
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
				entryToSave.IconID = file.ID
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
			returnForm.IconID = entryToSave.IconID.ToString()

			if entryToSave.ParentID != nil {
				returnForm.ParentID = (*entryToSave.ParentID).ToString()
			}

			ctx.JSON(returnForm)
		}
	}
}

func GetCurriculumCurriculumCourse(s3 *utils.StemexS3Client, dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {

	}
}

func CreateOrUpdateCurriculumCourse(s3 *utils.StemexS3Client, dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		// type InformationEntry struct {
		// 	IconID string `form:"icon_id"`
		// 	//IconFile []byte/**multipart.FileHeader*/ `form:"icon_file"`
		// 	Title   string `form:"title"`
		// 	Content string `form:"content"`
		// }
		//IconFile/**multipart.FileHeader */ []byte                                           `form:"icon_file"`
		// (2) ['curriculum_plan_file', File]
		// course_levels[*].icon_file
		// course_levels[*].Lessons[*].PresentationNotes(id, file_name, file)
		// course_levels[*].Lessons[*].StudentNotes(id, file_name, file)
		// course_levels[*].Lessons[*].TeacherNotes(id, file_name, file)
		// course_levels[*].Lessons[*].MiscMaterials(id, file_name, file)
		// course id
		type Form struct {
			ID                     string                                    `form:"id" json:"id"`
			IconID                 string                                    `form:"icon_id" json:"icon_id"`
			Description            string                                    `form:"description" json:"description"`
			ParentID               string                                    `form:"parent_id" json:"parent_id"`
			CourseID               string                                    `form:"course_id" json:"course_id"`
			CurriculumPlanID       string                                    `form:"curriculum_plan_id" json:"curriculum_plan_id"`
			CurriculumPlanFileName string                                    `form:"curriculum_plan_file_name" json:"curriculum_plan_file_name"` // uploaded
			BlogEntries            []dto.CurriculumCourseBlogEntries         `form:"blog_entries" json:"blog_entries"`
			YoutubeVideoEntries    []dto.CurriculumCourseYoutubeVideoEntries `form:"youtube_video_entries" json:"youtube_video_entries"`
			Levels                 []dto.CurriculumCourseLevels              `form:"levels" json:"levels"`
		}

		var returnForm Form
		var form Form
		err := ctx.ReadForm(&form)
		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
			return
		}

		var q = query.Use(dbInstance)
		var curriculumEntry model.CurriculumEntry = model.CurriculumEntry{}
		err = q.Transaction(func(tx *query.Query) error {
			if len(form.ID) > 1 {
				IDUUID, err := model.ValidUUIDExFromIDString(form.ID)
				if err != nil {
					return err
				}
				curriculumEntry.ID = IDUUID
			}

			curriculumEntry.Description = form.Description

			if len(form.ParentID) > 1 {
				ParentIDUUID, err := model.ValidUUIDExFromIDString(form.ParentID)
				if err != nil {
					return err
				}
				curriculumEntry.ParentID = &ParentIDUUID
			}

			if len(form.IconID) > 1 {
				IconIDUUID, err := model.ValidUUIDExFromIDString(form.IconID)
				if err != nil {
					return err
				}
				curriculumEntry.IconID = IconIDUUID
			}

			// // Get the max post value size passed via iris.WithPostMaxMemory.
			maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()

			err = ctx.Request().ParseMultipartForm(maxSize)
			if err != nil {
				return err
			}

			_, iconFileHeader, err := ctx.Request().FormFile("icon_file")
			if err == nil {
				file, err := utils.SaveUploadV2(iconFileHeader, &curriculumEntry.IconID, []string{utils.PrefixCourseResourses, curriculumEntry.Description}, s3, tx, ctx)
				if err != nil {
					return err
				}
				curriculumEntry.IconID = file.ID
			}

			err = tx.CurriculumEntry.Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(&curriculumEntry)
			if err != nil {
				return err
			}

			/* associations: CurriculumCourse */
			var curriculumCourse = model.CurriculumCourse{}
			if len(form.CourseID) > 1 {
				CourseIDUUID, err := model.ValidUUIDExFromIDString(form.CourseID)
				if err != nil {
					return err
				}
				curriculumCourse.ID = CourseIDUUID
			}

			if len(form.CurriculumPlanID) > 1 {
				curriculumPlanIDUUID, err := model.ValidUUIDExFromIDString(form.CurriculumPlanID)
				if err != nil {
					return err
				}
				curriculumCourse.CurriculumPlanID = curriculumPlanIDUUID
			}

			_, curriculumPlanFileHeader, err := ctx.Request().FormFile("curriculum_plan_file")
			if err == nil {
				file, err := utils.SaveUploadV2(curriculumPlanFileHeader, &curriculumCourse.CurriculumPlanID, []string{utils.PrefixCourseResourses, curriculumEntry.Description}, s3, tx, ctx)
				if err != nil {
					return err
				}
				curriculumCourse.CurriculumPlanID = file.ID
			}

			curriculumCourse.EntryID = curriculumEntry.ID

			err = tx.CurriculumCourse.Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(&curriculumCourse)
			if err != nil {
				return err
			}

			returnForm.CourseID = curriculumCourse.ID.ToString()

			/* associations: CurriculumCourseBlogEntries*/
			var blogs []*model.CurriculumCourseBlogEntries
			for _, dto := range form.BlogEntries {
				entity := model.CurriculumCourseBlogEntries{}

				if len(dto.ID) > 1 {
					IDUUID, err := model.ValidUUIDExFromIDString(dto.ID)
					if err != nil {
						return err
					}
					entity.ID = IDUUID
				}
				entity.ExternalURL = dto.ExternalURL
				entity.Title = dto.Title
				entity.EntryID = &curriculumEntry.ID
				blogs = append(blogs, &entity)
			}

			err = tx.CurriculumCourseBlogEntries.Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(blogs...)
			if err != nil {
				return err
			}

			for _, saved := range blogs {
				returnForm.BlogEntries = append(returnForm.BlogEntries, dto.CurriculumCourseBlogEntries{
					ExternalURL: saved.ExternalURL,
					Title:       saved.Title,
				})
			}

			/* associations: CurriculumCourseYoutubeVideoEntries*/
			var youtubes []*model.CurriculumCourseYoutubeVideoEntries
			for _, dto := range form.YoutubeVideoEntries {
				entity := model.CurriculumCourseYoutubeVideoEntries{}

				if len(dto.ID) > 1 {
					IDUUID, err := model.ValidUUIDExFromIDString(dto.ID)
					if err != nil {
						return err
					}
					entity.ID = IDUUID
				}
				entity.URL = dto.URL
				entity.EntryID = &curriculumEntry.ID
				youtubes = append(youtubes, &entity)
			}

			err = tx.CurriculumCourseYoutubeVideoEntries.Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(youtubes...)
			if err != nil {
				return err
			}

			for _, saved := range youtubes {
				returnForm.YoutubeVideoEntries = append(returnForm.YoutubeVideoEntries, dto.CurriculumCourseYoutubeVideoEntries{
					URL: saved.URL,
				})
			}

			presentationNotesType, _ := tx.CurriculumCourseLessonResourceType.Where(
				tx.CurriculumCourseLessonResourceType.Name.Eq("presentation_notes"),
			).First()

			studentNotesType, _ := tx.CurriculumCourseLessonResourceType.Where(
				tx.CurriculumCourseLessonResourceType.Name.Eq("student_notes"),
			).First()

			teacherNotesType, _ := tx.CurriculumCourseLessonResourceType.Where(
				tx.CurriculumCourseLessonResourceType.Name.Eq("teacher_notes"),
			).First()

			miscMaterialsType, _ := tx.CurriculumCourseLessonResourceType.Where(
				tx.CurriculumCourseLessonResourceType.Name.Eq("misc_materials"),
			).First()

			/* associations: CurriculumCourseLevels*/
			for i, level := range form.Levels {
				entityCourseLevel := model.CurriculumCourseLevel{}

				if len(level.ID) > 1 {
					IDUUID, err := model.ValidUUIDExFromIDString(level.ID)
					if err != nil {
						return err
					}
					entityCourseLevel.ID = IDUUID
				}
				entityCourseLevel.CourseID = curriculumCourse.ID
				entityCourseLevel.Name = level.Name

				err = tx.CurriculumCourseLevel.Clauses(clause.OnConflict{
					UpdateAll: true,
				}).Create(&entityCourseLevel)
				if err != nil {
					return err
				}

				returnLevels := dto.CurriculumCourseLevels{}
				returnLevels.ID = entityCourseLevel.ID.ToString()
				returnLevels.Name = entityCourseLevel.Name
				returnLevels.IconID = entityCourseLevel.IconID.ToString()
				returnLevels.Description = entityCourseLevel.Description

				for j, lesson := range level.Lessons {
					entityLesson := model.CurriculumCourseLevelLesson{}

					if len(lesson.ID) > 1 {
						lessonIDUUID, err := model.ValidUUIDExFromIDString(lesson.ID)
						if err != nil {
							return err
						}
						entityLesson.ID = lessonIDUUID
					} else {
						entityLesson.LessonNumber = uint64(j + 1)
						entityLesson.CourseLevelID = entityCourseLevel.ID
					}

					err = tx.CurriculumCourseLevelLesson.Clauses(clause.OnConflict{
						UpdateAll: true,
					}).Create(&entityLesson)
					if err != nil {
						return err
					}

					lessonDTO := dto.CurriculumCourseLevelLessons{}
					lessonDTO.ID = entityLesson.ID.ToString()
					lessonDTO.LessonNumber = entityLesson.LessonNumber

					for k, presentationNote := range lesson.PresentationNotes {
						entityPresentationNote := model.CurriculumCourseLevelLessonResources{}

						if len(presentationNote.ID) > 1 {
							presentationNoteIDUUID, err := model.ValidUUIDExFromIDString(presentationNote.ID)
							if err != nil {
								return err
							}
							entityPresentationNote.ID = presentationNoteIDUUID
						}

						if len(presentationNote.ResourseID) > 1 {
							presentationNoteResourseIDUUID, err := model.ValidUUIDExFromIDString(presentationNote.ResourseID)
							if err != nil {
								return err
							}
							entityPresentationNote.ResourseID = presentationNoteResourseIDUUID
						}

						_, presentationNoteFileHeader, err := ctx.Request().FormFile(fmt.Sprintf("levels.%d.lessons.%d.presentation_notes.%d.file", i, j, k))
						if err == nil {
							file, err := utils.SaveUploadV2(presentationNoteFileHeader, &entityPresentationNote.ResourseID, []string{utils.PrefixCourseResourses, curriculumEntry.Description}, s3, tx, ctx)
							if err != nil {
								return err
							}
							entityPresentationNote.ResourseID = file.ID
						}
						entityPresentationNote.LessonID = entityLesson.ID
						entityPresentationNote.ResourseTypeID = presentationNotesType.ID

						err = tx.CurriculumCourseLevelLessonResources.Clauses(clause.OnConflict{
							UpdateAll: true,
						}).Create(&entityPresentationNote)
						if err != nil {
							return err
						}

						presentationNoteDTO := dto.CurriculumCourseLevelLessonResources{}
						presentationNoteDTO.ID = entityPresentationNote.ID.ToString()
						presentationNoteDTO.ResourseID = entityPresentationNote.ResourseID.ToString()

						lessonDTO.PresentationNotes = append(lessonDTO.PresentationNotes, presentationNoteDTO)
					}

					for k, studentNote := range lesson.StudentNotes {
						entityStudentNote := model.CurriculumCourseLevelLessonResources{}

						if len(studentNote.ID) > 1 {
							studentNoteIDUUID, err := model.ValidUUIDExFromIDString(studentNote.ID)
							if err != nil {
								return err
							}
							entityStudentNote.ID = studentNoteIDUUID
						}

						if len(studentNote.ResourseID) > 1 {
							entityStudentNoteResourseIDUUID, err := model.ValidUUIDExFromIDString(studentNote.ResourseID)
							if err != nil {
								return err
							}
							entityStudentNote.ResourseID = entityStudentNoteResourseIDUUID
						}

						_, studentNoteFileHeader, err := ctx.Request().FormFile(fmt.Sprintf("levels.%d.lessons.%d.student_notes.%d.file", i, j, k))
						if err == nil {
							file, err := utils.SaveUploadV2(studentNoteFileHeader, &entityStudentNote.ResourseID, []string{utils.PrefixCourseResourses, curriculumEntry.Description}, s3, tx, ctx)
							if err != nil {
								return err
							}
							entityStudentNote.ResourseID = file.ID
						}
						entityStudentNote.LessonID = entityLesson.ID
						entityStudentNote.ResourseTypeID = studentNotesType.ID

						err = tx.CurriculumCourseLevelLessonResources.Clauses(clause.OnConflict{
							UpdateAll: true,
						}).Create(&entityStudentNote)
						if err != nil {
							return err
						}

						studentNoteDTO := dto.CurriculumCourseLevelLessonResources{}
						studentNoteDTO.ID = entityStudentNote.ID.ToString()
						studentNoteDTO.ResourseID = entityStudentNote.ResourseID.ToString()

						lessonDTO.StudentNotes = append(lessonDTO.StudentNotes, studentNoteDTO)
					}

					for k, teacherNote := range lesson.TeacherNotes {
						entityTeacherNote := model.CurriculumCourseLevelLessonResources{}

						if len(teacherNote.ID) > 1 {
							teacherNoteIDUUID, err := model.ValidUUIDExFromIDString(teacherNote.ID)
							if err != nil {
								return err
							}
							entityTeacherNote.ID = teacherNoteIDUUID
						}

						if len(teacherNote.ResourseID) > 1 {
							teacherNoteResourseIDUUID, err := model.ValidUUIDExFromIDString(teacherNote.ResourseID)
							if err != nil {
								return err
							}
							entityTeacherNote.ResourseID = teacherNoteResourseIDUUID
						}

						_, teacherNoteFileHeader, err := ctx.Request().FormFile(fmt.Sprintf("levels.%d.lessons.%d.teacher_notes.%d.file", i, j, k))
						if err == nil {
							file, err := utils.SaveUploadV2(teacherNoteFileHeader, &entityTeacherNote.ResourseID, []string{utils.PrefixCourseResourses, curriculumEntry.Description}, s3, tx, ctx)
							if err != nil {
								return err
							}
							entityTeacherNote.ResourseID = file.ID
						}
						entityTeacherNote.LessonID = entityLesson.ID
						entityTeacherNote.ResourseTypeID = teacherNotesType.ID
						err = tx.CurriculumCourseLevelLessonResources.Clauses(clause.OnConflict{
							UpdateAll: true,
						}).Create(&entityTeacherNote)
						if err != nil {
							return err
						}

						teacherNoteDTO := dto.CurriculumCourseLevelLessonResources{}
						teacherNoteDTO.ID = entityTeacherNote.ID.ToString()
						teacherNoteDTO.ResourseID = entityTeacherNote.ResourseID.ToString()

						lessonDTO.TeacherNotes = append(lessonDTO.TeacherNotes, teacherNoteDTO)
					}

					for k, miscMaterial := range lesson.MiscMaterials {
						entityMiscMaterial := model.CurriculumCourseLevelLessonResources{}

						if len(miscMaterial.ID) > 1 {
							miscMaterialIDUUID, err := model.ValidUUIDExFromIDString(miscMaterial.ID)
							if err != nil {
								return err
							}
							entityMiscMaterial.ID = miscMaterialIDUUID
						}

						if len(miscMaterial.ResourseID) > 1 {
							miscMaterialResourseIDUUID, err := model.ValidUUIDExFromIDString(miscMaterial.ResourseID)
							if err != nil {
								return err
							}
							entityMiscMaterial.ResourseID = miscMaterialResourseIDUUID
						}

						_, miscMaterialFileHeader, err := ctx.Request().FormFile(fmt.Sprintf("levels.%d.lessons.%d.misc_materials.%d.file", i, j, k))
						if err == nil {
							file, err := utils.SaveUploadV2(miscMaterialFileHeader, &entityMiscMaterial.ResourseID, []string{utils.PrefixCourseResourses, curriculumEntry.Description}, s3, tx, ctx)
							if err != nil {
								return err
							}
							entityMiscMaterial.ResourseID = file.ID
						}
						entityMiscMaterial.LessonID = entityLesson.ID
						entityMiscMaterial.ResourseTypeID = miscMaterialsType.ID
						err = tx.CurriculumCourseLevelLessonResources.Clauses(clause.OnConflict{
							UpdateAll: true,
						}).Create(&entityMiscMaterial)
						if err != nil {
							return err
						}

						miscMaterialDTO := dto.CurriculumCourseLevelLessonResources{}
						miscMaterialDTO.ID = entityMiscMaterial.ID.ToString()
						miscMaterialDTO.ResourseID = entityMiscMaterial.ResourseID.ToString()

						lessonDTO.TeacherNotes = append(lessonDTO.TeacherNotes, teacherNoteDTO)
					}
					returnLevels.Lessons = append(returnLevels.Lessons, lessonDTO)
				}

				returnForm.Levels = append(returnForm.Levels, returnLevels)
			}

			returnForm.ID = curriculumEntry.ID.ToString()
			returnForm.Description = curriculumEntry.Description
			returnForm.IconID = curriculumEntry.IconID.ToString()
			if curriculumEntry.ParentID != nil {
				returnForm.ParentID = (*curriculumEntry.ParentID).ToString()
			}
			// returnForm.CurriculumPlanID = curriculumCourse.CurriculumPlanID.ToString()
			// returnForm.CourseID = curriculumCourse.ID.ToString()

			// levels
			// youtube_video_entries

			return nil
		})

		// var entryToSave = model.CurriculumEntry{}
		// entryToSave.Description = form.Description

		// if len(form.ID) > 1 {
		// 	IDUUID, err := model.ValidUUIDExFromIDString(form.ID)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	tx.First(&entryToSave, "`id` = ?", IDUUID)
		// }

		// if len(form.IconID) > 1 {
		// 	IconIDUUID, err := model.ValidUUIDExFromIDString(form.IconID)
		// 	entryToSave.IconID = &IconIDUUID
		// 	if err != nil {
		// 		return err
		// 	}
		// }

		// // // Get the max post value size passed via iris.WithPostMaxMemory.
		// maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()

		// err = ctx.Request().ParseMultipartForm(maxSize)
		// if err != nil {
		// 	return err
		// }

		// _, iconFileHeader, err := ctx.Request().FormFile("icon_file")
		// if err == nil {
		// 	file, err := utils.SaveUpload(iconFileHeader, []string{utils.PrefixCourseResourses, entryToSave.Description}, s3, tx, ctx)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	entryToSave.IconID = &file.ID
		// }

		// if entryToSave.IconID == nil {
		// 	return fmt.Errorf("no icon id")
		// }

		// if len(form.ParentID) > 1 && form.ParentID != "null" {
		// 	parentIDUUID, err := model.ValidUUIDExFromIDString(form.ParentID)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	entryToSave.ParentID = &parentIDUUID

		// 	tx.Model(&model.CurriculumEntry{}).
		// 		Select("MAX(`seq_no_same_level`)").
		// 		Where("`parent_id` = ?", *entryToSave.ParentID).
		// 		Group("`parent_id`").
		// 		Scan(&entryToSave.SeqNoSameLevel)
		// 	entryToSave.SeqNoSameLevel = entryToSave.SeqNoSameLevel + 1
		// }

		// if err := tx.Save(&entryToSave).Error; err != nil {
		// 	return err
		// }

		// if err := tx.Delete(&model.CurriculumCourseBlogEntries{}, "`entry_id` = ?", entryToSave.ID).Error; err != nil {
		// 	return err
		// }

		// // if err := tx.Delete(&model.CurriculumCourseInformationEntries{}, "`entry_id` = ?", entryToSave.ID).Error; err != nil {
		// // 	return err
		// // }

		// if err := tx.Delete(&model.CurriculumCourseYoutubeVideoEntries{}, "`entry_id` = ?", entryToSave.ID).Error; err != nil {
		// 	return err
		// }

		// if form.BlogEntries != nil {
		// 	for _, blogEntry := range form.BlogEntries {
		// 		blogEntryModel := model.CurriculumCourseBlogEntries{}
		// 		blogEntryModel.ID = blogEntry.ID
		// 		blogEntryModel.ExternalURL = blogEntry.ExternalURL
		// 		blogEntryModel.Title = blogEntry.Title
		// 		blogEntryModel.EntryID = &entryToSave.ID

		// 		if err := tx.Clauses(clause.OnConflict{
		// 			Columns:   []clause.Column{{Name: "id"}},
		// 			DoUpdates: clause.AssignmentColumns([]string{"external_url", "title", "entry_id"}),
		// 		}).Create(&blogEntryModel).Error; err != nil {
		// 			return err
		// 		}
		// 	}
		// }

		// // if form.InformationEntries != nil {
		// // 	for i, informationEntry := range form.InformationEntries {
		// // 		informationEntryModel := model.CurriculumCourseInformationEntries{}
		// // 		informationEntryModel.Title = informationEntry.Title
		// // 		informationEntryModel.Content = informationEntry.Content
		// // 		// informationEntryModel.EntryID = &entryToSave.ID

		// // 		if len(informationEntry.IconID) > 1 {
		// // 			IconIDUUID, err := model.ValidUUIDExFromIDString(informationEntry.IconID)
		// // 			informationEntryModel.IconID = &IconIDUUID
		// // 			if err != nil {
		// // 				return err
		// // 			}
		// // 		}

		// // 		_, iconFileHeader, err := ctx.Request().FormFile(fmt.Sprintf("information_entries.%d.icon_file", i))
		// // 		if err == nil {
		// // 			file, err := utils.SaveUpload(iconFileHeader, []string{utils.PrefixCourseResourses, entryToSave.Description}, s3, tx, ctx)
		// // 			if err != nil {
		// // 				return err
		// // 			}
		// // 			informationEntryModel.IconID = &file.ID
		// // 		}

		// // 		if informationEntryModel.IconID == nil {
		// // 			return fmt.Errorf("no icon id")
		// // 		}

		// // 		if err := tx.Clauses(clause.OnConflict{
		// // 			Columns:   []clause.Column{{Name: "id"}},
		// // 			DoUpdates: clause.AssignmentColumns([]string{"icon_id", "title", "content", "entry_id"}),
		// // 		}).Create(&informationEntryModel).Error; err != nil {
		// // 			return err
		// // 		}
		// // 	}
		// // }

		// if form.YoutubeVideoEntries != nil {
		// 	for _, youtubeVideoEntry := range form.YoutubeVideoEntries {
		// 		youtubeVideoEntryModel := model.CurriculumCourseYoutubeVideoEntries{}
		// 		youtubeVideoEntryModel.ID = youtubeVideoEntry.ID
		// 		youtubeVideoEntryModel.URL = youtubeVideoEntry.URL
		// 		youtubeVideoEntryModel.EntryID = &entryToSave.ID

		// 		if err := tx.Clauses(clause.OnConflict{
		// 			Columns:   []clause.Column{{Name: "id"}},
		// 			DoUpdates: clause.AssignmentColumns([]string{"url", "title", "entry_id"}),
		// 		}).Create(&youtubeVideoEntryModel).Error; err != nil {
		// 			return err
		// 		}
		// 	}
		// }

		// // // return nil will commit the whole transaction

		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
		} else {
			ctx.JSON(returnForm)
		}

		// ctx.JSON(dto.CurriculumCourseDetails{
		// 	ID:          curriculumEntry.ID,
		// 	Description: curriculumEntry.Description,
		// 	IconID:      curriculumEntry.IconID,
		// 	ParentID:    curriculumEntry.ParentID,
		// 	//Prerequisites: []string
		// 	YoutubeVideoURLs: curriculumCourseYoutubeVideoEntries,
		// 	// InformationEntries: curriculumCourseInformationEntries,
		// 	BlogEntries: curriculumCourseBlogEntries,
		// })
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
