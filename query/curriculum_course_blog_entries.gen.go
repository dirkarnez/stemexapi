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

func newCurriculumCourseBlogEntries(db *gorm.DB, opts ...gen.DOOption) curriculumCourseBlogEntries {
	_curriculumCourseBlogEntries := curriculumCourseBlogEntries{}

	_curriculumCourseBlogEntries.curriculumCourseBlogEntriesDo.UseDB(db, opts...)
	_curriculumCourseBlogEntries.curriculumCourseBlogEntriesDo.UseModel(&model.CurriculumCourseBlogEntries{})

	tableName := _curriculumCourseBlogEntries.curriculumCourseBlogEntriesDo.TableName()
	_curriculumCourseBlogEntries.ALL = field.NewAsterisk(tableName)
	_curriculumCourseBlogEntries.ID = field.NewField(tableName, "id")
	_curriculumCourseBlogEntries.CreatedAt = field.NewTime(tableName, "created_at")
	_curriculumCourseBlogEntries.UpdatedAt = field.NewTime(tableName, "updated_at")
	_curriculumCourseBlogEntries.DeletedAt = field.NewField(tableName, "deleted_at")
	_curriculumCourseBlogEntries.ExternalURL = field.NewString(tableName, "external_url")
	_curriculumCourseBlogEntries.Title = field.NewString(tableName, "title")
	_curriculumCourseBlogEntries.EntryID = field.NewField(tableName, "entry_id")
	_curriculumCourseBlogEntries.Entry = curriculumCourseBlogEntriesBelongsToEntry{
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

	_curriculumCourseBlogEntries.fillFieldMap()

	return _curriculumCourseBlogEntries
}

type curriculumCourseBlogEntries struct {
	curriculumCourseBlogEntriesDo

	ALL         field.Asterisk
	ID          field.Field
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	ExternalURL field.String
	Title       field.String
	EntryID     field.Field
	Entry       curriculumCourseBlogEntriesBelongsToEntry

	fieldMap map[string]field.Expr
}

func (c curriculumCourseBlogEntries) Table(newTableName string) *curriculumCourseBlogEntries {
	c.curriculumCourseBlogEntriesDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c curriculumCourseBlogEntries) As(alias string) *curriculumCourseBlogEntries {
	c.curriculumCourseBlogEntriesDo.DO = *(c.curriculumCourseBlogEntriesDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *curriculumCourseBlogEntries) updateTableName(table string) *curriculumCourseBlogEntries {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewField(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.ExternalURL = field.NewString(table, "external_url")
	c.Title = field.NewString(table, "title")
	c.EntryID = field.NewField(table, "entry_id")

	c.fillFieldMap()

	return c
}

func (c *curriculumCourseBlogEntries) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *curriculumCourseBlogEntries) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 8)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["external_url"] = c.ExternalURL
	c.fieldMap["title"] = c.Title
	c.fieldMap["entry_id"] = c.EntryID

}

func (c curriculumCourseBlogEntries) clone(db *gorm.DB) curriculumCourseBlogEntries {
	c.curriculumCourseBlogEntriesDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c curriculumCourseBlogEntries) replaceDB(db *gorm.DB) curriculumCourseBlogEntries {
	c.curriculumCourseBlogEntriesDo.ReplaceDB(db)
	return c
}

type curriculumCourseBlogEntriesBelongsToEntry struct {
	db *gorm.DB

	field.RelationField

	Icon struct {
		field.RelationField
	}
	CurriculumPlan struct {
		field.RelationField
	}
}

func (a curriculumCourseBlogEntriesBelongsToEntry) Where(conds ...field.Expr) *curriculumCourseBlogEntriesBelongsToEntry {
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

func (a curriculumCourseBlogEntriesBelongsToEntry) WithContext(ctx context.Context) *curriculumCourseBlogEntriesBelongsToEntry {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a curriculumCourseBlogEntriesBelongsToEntry) Session(session *gorm.Session) *curriculumCourseBlogEntriesBelongsToEntry {
	a.db = a.db.Session(session)
	return &a
}

func (a curriculumCourseBlogEntriesBelongsToEntry) Model(m *model.CurriculumCourseBlogEntries) *curriculumCourseBlogEntriesBelongsToEntryTx {
	return &curriculumCourseBlogEntriesBelongsToEntryTx{a.db.Model(m).Association(a.Name())}
}

type curriculumCourseBlogEntriesBelongsToEntryTx struct{ tx *gorm.Association }

func (a curriculumCourseBlogEntriesBelongsToEntryTx) Find() (result *model.CurriculumEntry, err error) {
	return result, a.tx.Find(&result)
}

func (a curriculumCourseBlogEntriesBelongsToEntryTx) Append(values ...*model.CurriculumEntry) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a curriculumCourseBlogEntriesBelongsToEntryTx) Replace(values ...*model.CurriculumEntry) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a curriculumCourseBlogEntriesBelongsToEntryTx) Delete(values ...*model.CurriculumEntry) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a curriculumCourseBlogEntriesBelongsToEntryTx) Clear() error {
	return a.tx.Clear()
}

