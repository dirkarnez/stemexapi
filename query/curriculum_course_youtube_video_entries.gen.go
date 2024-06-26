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

func newCurriculumCourseYoutubeVideoEntries(db *gorm.DB, opts ...gen.DOOption) curriculumCourseYoutubeVideoEntries {
	_curriculumCourseYoutubeVideoEntries := curriculumCourseYoutubeVideoEntries{}

	_curriculumCourseYoutubeVideoEntries.curriculumCourseYoutubeVideoEntriesDo.UseDB(db, opts...)
	_curriculumCourseYoutubeVideoEntries.curriculumCourseYoutubeVideoEntriesDo.UseModel(&model.CurriculumCourseYoutubeVideoEntries{})

	tableName := _curriculumCourseYoutubeVideoEntries.curriculumCourseYoutubeVideoEntriesDo.TableName()
	_curriculumCourseYoutubeVideoEntries.ALL = field.NewAsterisk(tableName)
	_curriculumCourseYoutubeVideoEntries.ID = field.NewField(tableName, "id")
	_curriculumCourseYoutubeVideoEntries.CreatedAt = field.NewTime(tableName, "created_at")
	_curriculumCourseYoutubeVideoEntries.UpdatedAt = field.NewTime(tableName, "updated_at")
	_curriculumCourseYoutubeVideoEntries.DeletedAt = field.NewField(tableName, "deleted_at")
	_curriculumCourseYoutubeVideoEntries.URL = field.NewString(tableName, "url")
	_curriculumCourseYoutubeVideoEntries.EntryID = field.NewField(tableName, "entry_id")
	_curriculumCourseYoutubeVideoEntries.Entry = curriculumCourseYoutubeVideoEntriesBelongsToEntry{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Entry", "model.CurriculumEntry"),
		Icon: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Entry.Icon", "model.File"),
		},
	}

	_curriculumCourseYoutubeVideoEntries.fillFieldMap()

	return _curriculumCourseYoutubeVideoEntries
}

type curriculumCourseYoutubeVideoEntries struct {
	curriculumCourseYoutubeVideoEntriesDo

	ALL       field.Asterisk
	ID        field.Field
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	URL       field.String
	EntryID   field.Field
	Entry     curriculumCourseYoutubeVideoEntriesBelongsToEntry

	fieldMap map[string]field.Expr
}

func (c curriculumCourseYoutubeVideoEntries) Table(newTableName string) *curriculumCourseYoutubeVideoEntries {
	c.curriculumCourseYoutubeVideoEntriesDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c curriculumCourseYoutubeVideoEntries) As(alias string) *curriculumCourseYoutubeVideoEntries {
	c.curriculumCourseYoutubeVideoEntriesDo.DO = *(c.curriculumCourseYoutubeVideoEntriesDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *curriculumCourseYoutubeVideoEntries) updateTableName(table string) *curriculumCourseYoutubeVideoEntries {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewField(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.URL = field.NewString(table, "url")
	c.EntryID = field.NewField(table, "entry_id")

	c.fillFieldMap()

	return c
}

func (c *curriculumCourseYoutubeVideoEntries) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *curriculumCourseYoutubeVideoEntries) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["url"] = c.URL
	c.fieldMap["entry_id"] = c.EntryID

}

func (c curriculumCourseYoutubeVideoEntries) clone(db *gorm.DB) curriculumCourseYoutubeVideoEntries {
	c.curriculumCourseYoutubeVideoEntriesDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c curriculumCourseYoutubeVideoEntries) replaceDB(db *gorm.DB) curriculumCourseYoutubeVideoEntries {
	c.curriculumCourseYoutubeVideoEntriesDo.ReplaceDB(db)
	return c
}

type curriculumCourseYoutubeVideoEntriesBelongsToEntry struct {
	db *gorm.DB

	field.RelationField

	Icon struct {
		field.RelationField
	}
}

