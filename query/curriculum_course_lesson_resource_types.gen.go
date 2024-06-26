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

func newCurriculumCourseLessonResourceType(db *gorm.DB, opts ...gen.DOOption) curriculumCourseLessonResourceType {
	_curriculumCourseLessonResourceType := curriculumCourseLessonResourceType{}

	_curriculumCourseLessonResourceType.curriculumCourseLessonResourceTypeDo.UseDB(db, opts...)
	_curriculumCourseLessonResourceType.curriculumCourseLessonResourceTypeDo.UseModel(&model.CurriculumCourseLessonResourceType{})

	tableName := _curriculumCourseLessonResourceType.curriculumCourseLessonResourceTypeDo.TableName()
	_curriculumCourseLessonResourceType.ALL = field.NewAsterisk(tableName)
	_curriculumCourseLessonResourceType.ID = field.NewField(tableName, "id")
	_curriculumCourseLessonResourceType.CreatedAt = field.NewTime(tableName, "created_at")
	_curriculumCourseLessonResourceType.UpdatedAt = field.NewTime(tableName, "updated_at")
	_curriculumCourseLessonResourceType.DeletedAt = field.NewField(tableName, "deleted_at")
	_curriculumCourseLessonResourceType.Name = field.NewString(tableName, "name")

	_curriculumCourseLessonResourceType.fillFieldMap()

	return _curriculumCourseLessonResourceType
}

type curriculumCourseLessonResourceType struct {
	curriculumCourseLessonResourceTypeDo

	ALL       field.Asterisk
	ID        field.Field
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Name      field.String

	fieldMap map[string]field.Expr
}

func (c curriculumCourseLessonResourceType) Table(newTableName string) *curriculumCourseLessonResourceType {
	c.curriculumCourseLessonResourceTypeDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c curriculumCourseLessonResourceType) As(alias string) *curriculumCourseLessonResourceType {
	c.curriculumCourseLessonResourceTypeDo.DO = *(c.curriculumCourseLessonResourceTypeDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *curriculumCourseLessonResourceType) updateTableName(table string) *curriculumCourseLessonResourceType {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewField(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.Name = field.NewString(table, "name")

	c.fillFieldMap()

	return c
}

func (c *curriculumCourseLessonResourceType) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *curriculumCourseLessonResourceType) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 5)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["name"] = c.Name
}

func (c curriculumCourseLessonResourceType) clone(db *gorm.DB) curriculumCourseLessonResourceType {
	c.curriculumCourseLessonResourceTypeDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c curriculumCourseLessonResourceType) replaceDB(db *gorm.DB) curriculumCourseLessonResourceType {
	c.curriculumCourseLessonResourceTypeDo.ReplaceDB(db)
	return c
}

type curriculumCourseLessonResourceTypeDo struct{ gen.DO }

type ICurriculumCourseLessonResourceTypeDo interface {
	gen.SubQuery
	Debug() ICurriculumCourseLessonResourceTypeDo
	WithContext(ctx context.Context) ICurriculumCourseLessonResourceTypeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICurriculumCourseLessonResourceTypeDo
	WriteDB() ICurriculumCourseLessonResourceTypeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICurriculumCourseLessonResourceTypeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICurriculumCourseLessonResourceTypeDo
	Not(conds ...gen.Condition) ICurriculumCourseLessonResourceTypeDo
	Or(conds ...gen.Condition) ICurriculumCourseLessonResourceTypeDo
	Select(conds ...field.Expr) ICurriculumCourseLessonResourceTypeDo
	Where(conds ...gen.Condition) ICurriculumCourseLessonResourceTypeDo
	Order(conds ...field.Expr) ICurriculumCourseLessonResourceTypeDo
	Distinct(cols ...field.Expr) ICurriculumCourseLessonResourceTypeDo
	Omit(cols ...field.Expr) ICurriculumCourseLessonResourceTypeDo
	Join(table schema.Tabler, on ...field.Expr) ICurriculumCourseLessonResourceTypeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseLessonResourceTypeDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseLessonResourceTypeDo
	Group(cols ...field.Expr) ICurriculumCourseLessonResourceTypeDo
	Having(conds ...gen.Condition) ICurriculumCourseLessonResourceTypeDo
	Limit(limit int) ICurriculumCourseLessonResourceTypeDo
	Offset(offset int) ICurriculumCourseLessonResourceTypeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICurriculumCourseLessonResourceTypeDo
	Unscoped() ICurriculumCourseLessonResourceTypeDo
	Create(values ...*model.CurriculumCourseLessonResourceType) error
	CreateInBatches(values []*model.CurriculumCourseLessonResourceType, batchSize int) error
	Save(values ...*model.CurriculumCourseLessonResourceType) error
	First() (*model.CurriculumCourseLessonResourceType, error)
	Take() (*model.CurriculumCourseLessonResourceType, error)
	Last() (*model.CurriculumCourseLessonResourceType, error)
	Find() ([]*model.CurriculumCourseLessonResourceType, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CurriculumCourseLessonResourceType, err error)
	FindInBatches(result *[]*model.CurriculumCourseLessonResourceType, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CurriculumCourseLessonResourceType) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICurriculumCourseLessonResourceTypeDo
	Assign(attrs ...field.AssignExpr) ICurriculumCourseLessonResourceTypeDo
	Joins(fields ...field.RelationField) ICurriculumCourseLessonResourceTypeDo
	Preload(fields ...field.RelationField) ICurriculumCourseLessonResourceTypeDo
	FirstOrInit() (*model.CurriculumCourseLessonResourceType, error)
	FirstOrCreate() (*model.CurriculumCourseLessonResourceType, error)
	FindByPage(offset int, limit int) (result []*model.CurriculumCourseLessonResourceType, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICurriculumCourseLessonResourceTypeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c curriculumCourseLessonResourceTypeDo) Debug() ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Debug())
}

