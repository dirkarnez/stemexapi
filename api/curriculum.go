package api

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/albrow/forms"
	"github.com/dirkarnez/stemexapi/bo"
	"github.com/dirkarnez/stemexapi/datatypes"
	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/query"
	"github.com/dirkarnez/stemexapi/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/samber/lo"
	"gorm.io/gen/field"
	"gorm.io/gorm"
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
				ctx.StopWithError(iris.StatusNotFound, fmt.Errorf("invalid id"))
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
		// 		ctx.StopWithStatus(iris.StatusNotFound)
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
				ctx.StopWithError(iris.StatusNotFound, fmt.Errorf("invalid id"))
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
		} else {
			ctx.JSON(curriculumEntry)
		}
	}
}

// func GetCurriculumCourses(dbInstance *gorm.DB) context.Handler {
// 	return func(ctx iris.Context) {
// 		parentID := ctx.URLParam("parent-id")

// 		if len(parentID) < 1 {
// 			ctx.StopWithStatus(iris.StatusForbidden)
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

		form, err := MapRequestToCurriculumCourseTypeForm(ctx.Request())
		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
			return
		}

		returnForm, err := bo.CreateOrUpdateCurriculumCourseType(form, s3, query.Use(dbInstance))
		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
		} else {
			ctx.JSON(returnForm)
		}
	}
}