func (a curriculumCourseBlogEntriesBelongsToEntryTx) Count() int64 {
	return a.tx.Count()
}

type curriculumCourseBlogEntriesDo struct{ gen.DO }

type ICurriculumCourseBlogEntriesDo interface {
	gen.SubQuery
	Debug() ICurriculumCourseBlogEntriesDo
	WithContext(ctx context.Context) ICurriculumCourseBlogEntriesDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICurriculumCourseBlogEntriesDo
	WriteDB() ICurriculumCourseBlogEntriesDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICurriculumCourseBlogEntriesDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICurriculumCourseBlogEntriesDo
	Not(conds ...gen.Condition) ICurriculumCourseBlogEntriesDo
	Or(conds ...gen.Condition) ICurriculumCourseBlogEntriesDo
	Select(conds ...field.Expr) ICurriculumCourseBlogEntriesDo
	Where(conds ...gen.Condition) ICurriculumCourseBlogEntriesDo
	Order(conds ...field.Expr) ICurriculumCourseBlogEntriesDo
	Distinct(cols ...field.Expr) ICurriculumCourseBlogEntriesDo
	Omit(cols ...field.Expr) ICurriculumCourseBlogEntriesDo
	Join(table schema.Tabler, on ...field.Expr) ICurriculumCourseBlogEntriesDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseBlogEntriesDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseBlogEntriesDo
	Group(cols ...field.Expr) ICurriculumCourseBlogEntriesDo
	Having(conds ...gen.Condition) ICurriculumCourseBlogEntriesDo
	Limit(limit int) ICurriculumCourseBlogEntriesDo
	Offset(offset int) ICurriculumCourseBlogEntriesDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICurriculumCourseBlogEntriesDo
	Unscoped() ICurriculumCourseBlogEntriesDo
	Create(values ...*model.CurriculumCourseBlogEntries) error
	CreateInBatches(values []*model.CurriculumCourseBlogEntries, batchSize int) error
	Save(values ...*model.CurriculumCourseBlogEntries) error
	First() (*model.CurriculumCourseBlogEntries, error)
	Take() (*model.CurriculumCourseBlogEntries, error)
	Last() (*model.CurriculumCourseBlogEntries, error)
	Find() ([]*model.CurriculumCourseBlogEntries, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CurriculumCourseBlogEntries, err error)
	FindInBatches(result *[]*model.CurriculumCourseBlogEntries, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CurriculumCourseBlogEntries) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICurriculumCourseBlogEntriesDo
	Assign(attrs ...field.AssignExpr) ICurriculumCourseBlogEntriesDo
	Joins(fields ...field.RelationField) ICurriculumCourseBlogEntriesDo
	Preload(fields ...field.RelationField) ICurriculumCourseBlogEntriesDo
	FirstOrInit() (*model.CurriculumCourseBlogEntries, error)
	FirstOrCreate() (*model.CurriculumCourseBlogEntries, error)
	FindByPage(offset int, limit int) (result []*model.CurriculumCourseBlogEntries, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICurriculumCourseBlogEntriesDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithNameAndRole(name string, role string) (result []model.CurriculumCourseBlogEntries, err error)
}

// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
func (c curriculumCourseBlogEntriesDo) FilterWithNameAndRole(name string, role string) (result []model.CurriculumCourseBlogEntries, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM curriculum_course_blog_entries WHERE name = ? ")
	if role != "" {
		params = append(params, role)
		generateSQL.WriteString("AND role = ? ")
	}

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (c curriculumCourseBlogEntriesDo) Debug() ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Debug())
}

func (c curriculumCourseBlogEntriesDo) WithContext(ctx context.Context) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c curriculumCourseBlogEntriesDo) ReadDB() ICurriculumCourseBlogEntriesDo {
	return c.Clauses(dbresolver.Read)
}

