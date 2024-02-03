// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"github.com/dirkarnez/stemexapi/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newCurriculumCoursePrerequisites(db *gorm.DB, opts ...gen.DOOption) curriculumCoursePrerequisites {
	_curriculumCoursePrerequisites := curriculumCoursePrerequisites{}

	_curriculumCoursePrerequisites.curriculumCoursePrerequisitesDo.UseDB(db, opts...)
	_curriculumCoursePrerequisites.curriculumCoursePrerequisitesDo.UseModel(&model.CurriculumCoursePrerequisites{})

	tableName := _curriculumCoursePrerequisites.curriculumCoursePrerequisitesDo.TableName()
	_curriculumCoursePrerequisites.ALL = field.NewAsterisk(tableName)
	_curriculumCoursePrerequisites.ID = field.NewField(tableName, "id")
	_curriculumCoursePrerequisites.CreatedAt = field.NewTime(tableName, "created_at")
	_curriculumCoursePrerequisites.UpdatedAt = field.NewTime(tableName, "updated_at")
	_curriculumCoursePrerequisites.DeletedAt = field.NewField(tableName, "deleted_at")
	_curriculumCoursePrerequisites.Content = field.NewString(tableName, "content")
	_curriculumCoursePrerequisites.EntryID = field.NewField(tableName, "entry_id")
	_curriculumCoursePrerequisites.Entry = curriculumCoursePrerequisitesBelongsToEntry{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Entry", "model.CurriculumEntry"),
		Icon: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Entry.Icon", "model.File"),
		},
		CurriculumPlan: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Entry.CurriculumPlan", "model.File"),
		},
	}

	_curriculumCoursePrerequisites.fillFieldMap()

	return _curriculumCoursePrerequisites
}

type curriculumCoursePrerequisites struct {
	curriculumCoursePrerequisitesDo

	ALL       field.Asterisk
	ID        field.Field
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Content   field.String
	EntryID   field.Field
	Entry     curriculumCoursePrerequisitesBelongsToEntry

	fieldMap map[string]field.Expr
}

func (c curriculumCoursePrerequisites) Table(newTableName string) *curriculumCoursePrerequisites {
	c.curriculumCoursePrerequisitesDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c curriculumCoursePrerequisites) As(alias string) *curriculumCoursePrerequisites {
	c.curriculumCoursePrerequisitesDo.DO = *(c.curriculumCoursePrerequisitesDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *curriculumCoursePrerequisites) updateTableName(table string) *curriculumCoursePrerequisites {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewField(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.Content = field.NewString(table, "content")
	c.EntryID = field.NewField(table, "entry_id")

	c.fillFieldMap()

	return c
}

func (c *curriculumCoursePrerequisites) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *curriculumCoursePrerequisites) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["content"] = c.Content
	c.fieldMap["entry_id"] = c.EntryID

}

func (c curriculumCoursePrerequisites) clone(db *gorm.DB) curriculumCoursePrerequisites {
	c.curriculumCoursePrerequisitesDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c curriculumCoursePrerequisites) replaceDB(db *gorm.DB) curriculumCoursePrerequisites {
	c.curriculumCoursePrerequisitesDo.ReplaceDB(db)
	return c
}

type curriculumCoursePrerequisitesBelongsToEntry struct {
	db *gorm.DB

	field.RelationField

	Icon struct {
		field.RelationField
	}
	CurriculumPlan struct {
		field.RelationField
	}
}

func (a curriculumCoursePrerequisitesBelongsToEntry) Where(conds ...field.Expr) *curriculumCoursePrerequisitesBelongsToEntry {
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

func (a curriculumCoursePrerequisitesBelongsToEntry) WithContext(ctx context.Context) *curriculumCoursePrerequisitesBelongsToEntry {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a curriculumCoursePrerequisitesBelongsToEntry) Session(session *gorm.Session) *curriculumCoursePrerequisitesBelongsToEntry {
	a.db = a.db.Session(session)
	return &a
}

func (a curriculumCoursePrerequisitesBelongsToEntry) Model(m *model.CurriculumCoursePrerequisites) *curriculumCoursePrerequisitesBelongsToEntryTx {
	return &curriculumCoursePrerequisitesBelongsToEntryTx{a.db.Model(m).Association(a.Name())}
}

type curriculumCoursePrerequisitesBelongsToEntryTx struct{ tx *gorm.Association }

func (a curriculumCoursePrerequisitesBelongsToEntryTx) Find() (result *model.CurriculumEntry, err error) {
	return result, a.tx.Find(&result)
}

func (a curriculumCoursePrerequisitesBelongsToEntryTx) Append(values ...*model.CurriculumEntry) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a curriculumCoursePrerequisitesBelongsToEntryTx) Replace(values ...*model.CurriculumEntry) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a curriculumCoursePrerequisitesBelongsToEntryTx) Delete(values ...*model.CurriculumEntry) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a curriculumCoursePrerequisitesBelongsToEntryTx) Clear() error {
	return a.tx.Clear()
}

