// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"github.com/dirkarnez/stemexapi/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newCurriculumCourseLevelLessonResources(db *gorm.DB, opts ...gen.DOOption) curriculumCourseLevelLessonResources {
	_curriculumCourseLevelLessonResources := curriculumCourseLevelLessonResources{}

	_curriculumCourseLevelLessonResources.curriculumCourseLevelLessonResourcesDo.UseDB(db, opts...)
	_curriculumCourseLevelLessonResources.curriculumCourseLevelLessonResourcesDo.UseModel(&model.CurriculumCourseLevelLessonResources{})

	tableName := _curriculumCourseLevelLessonResources.curriculumCourseLevelLessonResourcesDo.TableName()
	_curriculumCourseLevelLessonResources.ALL = field.NewAsterisk(tableName)
	_curriculumCourseLevelLessonResources.ID = field.NewField(tableName, "id")
	_curriculumCourseLevelLessonResources.CreatedAt = field.NewTime(tableName, "created_at")
	_curriculumCourseLevelLessonResources.UpdatedAt = field.NewTime(tableName, "updated_at")
	_curriculumCourseLevelLessonResources.DeletedAt = field.NewField(tableName, "deleted_at")
	_curriculumCourseLevelLessonResources.LessonID = field.NewField(tableName, "lesson_id")
	_curriculumCourseLevelLessonResources.ResourseTypeID = field.NewField(tableName, "resourse_type_id")
	_curriculumCourseLevelLessonResources.ResourseID = field.NewField(tableName, "resourse_id")
	_curriculumCourseLevelLessonResources.Lesson = curriculumCourseLevelLessonResourcesBelongsToLesson{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Lesson", "model.CurriculumCourseLevelLesson"),
		CourseLevel: struct {
			field.RelationField
			Icon struct {
				field.RelationField
			}
			Course struct {
				field.RelationField
				Entry struct {
					field.RelationField
					Icon struct {
						field.RelationField
					}
				}
				CurriculumPlan struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("Lesson.CourseLevel", "model.CurriculumCourseLevel"),
			Icon: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Lesson.CourseLevel.Icon", "model.File"),
			},
			Course: struct {
				field.RelationField
				Entry struct {
					field.RelationField
					Icon struct {
						field.RelationField
					}
				}
				CurriculumPlan struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Lesson.CourseLevel.Course", "model.CurriculumCourse"),
				Entry: struct {
					field.RelationField
					Icon struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("Lesson.CourseLevel.Course.Entry", "model.CurriculumEntry"),
					Icon: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Lesson.CourseLevel.Course.Entry.Icon", "model.File"),
					},
				},
				CurriculumPlan: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Lesson.CourseLevel.Course.CurriculumPlan", "model.File"),
				},
			},
		},
	}

	_curriculumCourseLevelLessonResources.ResourseType = curriculumCourseLevelLessonResourcesBelongsToResourseType{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("ResourseType", "model.CurriculumCourseLessonResourceType"),
	}

	_curriculumCourseLevelLessonResources.Resourse = curriculumCourseLevelLessonResourcesBelongsToResourse{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Resourse", "model.File"),
	}

	_curriculumCourseLevelLessonResources.fillFieldMap()

	return _curriculumCourseLevelLessonResources
}

type curriculumCourseLevelLessonResources struct {
	curriculumCourseLevelLessonResourcesDo

	ALL            field.Asterisk
	ID             field.Field
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field
	LessonID       field.Field
	ResourseTypeID field.Field
	ResourseID     field.Field
	Lesson         curriculumCourseLevelLessonResourcesBelongsToLesson

	ResourseType curriculumCourseLevelLessonResourcesBelongsToResourseType

	Resourse curriculumCourseLevelLessonResourcesBelongsToResourse

	fieldMap map[string]field.Expr
}

func (c curriculumCourseLevelLessonResources) Table(newTableName string) *curriculumCourseLevelLessonResources {
	c.curriculumCourseLevelLessonResourcesDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c curriculumCourseLevelLessonResources) As(alias string) *curriculumCourseLevelLessonResources {
	c.curriculumCourseLevelLessonResourcesDo.DO = *(c.curriculumCourseLevelLessonResourcesDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *curriculumCourseLevelLessonResources) updateTableName(table string) *curriculumCourseLevelLessonResources {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewField(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.LessonID = field.NewField(table, "lesson_id")
	c.ResourseTypeID = field.NewField(table, "resourse_type_id")
	c.ResourseID = field.NewField(table, "resourse_id")

	c.fillFieldMap()

	return c
}

func (c *curriculumCourseLevelLessonResources) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *curriculumCourseLevelLessonResources) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 10)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["lesson_id"] = c.LessonID
	c.fieldMap["resourse_type_id"] = c.ResourseTypeID
	c.fieldMap["resourse_id"] = c.ResourseID

}

