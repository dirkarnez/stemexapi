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

func newCurriculumCourseLevelLesson(db *gorm.DB, opts ...gen.DOOption) curriculumCourseLevelLesson {
	_curriculumCourseLevelLesson := curriculumCourseLevelLesson{}

	_curriculumCourseLevelLesson.curriculumCourseLevelLessonDo.UseDB(db, opts...)
	_curriculumCourseLevelLesson.curriculumCourseLevelLessonDo.UseModel(&model.CurriculumCourseLevelLesson{})

	tableName := _curriculumCourseLevelLesson.curriculumCourseLevelLessonDo.TableName()
	_curriculumCourseLevelLesson.ALL = field.NewAsterisk(tableName)
	_curriculumCourseLevelLesson.ID = field.NewField(tableName, "id")
	_curriculumCourseLevelLesson.CreatedAt = field.NewTime(tableName, "created_at")
	_curriculumCourseLevelLesson.UpdatedAt = field.NewTime(tableName, "updated_at")
	_curriculumCourseLevelLesson.DeletedAt = field.NewField(tableName, "deleted_at")
	_curriculumCourseLevelLesson.LessonNumber = field.NewUint64(tableName, "lesson_number")
	_curriculumCourseLevelLesson.CourseLevelID = field.NewField(tableName, "course_level_id")
	_curriculumCourseLevelLesson.CourseLevel = curriculumCourseLevelLessonBelongsToCourseLevel{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("CourseLevel", "model.CurriculumCourseLevel"),
		Icon: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("CourseLevel.Icon", "model.File"),
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
			RelationField: field.NewRelation("CourseLevel.Course", "model.CurriculumCourse"),
			Entry: struct {
				field.RelationField
				Icon struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("CourseLevel.Course.Entry", "model.CurriculumEntry"),
				Icon: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("CourseLevel.Course.Entry.Icon", "model.File"),
				},
			},
			CurriculumPlan: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("CourseLevel.Course.CurriculumPlan", "model.File"),
			},
		},
	}

	_curriculumCourseLevelLesson.fillFieldMap()

	return _curriculumCourseLevelLesson
}

type curriculumCourseLevelLesson struct {
	curriculumCourseLevelLessonDo

	ALL           field.Asterisk
	ID            field.Field
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	LessonNumber  field.Uint64
	CourseLevelID field.Field
	CourseLevel   curriculumCourseLevelLessonBelongsToCourseLevel

	fieldMap map[string]field.Expr
}

func (c curriculumCourseLevelLesson) Table(newTableName string) *curriculumCourseLevelLesson {
	c.curriculumCourseLevelLessonDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c curriculumCourseLevelLesson) As(alias string) *curriculumCourseLevelLesson {
	c.curriculumCourseLevelLessonDo.DO = *(c.curriculumCourseLevelLessonDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *curriculumCourseLevelLesson) updateTableName(table string) *curriculumCourseLevelLesson {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewField(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.LessonNumber = field.NewUint64(table, "lesson_number")
	c.CourseLevelID = field.NewField(table, "course_level_id")

	c.fillFieldMap()

	return c
}

func (c *curriculumCourseLevelLesson) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *curriculumCourseLevelLesson) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["lesson_number"] = c.LessonNumber
	c.fieldMap["course_level_id"] = c.CourseLevelID

}

func (c curriculumCourseLevelLesson) clone(db *gorm.DB) curriculumCourseLevelLesson {
	c.curriculumCourseLevelLessonDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c curriculumCourseLevelLesson) replaceDB(db *gorm.DB) curriculumCourseLevelLesson {
	c.curriculumCourseLevelLessonDo.ReplaceDB(db)
	return c
}