func GetCurriculumCourse(s3 *utils.StemexS3Client, dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		ID := ctx.URLParamDefault("id", "")
		var returnForm dto.CurriculumCourseForm

		var curriculumEntry *model.CurriculumEntry
		var curriculumCourseYoutubeVideoEntries []*model.CurriculumCourseYoutubeVideoEntries
		var curriculumCourseBlogEntries []*model.CurriculumCourseBlogEntries

		var curriculumCourse *model.CurriculumCourse

		var curriculumCourseLevels []*model.CurriculumCourseLevel
		var curriculumCourseLevelLessons []*model.CurriculumCourseLevelLesson
		var ppts []*model.CurriculumCourseLevelLessonResources
		var students []*model.CurriculumCourseLevelLessonResources
		var teachers []*model.CurriculumCourseLevelLessonResources
		var miscs []*model.CurriculumCourseLevelLessonResources

		var err error
		var q = query.Use(dbInstance)

		if len(ID) == 0 {
			ctx.StopWithError(iris.StatusNotFound, fmt.Errorf("no id"))
			return
		}

		idUUID, err := model.ValidUUIDExFromIDString(ID)
		if err != nil {
			ctx.StopWithError(iris.StatusNotFound, fmt.Errorf("invalid id"))
			return
		}

		err = q.Transaction(func(tx *query.Query) error {
			curriculumEntry, err = tx.CurriculumEntry.
				Select(q.CurriculumEntry.ALL).
				LeftJoin(q.CurriculumCourse, q.CurriculumEntry.ID.EqCol(q.CurriculumCourse.EntryID)).
				Where(q.CurriculumEntry.ID.Eq(idUUID)).
				Where(q.CurriculumCourse.ID.IsNotNull()).
				First()
			if err != nil {
				return err
			}

			if curriculumEntry == nil {
				return fmt.Errorf("not found")
			}

			curriculumCourseYoutubeVideoEntries, err = tx.CurriculumCourseYoutubeVideoEntries.
				Select(q.CurriculumCourseYoutubeVideoEntries.ALL).
				LeftJoin(q.CurriculumEntry, q.CurriculumCourseYoutubeVideoEntries.EntryID.EqCol(q.CurriculumEntry.ID)).
				Where(q.CurriculumEntry.ID.Eq(idUUID)).
				Find()
			if err != nil {
				return err
			}

			curriculumCourseBlogEntries, err = tx.CurriculumCourseBlogEntries.
				Select(q.CurriculumCourseBlogEntries.ALL).
				LeftJoin(q.CurriculumEntry, q.CurriculumCourseBlogEntries.EntryID.EqCol(q.CurriculumEntry.ID)).
				Where(q.CurriculumEntry.ID.Eq(idUUID)).
				Find()
			if err != nil {
				return err
			}

			curriculumCourse, err = tx.CurriculumCourse.
				Select(q.CurriculumCourse.ALL).
				LeftJoin(q.CurriculumEntry, q.CurriculumCourse.EntryID.EqCol(q.CurriculumEntry.ID)).
				Where(q.CurriculumEntry.ID.Eq(idUUID)).
				Preload(field.Associations).
				First()
			if err != nil {
				return err
			}

			curriculumCourseLevels, err = tx.CurriculumCourseLevel.
				Select(q.CurriculumCourseLevel.ALL).
				Where(q.CurriculumCourseLevel.CourseID.Eq(curriculumCourse.ID)).
				Find()
			if err != nil {
				return err
			}

			for _, courseLevel := range curriculumCourseLevels {
				courseLevelDTO := dto.CurriculumCourseLevels{
					ID:          courseLevel.ID.ToString(),
					Name:        courseLevel.Name,
					Description: courseLevel.Description,
					Title:       courseLevel.Title,
					IconID:      courseLevel.IconID.ToString(),
				}

				curriculumCourseLevelLessons, err = tx.CurriculumCourseLevelLesson.
					Select(q.CurriculumCourseLevelLesson.ALL).
					Where(q.CurriculumCourseLevelLesson.CourseLevelID.Eq(courseLevel.ID)).
					Find()
				if err != nil {
					return err
				}
				for _, curriculumCourseLevelLesson := range curriculumCourseLevelLessons {
					curriculumCourseLevelLessonDTO := dto.CurriculumCourseLevelLessons{
						ID:           curriculumCourseLevelLesson.ID.ToString(),
						LessonNumber: curriculumCourseLevelLesson.LessonNumber,
					}

					ppts, err = tx.CurriculumCourseLevelLessonResources.
						Select(q.CurriculumCourseLevelLessonResources.ALL).
						LeftJoin(q.CurriculumCourseLessonResourceType, q.CurriculumCourseLevelLessonResources.ResourseTypeID.EqCol(q.CurriculumCourseLessonResourceType.ID)).
						Where(q.CurriculumCourseLevelLessonResources.LessonID.Eq(curriculumCourseLevelLesson.ID)).
						Where(q.CurriculumCourseLessonResourceType.Name.Eq("presentation_notes")).
						Preload(field.Associations).
						Find()
					if err != nil {
						return err
					}
					for _, ppt := range ppts {
						pptDTO := dto.CurriculumCourseLevelLessonResources{
							ID:         ppt.ID.ToString(),
							ResourseID: ppt.Resourse.ID.ToString(),
							Name:       ppt.Resourse.FileNameUploaded,
						}
						curriculumCourseLevelLessonDTO.PresentationNotes = append(curriculumCourseLevelLessonDTO.PresentationNotes, pptDTO)
					}

					students, err = tx.CurriculumCourseLevelLessonResources.
						Select(q.CurriculumCourseLevelLessonResources.ALL).
						LeftJoin(q.CurriculumCourseLessonResourceType, q.CurriculumCourseLevelLessonResources.ResourseTypeID.EqCol(q.CurriculumCourseLessonResourceType.ID)).
						Where(q.CurriculumCourseLevelLessonResources.LessonID.Eq(curriculumCourseLevelLesson.ID)).
						Where(q.CurriculumCourseLessonResourceType.Name.Eq("student_notes")).
						Preload(field.Associations).
						Find()
					if err != nil {
						return err
					}

					for _, student := range students {
						studentDTO := dto.CurriculumCourseLevelLessonResources{
							ID:         student.ID.ToString(),
							ResourseID: student.Resourse.ID.ToString(),
							Name:       student.Resourse.FileNameUploaded,
						}
						curriculumCourseLevelLessonDTO.StudentNotes = append(curriculumCourseLevelLessonDTO.StudentNotes, studentDTO)
					}

					teachers, err = tx.CurriculumCourseLevelLessonResources.
						Select(q.CurriculumCourseLevelLessonResources.ALL).
						LeftJoin(q.CurriculumCourseLessonResourceType, q.CurriculumCourseLevelLessonResources.ResourseTypeID.EqCol(q.CurriculumCourseLessonResourceType.ID)).
						Where(q.CurriculumCourseLevelLessonResources.LessonID.Eq(curriculumCourseLevelLesson.ID)).
						Where(q.CurriculumCourseLessonResourceType.Name.Eq("teacher_notes")).
						Preload(field.Associations).
						Find()
					if err != nil {
						return err
					}
					for _, teacher := range teachers {
						teacherDTO := dto.CurriculumCourseLevelLessonResources{
							ID:         teacher.ID.ToString(),
							ResourseID: teacher.Resourse.ID.ToString(),
							Name:       teacher.Resourse.FileNameUploaded,
						}
						curriculumCourseLevelLessonDTO.TeacherNotes = append(curriculumCourseLevelLessonDTO.TeacherNotes, teacherDTO)
					}

					miscs, err = tx.CurriculumCourseLevelLessonResources.
						Select(q.CurriculumCourseLevelLessonResources.ALL).
						LeftJoin(q.CurriculumCourseLessonResourceType, q.CurriculumCourseLevelLessonResources.ResourseTypeID.EqCol(q.CurriculumCourseLessonResourceType.ID)).
						Where(q.CurriculumCourseLevelLessonResources.LessonID.Eq(curriculumCourseLevelLesson.ID)).
						Where(q.CurriculumCourseLessonResourceType.Name.Eq("misc_materials")).
						Preload(field.Associations).
						Find()
					if err != nil {
						return err
					}
					for _, misc := range miscs {
						miscDTO := dto.CurriculumCourseLevelLessonResources{
							ID:         misc.ID.ToString(),
							ResourseID: misc.Resourse.ID.ToString(),
							Name:       misc.Resourse.FileNameUploaded,
						}
						curriculumCourseLevelLessonDTO.MiscMaterials = append(curriculumCourseLevelLessonDTO.MiscMaterials, miscDTO)
					}

					courseLevelDTO.Lessons = append(courseLevelDTO.Lessons, curriculumCourseLevelLessonDTO)
				}
				returnForm.Levels = append(returnForm.Levels, courseLevelDTO)
			}

			return nil
		})

		if err != nil {
			ctx.StopWithError(iris.StatusNotFound, err)
		} else {
			returnForm.ID = (*curriculumEntry).ID.ToString()
			returnForm.Description = (*curriculumEntry).Description
			returnForm.IconID = (*curriculumEntry).IconID.ToString()
			if (*curriculumEntry).ParentID != nil {
				returnForm.ParentID = (*(*curriculumEntry).ParentID).ToString()
			}

			returnForm.CourseID = (*curriculumCourse).ID.ToString()
			returnForm.CurriculumPlanID = (*curriculumCourse).CurriculumPlanID.ToString()
			returnForm.CurriculumPlanFileName = (*curriculumCourse).CurriculumPlan.FileNameUploaded

			for _, youtube := range curriculumCourseYoutubeVideoEntries {
				returnForm.YoutubeVideoEntries = append(returnForm.YoutubeVideoEntries, dto.CurriculumCourseYoutubeVideoEntries{
					ID:  youtube.ID.ToString(),
					URL: youtube.URL,
				})
			}

			for _, blog := range curriculumCourseBlogEntries {
				returnForm.BlogEntries = append(returnForm.BlogEntries, dto.CurriculumCourseBlogEntries{
					ID:          blog.ID.ToString(),
					ExternalURL: blog.ExternalURL,
					Title:       blog.Title,
				})
			}

			ctx.JSON(returnForm)
		}
	}
}