func (c curriculumCourseLevelLessonResources) clone(db *gorm.DB) curriculumCourseLevelLessonResources {
	c.curriculumCourseLevelLessonResourcesDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c curriculumCourseLevelLessonResources) replaceDB(db *gorm.DB) curriculumCourseLevelLessonResources {
	c.curriculumCourseLevelLessonResourcesDo.ReplaceDB(db)
	return c
}

type curriculumCourseLevelLessonResourcesBelongsToLesson struct {
	db *gorm.DB

	field.RelationField

	CourseLevel struct {
		field.RelationField
		Icon struct {
			field.RelationField
		}
		Course struct {
			field.RelationField
			Entry struct {
				field.RelationField
				Icon struct {
					field.RelationField
				}
			}
			CurriculumPlan struct {
				field.RelationField
			}
		}
	}
}

func (a curriculumCourseLevelLessonResourcesBelongsToLesson) Where(conds ...field.Expr) *curriculumCourseLevelLessonResourcesBelongsToLesson {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a curriculumCourseLevelLessonResourcesBelongsToLesson) WithContext(ctx context.Context) *curriculumCourseLevelLessonResourcesBelongsToLesson {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a curriculumCourseLevelLessonResourcesBelongsToLesson) Session(session *gorm.Session) *curriculumCourseLevelLessonResourcesBelongsToLesson {
	a.db = a.db.Session(session)
	return &a
}

func (a curriculumCourseLevelLessonResourcesBelongsToLesson) Model(m *model.CurriculumCourseLevelLessonResources) *curriculumCourseLevelLessonResourcesBelongsToLessonTx {
	return &curriculumCourseLevelLessonResourcesBelongsToLessonTx{a.db.Model(m).Association(a.Name())}
}

type curriculumCourseLevelLessonResourcesBelongsToLessonTx struct{ tx *gorm.Association }

func (a curriculumCourseLevelLessonResourcesBelongsToLessonTx) Find() (result *model.CurriculumCourseLevelLesson, err error) {
	return result, a.tx.Find(&result)
}

func (a curriculumCourseLevelLessonResourcesBelongsToLessonTx) Append(values ...*model.CurriculumCourseLevelLesson) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a curriculumCourseLevelLessonResourcesBelongsToLessonTx) Replace(values ...*model.CurriculumCourseLevelLesson) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a curriculumCourseLevelLessonResourcesBelongsToLessonTx) Delete(values ...*model.CurriculumCourseLevelLesson) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a curriculumCourseLevelLessonResourcesBelongsToLessonTx) Clear() error {
	return a.tx.Clear()
}

func (a curriculumCourseLevelLessonResourcesBelongsToLessonTx) Count() int64 {
	return a.tx.Count()
}

type curriculumCourseLevelLessonResourcesBelongsToResourseType struct {
	db *gorm.DB

	field.RelationField
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourseType) Where(conds ...field.Expr) *curriculumCourseLevelLessonResourcesBelongsToResourseType {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourseType) WithContext(ctx context.Context) *curriculumCourseLevelLessonResourcesBelongsToResourseType {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourseType) Session(session *gorm.Session) *curriculumCourseLevelLessonResourcesBelongsToResourseType {
	a.db = a.db.Session(session)
	return &a
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourseType) Model(m *model.CurriculumCourseLevelLessonResources) *curriculumCourseLevelLessonResourcesBelongsToResourseTypeTx {
	return &curriculumCourseLevelLessonResourcesBelongsToResourseTypeTx{a.db.Model(m).Association(a.Name())}
}

type curriculumCourseLevelLessonResourcesBelongsToResourseTypeTx struct{ tx *gorm.Association }

func (a curriculumCourseLevelLessonResourcesBelongsToResourseTypeTx) Find() (result *model.CurriculumCourseLessonResourceType, err error) {
	return result, a.tx.Find(&result)
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourseTypeTx) Append(values ...*model.CurriculumCourseLessonResourceType) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourseTypeTx) Replace(values ...*model.CurriculumCourseLessonResourceType) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourseTypeTx) Delete(values ...*model.CurriculumCourseLessonResourceType) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourseTypeTx) Clear() error {
	return a.tx.Clear()
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourseTypeTx) Count() int64 {
	return a.tx.Count()
}

