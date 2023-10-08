package api

import (
	"github.com/dirkarnez/stemexapi/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
)

func GetCurriculum(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		var curriculumEntryList []model.CurriculumEntry
		if err := dbInstance.Raw(`
			WITH RECURSIVE curriculum_entries_nested AS (
				SELECT id, description, parent_id FROM curriculum_entries WHERE description = 'Microbit'
				UNION
				SELECT curriculum_entries.id, curriculum_entries.description, curriculum_entries.parent_id FROM curriculum_entries JOIN curriculum_entries_nested ON curriculum_entries_nested.id = curriculum_entries.parent_id
			)
			SELECT * FROM curriculum_entries_nested
		`).
			Scan(&curriculumEntryList).Error; err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		} else {
			ctx.JSON(curriculumEntryList)
		}
	}
}