func (a curriculumCourseYoutubeVideoEntriesBelongsToEntry) Where(conds ...field.Expr) *curriculumCourseYoutubeVideoEntriesBelongsToEntry {
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

func (a curriculumCourseYoutubeVideoEntriesBelongsToEntry) WithContext(ctx context.Context) *curriculumCourseYoutubeVideoEntriesBelongsToEntry {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a curriculumCourseYoutubeVideoEntriesBelongsToEntry) Session(session *gorm.Session) *curriculumCourseYoutubeVideoEntriesBelongsToEntry {
	a.db = a.db.Session(session)
	return &a
}

func (a curriculumCourseYoutubeVideoEntriesBelongsToEntry) Model(m *model.CurriculumCourseYoutubeVideoEntries) *curriculumCourseYoutubeVideoEntriesBelongsToEntryTx {
	return &curriculumCourseYoutubeVideoEntriesBelongsToEntryTx{a.db.Model(m).Association(a.Name())}
}

type curriculumCourseYoutubeVideoEntriesBelongsToEntryTx struct{ tx *gorm.Association }

func (a curriculumCourseYoutubeVideoEntriesBelongsToEntryTx) Find() (result *model.CurriculumEntry, err error) {
	return result, a.tx.Find(&result)
}

func (a curriculumCourseYoutubeVideoEntriesBelongsToEntryTx) Append(values ...*model.CurriculumEntry) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a curriculumCourseYoutubeVideoEntriesBelongsToEntryTx) Replace(values ...*model.CurriculumEntry) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a curriculumCourseYoutubeVideoEntriesBelongsToEntryTx) Delete(values ...*model.CurriculumEntry) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a curriculumCourseYoutubeVideoEntriesBelongsToEntryTx) Clear() error {
	return a.tx.Clear()
}

func (a curriculumCourseYoutubeVideoEntriesBelongsToEntryTx) Count() int64 {
	return a.tx.Count()
}

type curriculumCourseYoutubeVideoEntriesDo struct{ gen.DO }

type ICurriculumCourseYoutubeVideoEntriesDo interface {
	gen.SubQuery
	Debug() ICurriculumCourseYoutubeVideoEntriesDo
	WithContext(ctx context.Context) ICurriculumCourseYoutubeVideoEntriesDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICurriculumCourseYoutubeVideoEntriesDo
	WriteDB() ICurriculumCourseYoutubeVideoEntriesDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICurriculumCourseYoutubeVideoEntriesDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICurriculumCourseYoutubeVideoEntriesDo
	Not(conds ...gen.Condition) ICurriculumCourseYoutubeVideoEntriesDo
	Or(conds ...gen.Condition) ICurriculumCourseYoutubeVideoEntriesDo
	Select(conds ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo
	Where(conds ...gen.Condition) ICurriculumCourseYoutubeVideoEntriesDo
	Order(conds ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo
	Distinct(cols ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo
	Omit(cols ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo
	Join(table schema.Tabler, on ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo
	Group(cols ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo
	Having(conds ...gen.Condition) ICurriculumCourseYoutubeVideoEntriesDo
	Limit(limit int) ICurriculumCourseYoutubeVideoEntriesDo
	Offset(offset int) ICurriculumCourseYoutubeVideoEntriesDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICurriculumCourseYoutubeVideoEntriesDo
	Unscoped() ICurriculumCourseYoutubeVideoEntriesDo
	Create(values ...*model.CurriculumCourseYoutubeVideoEntries) error
	CreateInBatches(values []*model.CurriculumCourseYoutubeVideoEntries, batchSize int) error
	Save(values ...*model.CurriculumCourseYoutubeVideoEntries) error
	First() (*model.CurriculumCourseYoutubeVideoEntries, error)
	Take() (*model.CurriculumCourseYoutubeVideoEntries, error)
	Last() (*model.CurriculumCourseYoutubeVideoEntries, error)
	Find() ([]*model.CurriculumCourseYoutubeVideoEntries, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CurriculumCourseYoutubeVideoEntries, err error)
	FindInBatches(result *[]*model.CurriculumCourseYoutubeVideoEntries, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CurriculumCourseYoutubeVideoEntries) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICurriculumCourseYoutubeVideoEntriesDo
	Assign(attrs ...field.AssignExpr) ICurriculumCourseYoutubeVideoEntriesDo
	Joins(fields ...field.RelationField) ICurriculumCourseYoutubeVideoEntriesDo
	Preload(fields ...field.RelationField) ICurriculumCourseYoutubeVideoEntriesDo
	FirstOrInit() (*model.CurriculumCourseYoutubeVideoEntries, error)
	FirstOrCreate() (*model.CurriculumCourseYoutubeVideoEntries, error)
	FindByPage(offset int, limit int) (result []*model.CurriculumCourseYoutubeVideoEntries, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICurriculumCourseYoutubeVideoEntriesDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c curriculumCourseYoutubeVideoEntriesDo) Debug() ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Debug())
}

func (c curriculumCourseYoutubeVideoEntriesDo) WithContext(ctx context.Context) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c curriculumCourseYoutubeVideoEntriesDo) ReadDB() ICurriculumCourseYoutubeVideoEntriesDo {
	return c.Clauses(dbresolver.Read)
}

func (c curriculumCourseYoutubeVideoEntriesDo) WriteDB() ICurriculumCourseYoutubeVideoEntriesDo {
	return c.Clauses(dbresolver.Write)
}

func (c curriculumCourseYoutubeVideoEntriesDo) Session(config *gorm.Session) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Session(config))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Clauses(conds ...clause.Expression) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Returning(value interface{}, columns ...string) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Not(conds ...gen.Condition) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Or(conds ...gen.Condition) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Select(conds ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Where(conds ...gen.Condition) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Order(conds ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Distinct(cols ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Omit(cols ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Join(table schema.Tabler, on ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) RightJoin(table schema.Tabler, on ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Group(cols ...field.Expr) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Having(conds ...gen.Condition) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Limit(limit int) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Offset(offset int) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Unscoped() ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Unscoped())
}

func (c curriculumCourseYoutubeVideoEntriesDo) Create(values ...*model.CurriculumCourseYoutubeVideoEntries) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c curriculumCourseYoutubeVideoEntriesDo) CreateInBatches(values []*model.CurriculumCourseYoutubeVideoEntries, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c curriculumCourseYoutubeVideoEntriesDo) Save(values ...*model.CurriculumCourseYoutubeVideoEntries) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c curriculumCourseYoutubeVideoEntriesDo) First() (*model.CurriculumCourseYoutubeVideoEntries, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseYoutubeVideoEntries), nil
	}
}

func (c curriculumCourseYoutubeVideoEntriesDo) Take() (*model.CurriculumCourseYoutubeVideoEntries, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseYoutubeVideoEntries), nil
	}
}