func (c curriculumCourseLessonResourceTypeDo) WithContext(ctx context.Context) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c curriculumCourseLessonResourceTypeDo) ReadDB() ICurriculumCourseLessonResourceTypeDo {
	return c.Clauses(dbresolver.Read)
}

func (c curriculumCourseLessonResourceTypeDo) WriteDB() ICurriculumCourseLessonResourceTypeDo {
	return c.Clauses(dbresolver.Write)
}

func (c curriculumCourseLessonResourceTypeDo) Session(config *gorm.Session) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Session(config))
}

func (c curriculumCourseLessonResourceTypeDo) Clauses(conds ...clause.Expression) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c curriculumCourseLessonResourceTypeDo) Returning(value interface{}, columns ...string) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c curriculumCourseLessonResourceTypeDo) Not(conds ...gen.Condition) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c curriculumCourseLessonResourceTypeDo) Or(conds ...gen.Condition) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c curriculumCourseLessonResourceTypeDo) Select(conds ...field.Expr) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c curriculumCourseLessonResourceTypeDo) Where(conds ...gen.Condition) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c curriculumCourseLessonResourceTypeDo) Order(conds ...field.Expr) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c curriculumCourseLessonResourceTypeDo) Distinct(cols ...field.Expr) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c curriculumCourseLessonResourceTypeDo) Omit(cols ...field.Expr) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c curriculumCourseLessonResourceTypeDo) Join(table schema.Tabler, on ...field.Expr) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c curriculumCourseLessonResourceTypeDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c curriculumCourseLessonResourceTypeDo) RightJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c curriculumCourseLessonResourceTypeDo) Group(cols ...field.Expr) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c curriculumCourseLessonResourceTypeDo) Having(conds ...gen.Condition) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c curriculumCourseLessonResourceTypeDo) Limit(limit int) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c curriculumCourseLessonResourceTypeDo) Offset(offset int) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c curriculumCourseLessonResourceTypeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c curriculumCourseLessonResourceTypeDo) Unscoped() ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Unscoped())
}

func (c curriculumCourseLessonResourceTypeDo) Create(values ...*model.CurriculumCourseLessonResourceType) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c curriculumCourseLessonResourceTypeDo) CreateInBatches(values []*model.CurriculumCourseLessonResourceType, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c curriculumCourseLessonResourceTypeDo) Save(values ...*model.CurriculumCourseLessonResourceType) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c curriculumCourseLessonResourceTypeDo) First() (*model.CurriculumCourseLessonResourceType, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseLessonResourceType), nil
	}
}

func (c curriculumCourseLessonResourceTypeDo) Take() (*model.CurriculumCourseLessonResourceType, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseLessonResourceType), nil
	}
}

func (c curriculumCourseLessonResourceTypeDo) Last() (*model.CurriculumCourseLessonResourceType, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseLessonResourceType), nil
	}
}

func (c curriculumCourseLessonResourceTypeDo) Find() ([]*model.CurriculumCourseLessonResourceType, error) {
	result, err := c.DO.Find()
	return result.([]*model.CurriculumCourseLessonResourceType), err
}

func (c curriculumCourseLessonResourceTypeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CurriculumCourseLessonResourceType, err error) {
	buf := make([]*model.CurriculumCourseLessonResourceType, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c curriculumCourseLessonResourceTypeDo) FindInBatches(result *[]*model.CurriculumCourseLessonResourceType, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c curriculumCourseLessonResourceTypeDo) Attrs(attrs ...field.AssignExpr) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c curriculumCourseLessonResourceTypeDo) Assign(attrs ...field.AssignExpr) ICurriculumCourseLessonResourceTypeDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c curriculumCourseLessonResourceTypeDo) Joins(fields ...field.RelationField) ICurriculumCourseLessonResourceTypeDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c curriculumCourseLessonResourceTypeDo) Preload(fields ...field.RelationField) ICurriculumCourseLessonResourceTypeDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c curriculumCourseLessonResourceTypeDo) FirstOrInit() (*model.CurriculumCourseLessonResourceType, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseLessonResourceType), nil
	}
}

func (c curriculumCourseLessonResourceTypeDo) FirstOrCreate() (*model.CurriculumCourseLessonResourceType, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseLessonResourceType), nil
	}
}

func (c curriculumCourseLessonResourceTypeDo) FindByPage(offset int, limit int) (result []*model.CurriculumCourseLessonResourceType, count int64, err error) {
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

func (c curriculumCourseLessonResourceTypeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c curriculumCourseLessonResourceTypeDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c curriculumCourseLessonResourceTypeDo) Delete(models ...*model.CurriculumCourseLessonResourceType) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *curriculumCourseLessonResourceTypeDo) withDO(do gen.Dao) *curriculumCourseLessonResourceTypeDo {
	c.DO = *do.(*gen.DO)
	return c
}