func (c curriculumCourseBlogEntriesDo) WriteDB() ICurriculumCourseBlogEntriesDo {
	return c.Clauses(dbresolver.Write)
}

func (c curriculumCourseBlogEntriesDo) Session(config *gorm.Session) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Session(config))
}

func (c curriculumCourseBlogEntriesDo) Clauses(conds ...clause.Expression) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c curriculumCourseBlogEntriesDo) Returning(value interface{}, columns ...string) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c curriculumCourseBlogEntriesDo) Not(conds ...gen.Condition) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c curriculumCourseBlogEntriesDo) Or(conds ...gen.Condition) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c curriculumCourseBlogEntriesDo) Select(conds ...field.Expr) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c curriculumCourseBlogEntriesDo) Where(conds ...gen.Condition) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c curriculumCourseBlogEntriesDo) Order(conds ...field.Expr) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c curriculumCourseBlogEntriesDo) Distinct(cols ...field.Expr) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c curriculumCourseBlogEntriesDo) Omit(cols ...field.Expr) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c curriculumCourseBlogEntriesDo) Join(table schema.Tabler, on ...field.Expr) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c curriculumCourseBlogEntriesDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c curriculumCourseBlogEntriesDo) RightJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c curriculumCourseBlogEntriesDo) Group(cols ...field.Expr) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c curriculumCourseBlogEntriesDo) Having(conds ...gen.Condition) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c curriculumCourseBlogEntriesDo) Limit(limit int) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c curriculumCourseBlogEntriesDo) Offset(offset int) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c curriculumCourseBlogEntriesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c curriculumCourseBlogEntriesDo) Unscoped() ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Unscoped())
}

func (c curriculumCourseBlogEntriesDo) Create(values ...*model.CurriculumCourseBlogEntries) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c curriculumCourseBlogEntriesDo) CreateInBatches(values []*model.CurriculumCourseBlogEntries, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c curriculumCourseBlogEntriesDo) Save(values ...*model.CurriculumCourseBlogEntries) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c curriculumCourseBlogEntriesDo) First() (*model.CurriculumCourseBlogEntries, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseBlogEntries), nil
	}
}

func (c curriculumCourseBlogEntriesDo) Take() (*model.CurriculumCourseBlogEntries, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseBlogEntries), nil
	}
}

func (c curriculumCourseBlogEntriesDo) Last() (*model.CurriculumCourseBlogEntries, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseBlogEntries), nil
	}
}

func (c curriculumCourseBlogEntriesDo) Find() ([]*model.CurriculumCourseBlogEntries, error) {
	result, err := c.DO.Find()
	return result.([]*model.CurriculumCourseBlogEntries), err
}

func (c curriculumCourseBlogEntriesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CurriculumCourseBlogEntries, err error) {
	buf := make([]*model.CurriculumCourseBlogEntries, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c curriculumCourseBlogEntriesDo) FindInBatches(result *[]*model.CurriculumCourseBlogEntries, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c curriculumCourseBlogEntriesDo) Attrs(attrs ...field.AssignExpr) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c curriculumCourseBlogEntriesDo) Assign(attrs ...field.AssignExpr) ICurriculumCourseBlogEntriesDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c curriculumCourseBlogEntriesDo) Joins(fields ...field.RelationField) ICurriculumCourseBlogEntriesDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c curriculumCourseBlogEntriesDo) Preload(fields ...field.RelationField) ICurriculumCourseBlogEntriesDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c curriculumCourseBlogEntriesDo) FirstOrInit() (*model.CurriculumCourseBlogEntries, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseBlogEntries), nil
	}
}

func (c curriculumCourseBlogEntriesDo) FirstOrCreate() (*model.CurriculumCourseBlogEntries, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseBlogEntries), nil
	}
}

func (c curriculumCourseBlogEntriesDo) FindByPage(offset int, limit int) (result []*model.CurriculumCourseBlogEntries, count int64, err error) {
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

func (c curriculumCourseBlogEntriesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c curriculumCourseBlogEntriesDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c curriculumCourseBlogEntriesDo) Delete(models ...*model.CurriculumCourseBlogEntries) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *curriculumCourseBlogEntriesDo) withDO(do gen.Dao) *curriculumCourseBlogEntriesDo {
	c.DO = *do.(*gen.DO)
	return c
}