func (a curriculumCoursePrerequisitesBelongsToEntryTx) Count() int64 {
	return a.tx.Count()
}

type curriculumCoursePrerequisitesDo struct{ gen.DO }

type ICurriculumCoursePrerequisitesDo interface {
	gen.SubQuery
	Debug() ICurriculumCoursePrerequisitesDo
	WithContext(ctx context.Context) ICurriculumCoursePrerequisitesDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICurriculumCoursePrerequisitesDo
	WriteDB() ICurriculumCoursePrerequisitesDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICurriculumCoursePrerequisitesDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICurriculumCoursePrerequisitesDo
	Not(conds ...gen.Condition) ICurriculumCoursePrerequisitesDo
	Or(conds ...gen.Condition) ICurriculumCoursePrerequisitesDo
	Select(conds ...field.Expr) ICurriculumCoursePrerequisitesDo
	Where(conds ...gen.Condition) ICurriculumCoursePrerequisitesDo
	Order(conds ...field.Expr) ICurriculumCoursePrerequisitesDo
	Distinct(cols ...field.Expr) ICurriculumCoursePrerequisitesDo
	Omit(cols ...field.Expr) ICurriculumCoursePrerequisitesDo
	Join(table schema.Tabler, on ...field.Expr) ICurriculumCoursePrerequisitesDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICurriculumCoursePrerequisitesDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICurriculumCoursePrerequisitesDo
	Group(cols ...field.Expr) ICurriculumCoursePrerequisitesDo
	Having(conds ...gen.Condition) ICurriculumCoursePrerequisitesDo
	Limit(limit int) ICurriculumCoursePrerequisitesDo
	Offset(offset int) ICurriculumCoursePrerequisitesDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICurriculumCoursePrerequisitesDo
	Unscoped() ICurriculumCoursePrerequisitesDo
	Create(values ...*model.CurriculumCoursePrerequisites) error
	CreateInBatches(values []*model.CurriculumCoursePrerequisites, batchSize int) error
	Save(values ...*model.CurriculumCoursePrerequisites) error
	First() (*model.CurriculumCoursePrerequisites, error)
	Take() (*model.CurriculumCoursePrerequisites, error)
	Last() (*model.CurriculumCoursePrerequisites, error)
	Find() ([]*model.CurriculumCoursePrerequisites, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CurriculumCoursePrerequisites, err error)
	FindInBatches(result *[]*model.CurriculumCoursePrerequisites, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CurriculumCoursePrerequisites) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICurriculumCoursePrerequisitesDo
	Assign(attrs ...field.AssignExpr) ICurriculumCoursePrerequisitesDo
	Joins(fields ...field.RelationField) ICurriculumCoursePrerequisitesDo
	Preload(fields ...field.RelationField) ICurriculumCoursePrerequisitesDo
	FirstOrInit() (*model.CurriculumCoursePrerequisites, error)
	FirstOrCreate() (*model.CurriculumCoursePrerequisites, error)
	FindByPage(offset int, limit int) (result []*model.CurriculumCoursePrerequisites, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICurriculumCoursePrerequisitesDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithNameAndRole(name string, role string) (result []model.CurriculumCoursePrerequisites, err error)
}

// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
func (c curriculumCoursePrerequisitesDo) FilterWithNameAndRole(name string, role string) (result []model.CurriculumCoursePrerequisites, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM curriculum_course_prerequisites WHERE name = ? ")
	if role != "" {
		params = append(params, role)
		generateSQL.WriteString("AND role = ? ")
	}

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (c curriculumCoursePrerequisitesDo) Debug() ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Debug())
}

func (c curriculumCoursePrerequisitesDo) WithContext(ctx context.Context) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c curriculumCoursePrerequisitesDo) ReadDB() ICurriculumCoursePrerequisitesDo {
	return c.Clauses(dbresolver.Read)
}