type curriculumCourseLevelLessonResourcesBelongsToResourse struct {
	db *gorm.DB

	field.RelationField
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourse) Where(conds ...field.Expr) *curriculumCourseLevelLessonResourcesBelongsToResourse {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourse) WithContext(ctx context.Context) *curriculumCourseLevelLessonResourcesBelongsToResourse {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourse) Session(session *gorm.Session) *curriculumCourseLevelLessonResourcesBelongsToResourse {
	a.db = a.db.Session(session)
	return &a
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourse) Model(m *model.CurriculumCourseLevelLessonResources) *curriculumCourseLevelLessonResourcesBelongsToResourseTx {
	return &curriculumCourseLevelLessonResourcesBelongsToResourseTx{a.db.Model(m).Association(a.Name())}
}

type curriculumCourseLevelLessonResourcesBelongsToResourseTx struct{ tx *gorm.Association }

func (a curriculumCourseLevelLessonResourcesBelongsToResourseTx) Find() (result *model.File, err error) {
	return result, a.tx.Find(&result)
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourseTx) Append(values ...*model.File) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourseTx) Replace(values ...*model.File) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourseTx) Delete(values ...*model.File) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourseTx) Clear() error {
	return a.tx.Clear()
}

func (a curriculumCourseLevelLessonResourcesBelongsToResourseTx) Count() int64 {
	return a.tx.Count()
}

type curriculumCourseLevelLessonResourcesDo struct{ gen.DO }

type ICurriculumCourseLevelLessonResourcesDo interface {
	gen.SubQuery
	Debug() ICurriculumCourseLevelLessonResourcesDo
	WithContext(ctx context.Context) ICurriculumCourseLevelLessonResourcesDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICurriculumCourseLevelLessonResourcesDo
	WriteDB() ICurriculumCourseLevelLessonResourcesDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICurriculumCourseLevelLessonResourcesDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICurriculumCourseLevelLessonResourcesDo
	Not(conds ...gen.Condition) ICurriculumCourseLevelLessonResourcesDo
	Or(conds ...gen.Condition) ICurriculumCourseLevelLessonResourcesDo
	Select(conds ...field.Expr) ICurriculumCourseLevelLessonResourcesDo
	Where(conds ...gen.Condition) ICurriculumCourseLevelLessonResourcesDo
	Order(conds ...field.Expr) ICurriculumCourseLevelLessonResourcesDo
	Distinct(cols ...field.Expr) ICurriculumCourseLevelLessonResourcesDo
	Omit(cols ...field.Expr) ICurriculumCourseLevelLessonResourcesDo
	Join(table schema.Tabler, on ...field.Expr) ICurriculumCourseLevelLessonResourcesDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseLevelLessonResourcesDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseLevelLessonResourcesDo
	Group(cols ...field.Expr) ICurriculumCourseLevelLessonResourcesDo
	Having(conds ...gen.Condition) ICurriculumCourseLevelLessonResourcesDo
	Limit(limit int) ICurriculumCourseLevelLessonResourcesDo
	Offset(offset int) ICurriculumCourseLevelLessonResourcesDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICurriculumCourseLevelLessonResourcesDo
	Unscoped() ICurriculumCourseLevelLessonResourcesDo
	Create(values ...*model.CurriculumCourseLevelLessonResources) error
	CreateInBatches(values []*model.CurriculumCourseLevelLessonResources, batchSize int) error
	Save(values ...*model.CurriculumCourseLevelLessonResources) error
	First() (*model.CurriculumCourseLevelLessonResources, error)
	Take() (*model.CurriculumCourseLevelLessonResources, error)
	Last() (*model.CurriculumCourseLevelLessonResources, error)
	Find() ([]*model.CurriculumCourseLevelLessonResources, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CurriculumCourseLevelLessonResources, err error)
	FindInBatches(result *[]*model.CurriculumCourseLevelLessonResources, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CurriculumCourseLevelLessonResources) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICurriculumCourseLevelLessonResourcesDo
	Assign(attrs ...field.AssignExpr) ICurriculumCourseLevelLessonResourcesDo
	Joins(fields ...field.RelationField) ICurriculumCourseLevelLessonResourcesDo
	Preload(fields ...field.RelationField) ICurriculumCourseLevelLessonResourcesDo
	FirstOrInit() (*model.CurriculumCourseLevelLessonResources, error)
	FirstOrCreate() (*model.CurriculumCourseLevelLessonResources, error)
	FindByPage(offset int, limit int) (result []*model.CurriculumCourseLevelLessonResources, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICurriculumCourseLevelLessonResourcesDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c curriculumCourseLevelLessonResourcesDo) Debug() ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Debug())
}

func (c curriculumCourseLevelLessonResourcesDo) WithContext(ctx context.Context) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c curriculumCourseLevelLessonResourcesDo) ReadDB() ICurriculumCourseLevelLessonResourcesDo {
	return c.Clauses(dbresolver.Read)
}