func CreateOrUpdateCurriculumCourse(s3 *utils.StemexS3Client, dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		form, err := MapRequestToCurriculumCourseForm(ctx.Request())
		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
			return
		}

		returnForm, err := bo.CreateOrUpdateCurriculumCourse(form, s3, query.Use(dbInstance))
		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
		} else {
			ctx.JSON(returnForm)
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
// 			ctx.StopWithStatus(iris.StatusForbidden)
// 			return
// 		}

// 		parentIDUUIDEx, err := model.ValidUUIDExFromIDString(parentID)
// 		if err != nil {
// 			ctx.StopWithStatus(iris.StatusNotFound)
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

func MapRequestToCurriculumCourseTypeForm(req *http.Request) (*dto.CurriculumCourseTypeForm, error) {
	var form dto.CurriculumCourseTypeForm
	// Parse request data.
	curriculumEntryFormData, err := forms.Parse(req)
	if err != nil {
		// ctx.StopWithError(iris.StatusInternalServerError, errParse)
		// return
		return nil, err
	}

	val := curriculumEntryFormData.Validator()
	val.Require("description")

	if !val.HasErrors() {
		form.ID = curriculumEntryFormData.Get("id")
		form.IconID = curriculumEntryFormData.Get("icon_id")
		form.IconFile = curriculumEntryFormData.GetFile("icon_file")
		form.ParentID = curriculumEntryFormData.Get("parent_id")
		form.Description = curriculumEntryFormData.Get("description")
	}
	return &form, nil
}

func MapRequestToCurriculumCourseForm(req *http.Request) (*dto.CurriculumCourseForm, error) {
	var form dto.CurriculumCourseForm
	// Parse request data.
	curriculumEntryFormData, err := forms.Parse(req)
	if err != nil {
		// ctx.StopWithError(iris.StatusInternalServerError, errParse)
		// return
		return nil, err
	}

	val := curriculumEntryFormData.Validator()
	val.Require("description")

	if !val.HasErrors() {
		form.ID = curriculumEntryFormData.Get("id")
		form.IconID = curriculumEntryFormData.Get("icon_id")
		form.IconFile = curriculumEntryFormData.GetFile("icon_file")
		form.ParentID = curriculumEntryFormData.Get("parent_id")
		form.Description = curriculumEntryFormData.Get("description")
		form.CourseID = curriculumEntryFormData.Get("course_id")
		form.CurriculumPlanID = curriculumEntryFormData.Get("curriculum_plan_id")
		form.CurriculumPlanFile = curriculumEntryFormData.GetFile("curriculum_plan_file")
		form.CurriculumPlanFileName = curriculumEntryFormData.Get("curriculum_plan_file_name")

		var youtubeVideoEntriesBaseKey = "youtube_video_entries[%d]"
		MapFormArray(curriculumEntryFormData, func() *dto.CurriculumCourseYoutubeVideoEntries { return &dto.CurriculumCourseYoutubeVideoEntries{} },
			[]datatypes.Pair[string, func(*dto.CurriculumCourseYoutubeVideoEntries, string)]{
				{
					First: youtubeVideoEntriesBaseKey + ".id",
					Second: func(dto *dto.CurriculumCourseYoutubeVideoEntries, s string) {
						dto.ID = s
					},
				},
				{
					First: youtubeVideoEntriesBaseKey + ".url",
					Second: func(dto *dto.CurriculumCourseYoutubeVideoEntries, s string) {
						dto.URL = s
					},
				}},
			[]datatypes.Pair[string, func(*dto.CurriculumCourseYoutubeVideoEntries, *multipart.FileHeader)]{},
			func(n *dto.CurriculumCourseYoutubeVideoEntries) {
				form.YoutubeVideoEntries = append(form.YoutubeVideoEntries, *n)
			},
		)

		var blogEntriesBaseKey = "blog_entries[%d]"

		MapFormArray(curriculumEntryFormData, func() *dto.CurriculumCourseBlogEntries { return &dto.CurriculumCourseBlogEntries{} },
			[]datatypes.Pair[string, func(*dto.CurriculumCourseBlogEntries, string)]{
				{
					First: blogEntriesBaseKey + ".id",
					Second: func(dto *dto.CurriculumCourseBlogEntries, s string) {
						dto.ID = s
					},
				},
				{
					First: blogEntriesBaseKey + ".title",
					Second: func(dto *dto.CurriculumCourseBlogEntries, s string) {
						dto.Title = s
					},
				},
				{
					First: blogEntriesBaseKey + ".external_url",
					Second: func(dto *dto.CurriculumCourseBlogEntries, s string) {
						dto.ExternalURL = s
					},
				},
			},
			[]datatypes.Pair[string, func(*dto.CurriculumCourseBlogEntries, *multipart.FileHeader)]{},
			func(n *dto.CurriculumCourseBlogEntries) {
				form.BlogEntries = append(form.BlogEntries, *n)
			},
		)

		var i = 0

		for {
			var levelsIDKey = fmt.Sprintf(`levels[%d].id`, i)
			var levelsNameKey = fmt.Sprintf(`levels[%d].name`, i)
			var levelsIconIDKey = fmt.Sprintf(`levels[%d].icon_id`, i)
			var levelsIconFileKey = fmt.Sprintf(`levels[%d].icon_file`, i)
			var levelsTitleKey = fmt.Sprintf(`levels[%d].title`, i)
			var levelsDescriptionKey = fmt.Sprintf(`levels[%d].description`, i)
			levelsIDKeyExists := curriculumEntryFormData.KeyExists(levelsIDKey)
			levelsNameKeyExists := curriculumEntryFormData.KeyExists(levelsNameKey)
			levelsTitleKeyExists := curriculumEntryFormData.KeyExists(levelsTitleKey)
			levelsIconIDKeyExists := curriculumEntryFormData.KeyExists(levelsIconIDKey)
			levelsIconFileKeyExists := curriculumEntryFormData.KeyExists(levelsIconFileKey)
			levelsDescriptionKeyExists := curriculumEntryFormData.KeyExists(levelsDescriptionKey)

			if levelsIDKeyExists || levelsNameKeyExists || levelsTitleKeyExists || levelsIconIDKeyExists || levelsIconFileKeyExists || levelsDescriptionKeyExists {
				level := dto.CurriculumCourseLevels{
					ID:          curriculumEntryFormData.Get(levelsIDKey),
					Name:        curriculumEntryFormData.Get(levelsNameKey),
					Title:       curriculumEntryFormData.Get(levelsTitleKey),
					IconFile:    curriculumEntryFormData.GetFile(levelsIconFileKey),
					IconID:      curriculumEntryFormData.Get(levelsIconIDKey),
					Description: curriculumEntryFormData.Get(levelsDescriptionKey),
				}

				var j = 0
				for {
					var lessonsArrayKey = fmt.Sprintf(`levels[%d].lessons[%d]`, i, j)

					mapDifferentTypesOfResources := func(baseKey string, callback func(dto *dto.CurriculumCourseLevelLessonResources)) {
						MapFormArray(curriculumEntryFormData, func() *dto.CurriculumCourseLevelLessonResources { return &dto.CurriculumCourseLevelLessonResources{} },
							[]datatypes.Pair[string, func(*dto.CurriculumCourseLevelLessonResources, string)]{{
								First: baseKey + ".id",
								Second: func(ccllr *dto.CurriculumCourseLevelLessonResources, s string) {
									ccllr.Name = s
								},
							}},
							[]datatypes.Pair[string, func(*dto.CurriculumCourseLevelLessonResources, *multipart.FileHeader)]{{
								First: baseKey + ".file",
								Second: func(ccllr *dto.CurriculumCourseLevelLessonResources, b *multipart.FileHeader) {
									ccllr.File = b
								},
							}},
							callback,
						)
					}

					presentationNotes := []dto.CurriculumCourseLevelLessonResources{}
					studentNotes := []dto.CurriculumCourseLevelLessonResources{}
					teacherNotes := []dto.CurriculumCourseLevelLessonResources{}
					miscMaterials := []dto.CurriculumCourseLevelLessonResources{}

					mapDifferentTypesOfResources(lessonsArrayKey+".presentation_notes[%d]", func(dto *dto.CurriculumCourseLevelLessonResources) {
						presentationNotes = append(presentationNotes, *dto)
					})

					mapDifferentTypesOfResources(lessonsArrayKey+".student_notes[%d]", func(dto *dto.CurriculumCourseLevelLessonResources) {
						studentNotes = append(studentNotes, *dto)
					})

					mapDifferentTypesOfResources(lessonsArrayKey+".teacher_notes[%d]", func(dto *dto.CurriculumCourseLevelLessonResources) {
						teacherNotes = append(teacherNotes, *dto)
					})

					mapDifferentTypesOfResources(lessonsArrayKey+".misc_materials[%d]", func(dto *dto.CurriculumCourseLevelLessonResources) {
						miscMaterials = append(miscMaterials, *dto)
					})

					if len(presentationNotes) > 0 || len(studentNotes) > 0 || len(teacherNotes) > 0 || len(miscMaterials) > 0 {
						level.Lessons = append(level.Lessons, dto.CurriculumCourseLevelLessons{
							PresentationNotes: presentationNotes,
							StudentNotes:      studentNotes,
							TeacherNotes:      teacherNotes,
							MiscMaterials:     miscMaterials,
						})
						j = j + 1
					} else {
						break
					}
				}

				form.Levels = append(form.Levels, level)
				i = i + 1
			} else {
				break
			}
		}
	}
	return &form, nil
}

func MapFormArray[N any](data *forms.Data, onNewItem func() *N, pairsForString []datatypes.Pair[string, func(*N, string)], pairsForFileBytes []datatypes.Pair[string, func(*N, *multipart.FileHeader)], onOK func(n *N)) {
	var k = 0
	var n *N = nil

	for {
		keysForString := lo.Map(pairsForString, func(pair datatypes.Pair[string, func(*N, string)], index int) string {
			return fmt.Sprintf(pair.First, k)
		})

		fmt.Println(keysForString)

		keysForFileBytes := lo.Map(pairsForFileBytes, func(pair datatypes.Pair[string, func(*N, *multipart.FileHeader)], index int) string {
			return fmt.Sprintf(pair.First, k)
		})

		if lo.SomeBy(keysForString, func(key string) bool {
			return data.KeyExists(key)
		}) || lo.SomeBy(keysForFileBytes, func(key string) bool {
			return data.FileExists(key)
		}) {
			n = onNewItem()

			lo.ForEach(pairsForString, func(pair datatypes.Pair[string, func(*N, string)], index int) {
				key := keysForString[index]
				content := data.Get(key)
				pair.Second(n, content)
			})

			lo.ForEach(pairsForFileBytes, func(pair datatypes.Pair[string, func(*N, *multipart.FileHeader)], index int) {
				key := keysForFileBytes[index]
				file := data.GetFile(key)
				pair.Second(n, file)
			})

			onOK(n)

			k = k + 1
		} else {
			break
		}
	}
}