func (c curriculumCourseYoutubeVideoEntriesDo) Last() (*model.CurriculumCourseYoutubeVideoEntries, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseYoutubeVideoEntries), nil
	}
}

func (c curriculumCourseYoutubeVideoEntriesDo) Find() ([]*model.CurriculumCourseYoutubeVideoEntries, error) {
	result, err := c.DO.Find()
	return result.([]*model.CurriculumCourseYoutubeVideoEntries), err
}

func (c curriculumCourseYoutubeVideoEntriesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CurriculumCourseYoutubeVideoEntries, err error) {
	buf := make([]*model.CurriculumCourseYoutubeVideoEntries, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c curriculumCourseYoutubeVideoEntriesDo) FindInBatches(result *[]*model.CurriculumCourseYoutubeVideoEntries, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c curriculumCourseYoutubeVideoEntriesDo) Attrs(attrs ...field.AssignExpr) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Assign(attrs ...field.AssignExpr) ICurriculumCourseYoutubeVideoEntriesDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c curriculumCourseYoutubeVideoEntriesDo) Joins(fields ...field.RelationField) ICurriculumCourseYoutubeVideoEntriesDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c curriculumCourseYoutubeVideoEntriesDo) Preload(fields ...field.RelationField) ICurriculumCourseYoutubeVideoEntriesDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c curriculumCourseYoutubeVideoEntriesDo) FirstOrInit() (*model.CurriculumCourseYoutubeVideoEntries, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseYoutubeVideoEntries), nil
	}
}

func (c curriculumCourseYoutubeVideoEntriesDo) FirstOrCreate() (*model.CurriculumCourseYoutubeVideoEntries, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumCourseYoutubeVideoEntries), nil
	}
}

func (c curriculumCourseYoutubeVideoEntriesDo) FindByPage(offset int, limit int) (result []*model.CurriculumCourseYoutubeVideoEntries, count int64, err error) {
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

func (c curriculumCourseYoutubeVideoEntriesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c curriculumCourseYoutubeVideoEntriesDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c curriculumCourseYoutubeVideoEntriesDo) Delete(models ...*model.CurriculumCourseYoutubeVideoEntries) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *curriculumCourseYoutubeVideoEntriesDo) withDO(do gen.Dao) *curriculumCourseYoutubeVideoEntriesDo {
	c.DO = *do.(*gen.DO)
	return c
}