func (c curriculumCourseLevelLessonResourcesDo) WriteDB() ICurriculumCourseLevelLessonResourcesDo {
	return c.Clauses(dbresolver.Write)
}

func (c curriculumCourseLevelLessonResourcesDo) Session(config *gorm.Session) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Session(config))
}

func (c curriculumCourseLevelLessonResourcesDo) Clauses(conds ...clause.Expression) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c curriculumCourseLevelLessonResourcesDo) Returning(value interface{}, columns ...string) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c curriculumCourseLevelLessonResourcesDo) Not(conds ...gen.Condition) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c curriculumCourseLevelLessonResourcesDo) Or(conds ...gen.Condition) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c curriculumCourseLevelLessonResourcesDo) Select(conds ...field.Expr) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c curriculumCourseLevelLessonResourcesDo) Where(conds ...gen.Condition) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c curriculumCourseLevelLessonResourcesDo) Order(conds ...field.Expr) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c curriculumCourseLevelLessonResourcesDo) Distinct(cols ...field.Expr) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c curriculumCourseLevelLessonResourcesDo) Omit(cols ...field.Expr) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c curriculumCourseLevelLessonResourcesDo) Join(table schema.Tabler, on ...field.Expr) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c curriculumCourseLevelLessonResourcesDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c curriculumCourseLevelLessonResourcesDo) RightJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c curriculumCourseLevelLessonResourcesDo) Group(cols ...field.Expr) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c curriculumCourseLevelLessonResourcesDo) Having(conds ...gen.Condition) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c curriculumCourseLevelLessonResourcesDo) Limit(limit int) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c curriculumCourseLevelLessonResourcesDo) Offset(offset int) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c curriculumCourseLevelLessonResourcesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c curriculumCourseLevelLessonResourcesDo) Unscoped() ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Unscoped())
}

func (c curriculumCourseLevelLessonResourcesDo) Create(values ...*model.CurriculumCourseLevelLessonResources) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c curriculumCourseLevelLessonResourcesDo) CreateInBatches(values []*model.CurriculumCourseLevelLessonResources, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c curriculumCourseLevelLessonResourcesDo) Save(values ...*model.CurriculumCourseLevelLessonResources) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c curriculumCourseLevelLessonResourcesDo) First() (*model.CurriculumCourseLevelLessonResources, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseLevelLessonResources), nil
	}
}

func (c curriculumCourseLevelLessonResourcesDo) Take() (*model.CurriculumCourseLevelLessonResources, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseLevelLessonResources), nil
	}
}

func (c curriculumCourseLevelLessonResourcesDo) Last() (*model.CurriculumCourseLevelLessonResources, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseLevelLessonResources), nil
	}
}

func (c curriculumCourseLevelLessonResourcesDo) Find() ([]*model.CurriculumCourseLevelLessonResources, error) {
	result, err := c.DO.Find()
	return result.([]*model.CurriculumCourseLevelLessonResources), err
}

func (c curriculumCourseLevelLessonResourcesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CurriculumCourseLevelLessonResources, err error) {
	buf := make([]*model.CurriculumCourseLevelLessonResources, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c curriculumCourseLevelLessonResourcesDo) FindInBatches(result *[]*model.CurriculumCourseLevelLessonResources, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c curriculumCourseLevelLessonResourcesDo) Attrs(attrs ...field.AssignExpr) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c curriculumCourseLevelLessonResourcesDo) Assign(attrs ...field.AssignExpr) ICurriculumCourseLevelLessonResourcesDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c curriculumCourseLevelLessonResourcesDo) Joins(fields ...field.RelationField) ICurriculumCourseLevelLessonResourcesDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c curriculumCourseLevelLessonResourcesDo) Preload(fields ...field.RelationField) ICurriculumCourseLevelLessonResourcesDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c curriculumCourseLevelLessonResourcesDo) FirstOrInit() (*model.CurriculumCourseLevelLessonResources, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseLevelLessonResources), nil
	}
}

func (c curriculumCourseLevelLessonResourcesDo) FirstOrCreate() (*model.CurriculumCourseLevelLessonResources, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseLevelLessonResources), nil
	}
}

func (c curriculumCourseLevelLessonResourcesDo) FindByPage(offset int, limit int) (result []*model.CurriculumCourseLevelLessonResources, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c curriculumCourseLevelLessonResourcesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c curriculumCourseLevelLessonResourcesDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c curriculumCourseLevelLessonResourcesDo) Delete(models ...*model.CurriculumCourseLevelLessonResources) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *curriculumCourseLevelLessonResourcesDo) withDO(do gen.Dao) *curriculumCourseLevelLessonResourcesDo {
	c.DO = *do.(*gen.DO)
	return c
}