type curriculumCourseLevelLessonBelongsToCourseLevel struct {
	db *gorm.DB

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

func (a curriculumCourseLevelLessonBelongsToCourseLevel) Where(conds ...field.Expr) *curriculumCourseLevelLessonBelongsToCourseLevel {
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

func (a curriculumCourseLevelLessonBelongsToCourseLevel) WithContext(ctx context.Context) *curriculumCourseLevelLessonBelongsToCourseLevel {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a curriculumCourseLevelLessonBelongsToCourseLevel) Session(session *gorm.Session) *curriculumCourseLevelLessonBelongsToCourseLevel {
	a.db = a.db.Session(session)
	return &a
}

func (a curriculumCourseLevelLessonBelongsToCourseLevel) Model(m *model.CurriculumCourseLevelLesson) *curriculumCourseLevelLessonBelongsToCourseLevelTx {
	return &curriculumCourseLevelLessonBelongsToCourseLevelTx{a.db.Model(m).Association(a.Name())}
}

type curriculumCourseLevelLessonBelongsToCourseLevelTx struct{ tx *gorm.Association }

func (a curriculumCourseLevelLessonBelongsToCourseLevelTx) Find() (result *model.CurriculumCourseLevel, err error) {
	return result, a.tx.Find(&result)
}

func (a curriculumCourseLevelLessonBelongsToCourseLevelTx) Append(values ...*model.CurriculumCourseLevel) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a curriculumCourseLevelLessonBelongsToCourseLevelTx) Replace(values ...*model.CurriculumCourseLevel) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a curriculumCourseLevelLessonBelongsToCourseLevelTx) Delete(values ...*model.CurriculumCourseLevel) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a curriculumCourseLevelLessonBelongsToCourseLevelTx) Clear() error {
	return a.tx.Clear()
}

func (a curriculumCourseLevelLessonBelongsToCourseLevelTx) Count() int64 {
	return a.tx.Count()
}

type curriculumCourseLevelLessonDo struct{ gen.DO }

type ICurriculumCourseLevelLessonDo interface {
	gen.SubQuery
	Debug() ICurriculumCourseLevelLessonDo
	WithContext(ctx context.Context) ICurriculumCourseLevelLessonDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICurriculumCourseLevelLessonDo
	WriteDB() ICurriculumCourseLevelLessonDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICurriculumCourseLevelLessonDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICurriculumCourseLevelLessonDo
	Not(conds ...gen.Condition) ICurriculumCourseLevelLessonDo
	Or(conds ...gen.Condition) ICurriculumCourseLevelLessonDo
	Select(conds ...field.Expr) ICurriculumCourseLevelLessonDo
	Where(conds ...gen.Condition) ICurriculumCourseLevelLessonDo
	Order(conds ...field.Expr) ICurriculumCourseLevelLessonDo
	Distinct(cols ...field.Expr) ICurriculumCourseLevelLessonDo
	Omit(cols ...field.Expr) ICurriculumCourseLevelLessonDo
	Join(table schema.Tabler, on ...field.Expr) ICurriculumCourseLevelLessonDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseLevelLessonDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseLevelLessonDo
	Group(cols ...field.Expr) ICurriculumCourseLevelLessonDo
	Having(conds ...gen.Condition) ICurriculumCourseLevelLessonDo
	Limit(limit int) ICurriculumCourseLevelLessonDo
	Offset(offset int) ICurriculumCourseLevelLessonDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICurriculumCourseLevelLessonDo
	Unscoped() ICurriculumCourseLevelLessonDo
	Create(values ...*model.CurriculumCourseLevelLesson) error
	CreateInBatches(values []*model.CurriculumCourseLevelLesson, batchSize int) error
	Save(values ...*model.CurriculumCourseLevelLesson) error
	First() (*model.CurriculumCourseLevelLesson, error)
	Take() (*model.CurriculumCourseLevelLesson, error)
	Last() (*model.CurriculumCourseLevelLesson, error)
	Find() ([]*model.CurriculumCourseLevelLesson, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CurriculumCourseLevelLesson, err error)
	FindInBatches(result *[]*model.CurriculumCourseLevelLesson, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CurriculumCourseLevelLesson) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICurriculumCourseLevelLessonDo
	Assign(attrs ...field.AssignExpr) ICurriculumCourseLevelLessonDo
	Joins(fields ...field.RelationField) ICurriculumCourseLevelLessonDo
	Preload(fields ...field.RelationField) ICurriculumCourseLevelLessonDo
	FirstOrInit() (*model.CurriculumCourseLevelLesson, error)
	FirstOrCreate() (*model.CurriculumCourseLevelLesson, error)
	FindByPage(offset int, limit int) (result []*model.CurriculumCourseLevelLesson, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICurriculumCourseLevelLessonDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c curriculumCourseLevelLessonDo) Debug() ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Debug())
}

