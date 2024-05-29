package bo

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/mappings"
	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/query"
	"github.com/dirkarnez/stemexapi/utils"
	"github.com/samber/lo"
	"gorm.io/gorm/clause"
)

func CreateOrUpdateCurriculumCourseType(form *dto.CurriculumCourseTypeForm, s3 *utils.StemexS3Client, txOrQ *query.Query) (*dto.CurriculumCourseTypeForm, error) {
	var returnForm dto.CurriculumCourseTypeForm

	err := txOrQ.Transaction(func(tx *query.Query) error {
		curriculumEntry := model.CurriculumEntry{}

		err := mappings.MapCurriculumCourseTypeFormToCurriculumEntry(form, &curriculumEntry, s3, tx)
		if err != nil {
			return err
		}

		s3Prefix := []string{utils.PrefixCourseResourses, strings.ToLower(curriculumEntry.Description)}

		curriculumEntry.IconID, err = model.ValidUUIDExFromIDString(form.IconID)
		if err != nil {
			return err
		}

		file, err := utils.SaveUploadV2(form.IconFile, &curriculumEntry.IconID, s3Prefix, s3, txOrQ)
		if err != nil {
			return fmt.Errorf("form.IconFile %s", err.Error())
		}

		curriculumEntry.IconID = file.ID

		err = tx.CurriculumEntry.Clauses(clause.OnConflict{UpdateAll: true}).Create(&curriculumEntry)
		if err != nil {
			return err
		} else {
			returnForm.ID = curriculumEntry.ID.ToString()
			returnForm.Description = curriculumEntry.Description
			returnForm.IconID = curriculumEntry.IconID.ToString()
			if curriculumEntry.ParentID != nil {
				returnForm.ParentID = (*curriculumEntry.ParentID).ToString()
			}

			return nil
		}
	})

	if err != nil {
		return nil, err
	} else {
		return &returnForm, nil
	}

	// var entryToSave = model.CurriculumEntry{}

	// err := dbInstance.Transaction(func(tx *gorm.DB) error {
	// 	// type InformationEntry struct {
	// 	// 	IconID string `form:"icon_id"`
	// 	// 	//IconFile []byte/**multipart.FileHeader*/ `form:"icon_file"`
	// 	// 	Title   string `form:"title"`
	// 	// 	Content string `form:"content"`
	// 	// }

	// 	var form dto.CurriculumCourseTypeForm
	// 	err := ctx.ReadForm(&form)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// if len(form.ID) > 1 {

	// 	// }

	// 	entryToSave.ID, err = model.ValidUUIDExFromIDString(form.ID)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	tx.First(&entryToSave, "`id` = ?", entryToSave.ID)

	// 	entryToSave.Description = form.Description
	// 	entryToSave.IconID, err = model.ValidUUIDExFromIDString(form.IconID)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// // Get the max post value size passed via iris.WithPostMaxMemory.
	// 	maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()

	// 	err = ctx.Request().ParseMultipartForm(maxSize)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	_, iconFileHeader, err := ctx.Request().FormFile("icon_file")
	// 	entryToSave.IconID, err = model.ValidUUIDExFromIDString(form.IconID)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	file, err := utils.SaveUploadV2(iconFileHeader, &entryToSave.IconID, []string{utils.PrefixCourseResourses, strings.ToLower(entryToSave.Description)}, s3, query.Use(tx))
	// 	if err != nil {
	// 		return fmt.Errorf("CurriculumPlanFile: %s", err.Error())
	// 	}
	// 	entryToSave.IconID = file.ID

	// 	if len(form.ParentID) > 1 && form.ParentID != "null" {
	// 		parentIDUUID, err := model.ValidUUIDExFromIDString(form.ParentID)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		entryToSave.ParentID = &parentIDUUID

	// 		tx.Model(&model.CurriculumEntry{}).
	// 			Select("MAX(`seq_no_same_level`)").
	// 			Where("`parent_id` = ?", *entryToSave.ParentID).
	// 			Group("`parent_id`").
	// 			Scan(&entryToSave.SeqNoSameLevel)
	// 		entryToSave.SeqNoSameLevel = entryToSave.SeqNoSameLevel + 1
	// 	}

	// 	if err := tx.Save(&entryToSave).Error; err != nil {
	// 		return err
	// 	}

	// 	// // return nil will commit the whole transaction
	// 	return nil
	// })

	// if err != nil {
	// 	ctx.StopWithError(iris.StatusInternalServerError, err)
	// } else {
	// 	var returnForm dto.CurriculumCourseTypeForm
	// 	returnForm.Description = entryToSave.Description
	// 	returnForm.ID = entryToSave.ID.ToString()
	// 	returnForm.IconID = entryToSave.IconID.ToString()

	// 	if entryToSave.ParentID != nil {
	// 		returnForm.ParentID = (*entryToSave.ParentID).ToString()
	// 	}

	// 	ctx.JSON(returnForm)
	// }
	// return
}