func (c curriculumCoursePrerequisitesDo) WriteDB() ICurriculumCoursePrerequisitesDo {
	return c.Clauses(dbresolver.Write)
}

func (c curriculumCoursePrerequisitesDo) Session(config *gorm.Session) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Session(config))
}

func (c curriculumCoursePrerequisitesDo) Clauses(conds ...clause.Expression) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c curriculumCoursePrerequisitesDo) Returning(value interface{}, columns ...string) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c curriculumCoursePrerequisitesDo) Not(conds ...gen.Condition) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c curriculumCoursePrerequisitesDo) Or(conds ...gen.Condition) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c curriculumCoursePrerequisitesDo) Select(conds ...field.Expr) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c curriculumCoursePrerequisitesDo) Where(conds ...gen.Condition) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c curriculumCoursePrerequisitesDo) Order(conds ...field.Expr) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c curriculumCoursePrerequisitesDo) Distinct(cols ...field.Expr) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c curriculumCoursePrerequisitesDo) Omit(cols ...field.Expr) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c curriculumCoursePrerequisitesDo) Join(table schema.Tabler, on ...field.Expr) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c curriculumCoursePrerequisitesDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c curriculumCoursePrerequisitesDo) RightJoin(table schema.Tabler, on ...field.Expr) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c curriculumCoursePrerequisitesDo) Group(cols ...field.Expr) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c curriculumCoursePrerequisitesDo) Having(conds ...gen.Condition) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c curriculumCoursePrerequisitesDo) Limit(limit int) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c curriculumCoursePrerequisitesDo) Offset(offset int) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c curriculumCoursePrerequisitesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c curriculumCoursePrerequisitesDo) Unscoped() ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Unscoped())
}

func (c curriculumCoursePrerequisitesDo) Create(values ...*model.CurriculumCoursePrerequisites) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c curriculumCoursePrerequisitesDo) CreateInBatches(values []*model.CurriculumCoursePrerequisites, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c curriculumCoursePrerequisitesDo) Save(values ...*model.CurriculumCoursePrerequisites) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c curriculumCoursePrerequisitesDo) First() (*model.CurriculumCoursePrerequisites, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCoursePrerequisites), nil
	}
}

func (c curriculumCoursePrerequisitesDo) Take() (*model.CurriculumCoursePrerequisites, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCoursePrerequisites), nil
	}
}

func (c curriculumCoursePrerequisitesDo) Last() (*model.CurriculumCoursePrerequisites, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCoursePrerequisites), nil
	}
}

func (c curriculumCoursePrerequisitesDo) Find() ([]*model.CurriculumCoursePrerequisites, error) {
	result, err := c.DO.Find()
	return result.([]*model.CurriculumCoursePrerequisites), err
}

func (c curriculumCoursePrerequisitesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CurriculumCoursePrerequisites, err error) {
	buf := make([]*model.CurriculumCoursePrerequisites, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c curriculumCoursePrerequisitesDo) FindInBatches(result *[]*model.CurriculumCoursePrerequisites, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c curriculumCoursePrerequisitesDo) Attrs(attrs ...field.AssignExpr) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c curriculumCoursePrerequisitesDo) Assign(attrs ...field.AssignExpr) ICurriculumCoursePrerequisitesDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c curriculumCoursePrerequisitesDo) Joins(fields ...field.RelationField) ICurriculumCoursePrerequisitesDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c curriculumCoursePrerequisitesDo) Preload(fields ...field.RelationField) ICurriculumCoursePrerequisitesDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c curriculumCoursePrerequisitesDo) FirstOrInit() (*model.CurriculumCoursePrerequisites, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCoursePrerequisites), nil
	}
}

func (c curriculumCoursePrerequisitesDo) FirstOrCreate() (*model.CurriculumCoursePrerequisites, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCoursePrerequisites), nil
	}
}

func (c curriculumCoursePrerequisitesDo) FindByPage(offset int, limit int) (result []*model.CurriculumCoursePrerequisites, count int64, err error) {
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

func (c curriculumCoursePrerequisitesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c curriculumCoursePrerequisitesDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c curriculumCoursePrerequisitesDo) Delete(models ...*model.CurriculumCoursePrerequisites) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *curriculumCoursePrerequisitesDo) withDO(do gen.Dao) *curriculumCoursePrerequisitesDo {
	c.DO = *do.(*gen.DO)
	return c
}