func (c curriculumCourseLevelLessonDo) WithContext(ctx context.Context) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c curriculumCourseLevelLessonDo) ReadDB() ICurriculumCourseLevelLessonDo {
	return c.Clauses(dbresolver.Read)
}

func (c curriculumCourseLevelLessonDo) WriteDB() ICurriculumCourseLevelLessonDo {
	return c.Clauses(dbresolver.Write)
}

func (c curriculumCourseLevelLessonDo) Session(config *gorm.Session) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Session(config))
}

func (c curriculumCourseLevelLessonDo) Clauses(conds ...clause.Expression) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c curriculumCourseLevelLessonDo) Returning(value interface{}, columns ...string) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c curriculumCourseLevelLessonDo) Not(conds ...gen.Condition) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c curriculumCourseLevelLessonDo) Or(conds ...gen.Condition) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c curriculumCourseLevelLessonDo) Select(conds ...field.Expr) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c curriculumCourseLevelLessonDo) Where(conds ...gen.Condition) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c curriculumCourseLevelLessonDo) Order(conds ...field.Expr) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c curriculumCourseLevelLessonDo) Distinct(cols ...field.Expr) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c curriculumCourseLevelLessonDo) Omit(cols ...field.Expr) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c curriculumCourseLevelLessonDo) Join(table schema.Tabler, on ...field.Expr) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c curriculumCourseLevelLessonDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c curriculumCourseLevelLessonDo) RightJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c curriculumCourseLevelLessonDo) Group(cols ...field.Expr) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c curriculumCourseLevelLessonDo) Having(conds ...gen.Condition) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c curriculumCourseLevelLessonDo) Limit(limit int) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c curriculumCourseLevelLessonDo) Offset(offset int) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c curriculumCourseLevelLessonDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c curriculumCourseLevelLessonDo) Unscoped() ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Unscoped())
}

func (c curriculumCourseLevelLessonDo) Create(values ...*model.CurriculumCourseLevelLesson) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c curriculumCourseLevelLessonDo) CreateInBatches(values []*model.CurriculumCourseLevelLesson, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c curriculumCourseLevelLessonDo) Save(values ...*model.CurriculumCourseLevelLesson) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c curriculumCourseLevelLessonDo) First() (*model.CurriculumCourseLevelLesson, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseLevelLesson), nil
	}
}

func (c curriculumCourseLevelLessonDo) Take() (*model.CurriculumCourseLevelLesson, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseLevelLesson), nil
	}
}

func (c curriculumCourseLevelLessonDo) Last() (*model.CurriculumCourseLevelLesson, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseLevelLesson), nil
	}
}

func (c curriculumCourseLevelLessonDo) Find() ([]*model.CurriculumCourseLevelLesson, error) {
	result, err := c.DO.Find()
	return result.([]*model.CurriculumCourseLevelLesson), err
}

func (c curriculumCourseLevelLessonDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CurriculumCourseLevelLesson, err error) {
	buf := make([]*model.CurriculumCourseLevelLesson, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c curriculumCourseLevelLessonDo) FindInBatches(result *[]*model.CurriculumCourseLevelLesson, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c curriculumCourseLevelLessonDo) Attrs(attrs ...field.AssignExpr) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c curriculumCourseLevelLessonDo) Assign(attrs ...field.AssignExpr) ICurriculumCourseLevelLessonDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c curriculumCourseLevelLessonDo) Joins(fields ...field.RelationField) ICurriculumCourseLevelLessonDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c curriculumCourseLevelLessonDo) Preload(fields ...field.RelationField) ICurriculumCourseLevelLessonDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c curriculumCourseLevelLessonDo) FirstOrInit() (*model.CurriculumCourseLevelLesson, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseLevelLesson), nil
	}
}

func (c curriculumCourseLevelLessonDo) FirstOrCreate() (*model.CurriculumCourseLevelLesson, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseLevelLesson), nil
	}
}

func (c curriculumCourseLevelLessonDo) FindByPage(offset int, limit int) (result []*model.CurriculumCourseLevelLesson, count int64, err error) {
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

func (c curriculumCourseLevelLessonDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c curriculumCourseLevelLessonDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c curriculumCourseLevelLessonDo) Delete(models ...*model.CurriculumCourseLevelLesson) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *curriculumCourseLevelLessonDo) withDO(do gen.Dao) *curriculumCourseLevelLessonDo {
	c.DO = *do.(*gen.DO)
	return c
}