func CreateOrUpdateCurriculumCourse(form *dto.CurriculumCourseForm, s3 *utils.StemexS3Client, txOrQ *query.Query) (*dto.CurriculumCourseForm, error) {
	var returnForm dto.CurriculumCourseForm

	err := txOrQ.Transaction(func(tx *query.Query) error {
		curriculumEntry := model.CurriculumEntry{}

		err := mappings.MapCurriculumCourseFormToCurriculumEntry(form, &curriculumEntry, s3, tx)
		if err != nil {
			return err
		}

		s3Prefix := []string{utils.PrefixCourseResourses, strings.ToLower(curriculumEntry.Description)}

		curriculumEntry.IconID, err = model.ValidUUIDExFromIDString(form.IconID)
		if err != nil {
			return err
		}

		file, err := utils.SaveUploadV2(form.IconFile, &curriculumEntry.IconID, s3Prefix, s3, txOrQ)
		if err != nil {
			return fmt.Errorf("form.IconFile %s", err.Error())
		}

		curriculumEntry.IconID = file.ID

		err = tx.CurriculumEntry.Clauses(clause.OnConflict{UpdateAll: true}).Create(&curriculumEntry)
		if err != nil {
			return err
		}

		/* associations: CurriculumCourse */
		var curriculumCourse = model.CurriculumCourse{EntryID: curriculumEntry.ID}
		curriculumCourse.ID, err = model.ValidUUIDExFromIDString(form.CourseID)
		if err != nil {
			return err
		}

		if form.CurriculumPlanFile != nil {
			curriculumCourse.CurriculumPlanID, err = model.ValidUUIDExPointerFromIDString(form.CurriculumPlanID)
			if err != nil {
				return err
			}

			file, err = utils.SaveUploadV2(form.CurriculumPlanFile, curriculumCourse.CurriculumPlanID, s3Prefix, s3, tx)
			if err != nil {
				return fmt.Errorf("CurriculumPlanFile: %s", err.Error())
			}
			curriculumCourse.CurriculumPlanID = &file.ID
		}

		err = tx.CurriculumCourse.Clauses(clause.OnConflict{UpdateAll: true}).Create(&curriculumCourse)
		if err != nil {
			return err
		}

		returnForm.CourseID = curriculumCourse.ID.ToString()

		/* associations: CurriculumCourseBlogEntries*/
		var blogs []*model.CurriculumCourseBlogEntries
		for _, dto := range form.BlogEntries {
			entity := model.CurriculumCourseBlogEntries{}
			entity.ID, err = model.ValidUUIDExFromIDString(dto.ID)
			if err != nil {
				return err
			}
			entity.ExternalURL = dto.ExternalURL
			entity.Title = dto.Title
			entity.EntryID = &curriculumEntry.ID
			blogs = append(blogs, &entity)
		}

		err = tx.CurriculumCourseBlogEntries.Clauses(clause.OnConflict{UpdateAll: true}).Create(blogs...)
		if err != nil {
			return err
		}

		tx.CurriculumCourseBlogEntries.
			Unscoped().
			Where(tx.CurriculumCourseBlogEntries.EntryID.Eq(curriculumEntry.ID)).
			Not(tx.CurriculumCourseBlogEntries.ID.In(lo.Map(blogs, func(blog *model.CurriculumCourseBlogEntries, index int) driver.Valuer {
				return blog.ID
			})...)).
			Delete()

		returnForm.BlogEntries = lo.Map(blogs, func(saved *model.CurriculumCourseBlogEntries, index int) dto.CurriculumCourseBlogEntries {
			return dto.CurriculumCourseBlogEntries{
				ID:          saved.ID.ToString(),
				ExternalURL: saved.ExternalURL,
				Title:       saved.Title,
			}
		})

		/* associations: CurriculumCourseYoutubeVideoEntries*/
		var youtubes []*model.CurriculumCourseYoutubeVideoEntries
		for _, dto := range form.YoutubeVideoEntries {
			entity := model.CurriculumCourseYoutubeVideoEntries{}

			entity.ID, err = model.ValidUUIDExFromIDString(dto.ID)
			if err != nil {
				return err
			}

			entity.URL = dto.URL
			entity.EntryID = &curriculumEntry.ID
			youtubes = append(youtubes, &entity)
		}

		err = tx.CurriculumCourseYoutubeVideoEntries.Clauses(clause.OnConflict{UpdateAll: true}).Create(youtubes...)
		if err != nil {
			return err
		}

		tx.CurriculumCourseYoutubeVideoEntries.
			Unscoped().
			Where(tx.CurriculumCourseYoutubeVideoEntries.EntryID.Eq(curriculumEntry.ID)).
			Not(tx.CurriculumCourseYoutubeVideoEntries.ID.In(lo.Map(youtubes, func(youtube *model.CurriculumCourseYoutubeVideoEntries, index int) driver.Valuer {
				return youtube.ID
			})...)).
			Delete()

		returnForm.YoutubeVideoEntries = lo.Map(youtubes, func(saved *model.CurriculumCourseYoutubeVideoEntries, index int) dto.CurriculumCourseYoutubeVideoEntries {
			return dto.CurriculumCourseYoutubeVideoEntries{
				ID:  saved.ID.ToString(),
				URL: saved.URL,
			}
		})

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
		var levelEntityList []*model.CurriculumCourseLevel
		for _, level := range form.Levels {
			entityCourseLevel := model.CurriculumCourseLevel{}

			entityCourseLevel.ID, err = model.ValidUUIDExFromIDString(level.ID)
			if err != nil {
				return err
			}

			entityCourseLevel.IconID, err = model.ValidUUIDExFromIDString(level.IconID)
			if err != nil {
				return err
			}

			iconFile, err := utils.SaveUploadV2(level.IconFile, &entityCourseLevel.IconID, s3Prefix, s3, tx)
			if err != nil {
				return fmt.Errorf("iconFile: %s", err.Error())
			}
			entityCourseLevel.IconID = iconFile.ID

			entityCourseLevel.CourseID = curriculumCourse.ID
			entityCourseLevel.Name = level.Name
			entityCourseLevel.Title = level.Title
			entityCourseLevel.Description = level.Description

			err = tx.CurriculumCourseLevel.Clauses(clause.OnConflict{UpdateAll: true}).Create(&entityCourseLevel)
			if err != nil {
				return err
			}

			returnLevels := dto.CurriculumCourseLevels{}
			returnLevels.ID = entityCourseLevel.ID.ToString()
			returnLevels.Name = entityCourseLevel.Name
			returnLevels.IconID = entityCourseLevel.IconID.ToString()
			returnLevels.Title = entityCourseLevel.Title
			returnLevels.Description = entityCourseLevel.Description

			var lessonEntityList []*model.CurriculumCourseLevelLesson
			for _, lesson := range level.Lessons {
				entityLesson := model.CurriculumCourseLevelLesson{}

				entityLesson.ID, err = model.ValidUUIDExFromIDString(lesson.ID)
				if err != nil {
					return err
				}

				entityLesson.LessonNumber = lesson.LessonNumber
				entityLesson.CourseLevelID = entityCourseLevel.ID

				err = tx.CurriculumCourseLevelLesson.Clauses(clause.OnConflict{UpdateAll: true}).Create(&entityLesson)
				if err != nil {
					return err
				}

				fmt.Printf("lesson %s is referencing level %s", entityLesson.ID.ToString(), entityCourseLevel.ID.ToString())

				lessonDTO := dto.CurriculumCourseLevelLessons{}
				lessonDTO.ID = entityLesson.ID.ToString()
				lessonDTO.LessonNumber = entityLesson.LessonNumber

				var presentationNoteInsertedList []*model.CurriculumCourseLevelLessonResources
				for _, presentationNote := range lesson.PresentationNotes {
					entityPresentationNote := model.CurriculumCourseLevelLessonResources{}

					entityPresentationNote.ID, err = model.ValidUUIDExFromIDString(presentationNote.ID)
					if err != nil {
						return err
					}

					entityPresentationNote.ResourseID, err = model.ValidUUIDExFromIDString(presentationNote.ResourseID)
					if err != nil {
						return err
					}

					file, err := utils.SaveUploadV2(presentationNote.File, &entityPresentationNote.ResourseID, s3Prefix, s3, tx)
					if err != nil {
						return fmt.Errorf("presentationNote: %s", err.Error())
					}
					entityPresentationNote.ResourseID = file.ID

					entityPresentationNote.LessonID = entityLesson.ID
					entityPresentationNote.ResourseTypeID = presentationNotesType.ID

					err = tx.CurriculumCourseLevelLessonResources.Clauses(clause.OnConflict{UpdateAll: true}).Create(&entityPresentationNote)
					if err != nil {
						return err
					}

					presentationNoteDTO := dto.CurriculumCourseLevelLessonResources{}
					presentationNoteDTO.ID = entityPresentationNote.ID.ToString()
					presentationNoteDTO.ResourseID = entityPresentationNote.ResourseID.ToString()

					presentationNoteInsertedList = append(presentationNoteInsertedList, &entityPresentationNote)
					lessonDTO.PresentationNotes = append(lessonDTO.PresentationNotes, presentationNoteDTO)
				}

				tx.CurriculumCourseLevelLessonResources.
					Unscoped().
					Where(tx.CurriculumCourseLevelLessonResources.LessonID.Eq(entityLesson.ID), tx.CurriculumCourseLevelLessonResources.ResourseTypeID.Eq(presentationNotesType.ID)).
					Not(tx.CurriculumCourseLevelLessonResources.ID.In(lo.Map(presentationNoteInsertedList, func(presentationNoteInserted *model.CurriculumCourseLevelLessonResources, index int) driver.Valuer {
						return presentationNoteInserted.ID
					})...)).
					Delete()

				var studentNoteInsertedList []*model.CurriculumCourseLevelLessonResources
				for _, studentNote := range lesson.StudentNotes {
					entityStudentNote := model.CurriculumCourseLevelLessonResources{}

					entityStudentNote.ID, err = model.ValidUUIDExFromIDString(studentNote.ID)
					if err != nil {
						return err
					}

					entityStudentNote.ResourseID, err = model.ValidUUIDExFromIDString(studentNote.ResourseID)
					if err != nil {
						return err
					}

					file, err := utils.SaveUploadV2(studentNote.File, &entityStudentNote.ResourseID, s3Prefix, s3, tx)
					if err != nil {
						return fmt.Errorf("studentNote: %s", err.Error())
					}
					entityStudentNote.ResourseID = file.ID

					entityStudentNote.LessonID = entityLesson.ID
					entityStudentNote.ResourseTypeID = studentNotesType.ID

					err = tx.CurriculumCourseLevelLessonResources.Clauses(clause.OnConflict{UpdateAll: true}).Create(&entityStudentNote)
					if err != nil {
						return err
					}

					studentNoteDTO := dto.CurriculumCourseLevelLessonResources{}
					studentNoteDTO.ID = entityStudentNote.ID.ToString()
					studentNoteDTO.ResourseID = entityStudentNote.ResourseID.ToString()

					studentNoteInsertedList = append(studentNoteInsertedList, &entityStudentNote)
					lessonDTO.StudentNotes = append(lessonDTO.StudentNotes, studentNoteDTO)
				}

				tx.CurriculumCourseLevelLessonResources.
					Unscoped().
					Where(tx.CurriculumCourseLevelLessonResources.LessonID.Eq(entityLesson.ID), tx.CurriculumCourseLevelLessonResources.ResourseTypeID.Eq(studentNotesType.ID)).
					Not(tx.CurriculumCourseLevelLessonResources.ID.In(lo.Map(studentNoteInsertedList, func(studentNoteInserted *model.CurriculumCourseLevelLessonResources, index int) driver.Valuer {
						return studentNoteInserted.ID
					})...)).
					Delete()

				var teacherNoteInsertedList []*model.CurriculumCourseLevelLessonResources
				for _, teacherNote := range lesson.TeacherNotes {
					entityTeacherNote := model.CurriculumCourseLevelLessonResources{}

					entityTeacherNote.ID, err = model.ValidUUIDExFromIDString(teacherNote.ID)
					if err != nil {
						return err
					}

					entityTeacherNote.ResourseID, err = model.ValidUUIDExFromIDString(teacherNote.ResourseID)
					if err != nil {
						return err
					}

					file, err := utils.SaveUploadV2(teacherNote.File, &entityTeacherNote.ResourseID, s3Prefix, s3, tx)
					if err != nil {
						return fmt.Errorf("teacherNote: %s", err.Error())
					}
					entityTeacherNote.ResourseID = file.ID

					entityTeacherNote.LessonID = entityLesson.ID
					entityTeacherNote.ResourseTypeID = teacherNotesType.ID
					err = tx.CurriculumCourseLevelLessonResources.Clauses(clause.OnConflict{UpdateAll: true}).Create(&entityTeacherNote)
					if err != nil {
						return err
					}

					teacherNoteDTO := dto.CurriculumCourseLevelLessonResources{}
					teacherNoteDTO.ID = entityTeacherNote.ID.ToString()
					teacherNoteDTO.ResourseID = entityTeacherNote.ResourseID.ToString()

					teacherNoteInsertedList = append(teacherNoteInsertedList, &entityTeacherNote)
					lessonDTO.TeacherNotes = append(lessonDTO.TeacherNotes, teacherNoteDTO)
				}

				tx.CurriculumCourseLevelLessonResources.
					Unscoped().
					Where(tx.CurriculumCourseLevelLessonResources.LessonID.Eq(entityLesson.ID), tx.CurriculumCourseLevelLessonResources.ResourseTypeID.Eq(teacherNotesType.ID)).
					Not(tx.CurriculumCourseLevelLessonResources.ID.In(lo.Map(teacherNoteInsertedList, func(teacherNoteInserted *model.CurriculumCourseLevelLessonResources, index int) driver.Valuer {
						return teacherNoteInserted.ID
					})...)).
					Delete()

				var miscMaterialInsertedList []*model.CurriculumCourseLevelLessonResources
				for _, miscMaterial := range lesson.MiscMaterials {
					entityMiscMaterial := model.CurriculumCourseLevelLessonResources{}

					if len(miscMaterial.ID) > 1 {
						miscMaterialIDUUID, err := model.ValidUUIDExFromIDString(miscMaterial.ID)
						if err != nil {
							return err
						}
						entityMiscMaterial.ID = miscMaterialIDUUID
					}

					entityMiscMaterial.ResourseID, err = model.ValidUUIDExFromIDString(miscMaterial.ResourseID)
					if err != nil {
						return err
					}

					file, err := utils.SaveUploadV2(miscMaterial.File, &entityMiscMaterial.ResourseID, s3Prefix, s3, tx)
					if err != nil {
						return fmt.Errorf("miscMaterial: %s", err.Error())
					}
					entityMiscMaterial.ResourseID = file.ID

					entityMiscMaterial.LessonID = entityLesson.ID
					entityMiscMaterial.ResourseTypeID = miscMaterialsType.ID

					err = tx.CurriculumCourseLevelLessonResources.Clauses(clause.OnConflict{UpdateAll: true}).Create(&entityMiscMaterial)
					if err != nil {
						return err
					}

					miscMaterialDTO := dto.CurriculumCourseLevelLessonResources{}
					miscMaterialDTO.ID = entityMiscMaterial.ID.ToString()
					miscMaterialDTO.ResourseID = entityMiscMaterial.ResourseID.ToString()

					miscMaterialInsertedList = append(miscMaterialInsertedList, &entityMiscMaterial)
					lessonDTO.MiscMaterials = append(lessonDTO.MiscMaterials, miscMaterialDTO)
				}

				tx.CurriculumCourseLevelLessonResources.
					Unscoped().
					Where(tx.CurriculumCourseLevelLessonResources.LessonID.Eq(entityLesson.ID), tx.CurriculumCourseLevelLessonResources.ResourseTypeID.Eq(miscMaterialsType.ID)).
					Not(tx.CurriculumCourseLevelLessonResources.ID.In(lo.Map(miscMaterialInsertedList, func(miscMaterialInserted *model.CurriculumCourseLevelLessonResources, index int) driver.Valuer {
						return miscMaterialInserted.ID
					})...)).
					Delete()

				lessonEntityList = append(lessonEntityList, &entityLesson)
				returnLevels.Lessons = append(returnLevels.Lessons, lessonDTO)
			}

			tx.CurriculumCourseLevelLesson.
				Unscoped().
				Where(tx.CurriculumCourseLevelLesson.CourseLevelID.Eq(entityCourseLevel.ID)).
				Not(tx.CurriculumCourseLevelLesson.ID.In(lo.Map(lessonEntityList, func(lessonEntity *model.CurriculumCourseLevelLesson, index int) driver.Valuer {
					return lessonEntity.ID
				})...)).
				Delete()

			levelEntityList = append(levelEntityList, &entityCourseLevel)
			returnForm.Levels = append(returnForm.Levels, returnLevels)
		}

		tx.CurriculumCourseLevel.
			Unscoped().
			Where(tx.CurriculumCourseLevel.CourseID.Eq(curriculumCourse.ID)).
			Not(tx.CurriculumCourseLevel.ID.In(lo.Map(levelEntityList, func(levelEntity *model.CurriculumCourseLevel, index int) driver.Valuer {
				return levelEntity.ID
			})...)).
			Delete()

		returnForm.ID = curriculumEntry.ID.ToString()
		returnForm.Description = curriculumEntry.Description
		returnForm.IconID = curriculumEntry.IconID.ToString()
		if curriculumEntry.ParentID != nil {
			returnForm.ParentID = (*curriculumEntry.ParentID).ToString()
		}

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

	if err != nil {
		return nil, err
	} else {
		return &returnForm, nil
	}
}
