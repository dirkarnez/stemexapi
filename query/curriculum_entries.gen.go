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

func newCurriculumEntry(db *gorm.DB, opts ...gen.DOOption) curriculumEntry {
	_curriculumEntry := curriculumEntry{}

	_curriculumEntry.curriculumEntryDo.UseDB(db, opts...)
	_curriculumEntry.curriculumEntryDo.UseModel(&model.CurriculumEntry{})

	tableName := _curriculumEntry.curriculumEntryDo.TableName()
	_curriculumEntry.ALL = field.NewAsterisk(tableName)
	_curriculumEntry.ID = field.NewField(tableName, "id")
	_curriculumEntry.CreatedAt = field.NewTime(tableName, "created_at")
	_curriculumEntry.UpdatedAt = field.NewTime(tableName, "updated_at")
	_curriculumEntry.DeletedAt = field.NewField(tableName, "deleted_at")
	_curriculumEntry.IconID = field.NewField(tableName, "icon_id")
	_curriculumEntry.Description = field.NewString(tableName, "description")
	_curriculumEntry.ParentID = field.NewField(tableName, "parent_id")
	_curriculumEntry.SeqNoSameLevel = field.NewUint64(tableName, "seq_no_same_level")
	_curriculumEntry.Icon = curriculumEntryBelongsToIcon{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Icon", "model.File"),
	}

	_curriculumEntry.fillFieldMap()

	return _curriculumEntry
}

type curriculumEntry struct {
	curriculumEntryDo

	ALL            field.Asterisk
	ID             field.Field
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field
	IconID         field.Field
	Description    field.String
	ParentID       field.Field
	SeqNoSameLevel field.Uint64
	Icon           curriculumEntryBelongsToIcon

	fieldMap map[string]field.Expr
}

func (c curriculumEntry) Table(newTableName string) *curriculumEntry {
	c.curriculumEntryDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c curriculumEntry) As(alias string) *curriculumEntry {
	c.curriculumEntryDo.DO = *(c.curriculumEntryDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *curriculumEntry) updateTableName(table string) *curriculumEntry {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewField(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.IconID = field.NewField(table, "icon_id")
	c.Description = field.NewString(table, "description")
	c.ParentID = field.NewField(table, "parent_id")
	c.SeqNoSameLevel = field.NewUint64(table, "seq_no_same_level")

	c.fillFieldMap()

	return c
}

func (c *curriculumEntry) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *curriculumEntry) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["icon_id"] = c.IconID
	c.fieldMap["description"] = c.Description
	c.fieldMap["parent_id"] = c.ParentID
	c.fieldMap["seq_no_same_level"] = c.SeqNoSameLevel

}

func (c curriculumEntry) clone(db *gorm.DB) curriculumEntry {
	c.curriculumEntryDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c curriculumEntry) replaceDB(db *gorm.DB) curriculumEntry {
	c.curriculumEntryDo.ReplaceDB(db)
	return c
}

type curriculumEntryBelongsToIcon struct {
	db *gorm.DB

	field.RelationField
}

func (a curriculumEntryBelongsToIcon) Where(conds ...field.Expr) *curriculumEntryBelongsToIcon {
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

func (a curriculumEntryBelongsToIcon) WithContext(ctx context.Context) *curriculumEntryBelongsToIcon {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a curriculumEntryBelongsToIcon) Session(session *gorm.Session) *curriculumEntryBelongsToIcon {
	a.db = a.db.Session(session)
	return &a
}

func (a curriculumEntryBelongsToIcon) Model(m *model.CurriculumEntry) *curriculumEntryBelongsToIconTx {
	return &curriculumEntryBelongsToIconTx{a.db.Model(m).Association(a.Name())}
}

type curriculumEntryBelongsToIconTx struct{ tx *gorm.Association }

func (a curriculumEntryBelongsToIconTx) Find() (result *model.File, err error) {
	return result, a.tx.Find(&result)
}

func (a curriculumEntryBelongsToIconTx) Append(values ...*model.File) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a curriculumEntryBelongsToIconTx) Replace(values ...*model.File) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a curriculumEntryBelongsToIconTx) Delete(values ...*model.File) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a curriculumEntryBelongsToIconTx) Clear() error {
	return a.tx.Clear()
}

func (a curriculumEntryBelongsToIconTx) Count() int64 {
	return a.tx.Count()
}

type curriculumEntryDo struct{ gen.DO }

type ICurriculumEntryDo interface {
	gen.SubQuery
	Debug() ICurriculumEntryDo
	WithContext(ctx context.Context) ICurriculumEntryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICurriculumEntryDo
	WriteDB() ICurriculumEntryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICurriculumEntryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICurriculumEntryDo
	Not(conds ...gen.Condition) ICurriculumEntryDo
	Or(conds ...gen.Condition) ICurriculumEntryDo
	Select(conds ...field.Expr) ICurriculumEntryDo
	Where(conds ...gen.Condition) ICurriculumEntryDo
	Order(conds ...field.Expr) ICurriculumEntryDo
	Distinct(cols ...field.Expr) ICurriculumEntryDo
	Omit(cols ...field.Expr) ICurriculumEntryDo
	Join(table schema.Tabler, on ...field.Expr) ICurriculumEntryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICurriculumEntryDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICurriculumEntryDo
	Group(cols ...field.Expr) ICurriculumEntryDo
	Having(conds ...gen.Condition) ICurriculumEntryDo
	Limit(limit int) ICurriculumEntryDo
	Offset(offset int) ICurriculumEntryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICurriculumEntryDo
	Unscoped() ICurriculumEntryDo
	Create(values ...*model.CurriculumEntry) error
	CreateInBatches(values []*model.CurriculumEntry, batchSize int) error
	Save(values ...*model.CurriculumEntry) error
	First() (*model.CurriculumEntry, error)
	Take() (*model.CurriculumEntry, error)
	Last() (*model.CurriculumEntry, error)
	Find() ([]*model.CurriculumEntry, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CurriculumEntry, err error)
	FindInBatches(result *[]*model.CurriculumEntry, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CurriculumEntry) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICurriculumEntryDo
	Assign(attrs ...field.AssignExpr) ICurriculumEntryDo
	Joins(fields ...field.RelationField) ICurriculumEntryDo
	Preload(fields ...field.RelationField) ICurriculumEntryDo
	FirstOrInit() (*model.CurriculumEntry, error)
	FirstOrCreate() (*model.CurriculumEntry, error)
	FindByPage(offset int, limit int) (result []*model.CurriculumEntry, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICurriculumEntryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithNameAndRole(name string, role string) (result []model.CurriculumEntry, err error)
}

// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
func (c curriculumEntryDo) FilterWithNameAndRole(name string, role string) (result []model.CurriculumEntry, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM curriculum_entries WHERE name = ? ")
	if role != "" {
		params = append(params, role)
		generateSQL.WriteString("AND role = ? ")
	}

	var executeSQL *gorm.DB
	executeSQL = c.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (c curriculumEntryDo) Debug() ICurriculumEntryDo {
	return c.withDO(c.DO.Debug())
}

func (c curriculumEntryDo) WithContext(ctx context.Context) ICurriculumEntryDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c curriculumEntryDo) ReadDB() ICurriculumEntryDo {
	return c.Clauses(dbresolver.Read)
}

func (c curriculumEntryDo) WriteDB() ICurriculumEntryDo {
	return c.Clauses(dbresolver.Write)
}

func (c curriculumEntryDo) Session(config *gorm.Session) ICurriculumEntryDo {
	return c.withDO(c.DO.Session(config))
}

func (c curriculumEntryDo) Clauses(conds ...clause.Expression) ICurriculumEntryDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c curriculumEntryDo) Returning(value interface{}, columns ...string) ICurriculumEntryDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c curriculumEntryDo) Not(conds ...gen.Condition) ICurriculumEntryDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c curriculumEntryDo) Or(conds ...gen.Condition) ICurriculumEntryDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c curriculumEntryDo) Select(conds ...field.Expr) ICurriculumEntryDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c curriculumEntryDo) Where(conds ...gen.Condition) ICurriculumEntryDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c curriculumEntryDo) Order(conds ...field.Expr) ICurriculumEntryDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c curriculumEntryDo) Distinct(cols ...field.Expr) ICurriculumEntryDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c curriculumEntryDo) Omit(cols ...field.Expr) ICurriculumEntryDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c curriculumEntryDo) Join(table schema.Tabler, on ...field.Expr) ICurriculumEntryDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c curriculumEntryDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICurriculumEntryDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c curriculumEntryDo) RightJoin(table schema.Tabler, on ...field.Expr) ICurriculumEntryDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c curriculumEntryDo) Group(cols ...field.Expr) ICurriculumEntryDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c curriculumEntryDo) Having(conds ...gen.Condition) ICurriculumEntryDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c curriculumEntryDo) Limit(limit int) ICurriculumEntryDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c curriculumEntryDo) Offset(offset int) ICurriculumEntryDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c curriculumEntryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICurriculumEntryDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c curriculumEntryDo) Unscoped() ICurriculumEntryDo {
	return c.withDO(c.DO.Unscoped())
}

func (c curriculumEntryDo) Create(values ...*model.CurriculumEntry) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c curriculumEntryDo) CreateInBatches(values []*model.CurriculumEntry, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c curriculumEntryDo) Save(values ...*model.CurriculumEntry) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c curriculumEntryDo) First() (*model.CurriculumEntry, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumEntry), nil
	}
}

func (c curriculumEntryDo) Take() (*model.CurriculumEntry, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumEntry), nil
	}
}

func (c curriculumEntryDo) Last() (*model.CurriculumEntry, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumEntry), nil
	}
}

func (c curriculumEntryDo) Find() ([]*model.CurriculumEntry, error) {
	result, err := c.DO.Find()
	return result.([]*model.CurriculumEntry), err
}

func (c curriculumEntryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CurriculumEntry, err error) {
	buf := make([]*model.CurriculumEntry, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c curriculumEntryDo) FindInBatches(result *[]*model.CurriculumEntry, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c curriculumEntryDo) Attrs(attrs ...field.AssignExpr) ICurriculumEntryDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c curriculumEntryDo) Assign(attrs ...field.AssignExpr) ICurriculumEntryDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c curriculumEntryDo) Joins(fields ...field.RelationField) ICurriculumEntryDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c curriculumEntryDo) Preload(fields ...field.RelationField) ICurriculumEntryDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c curriculumEntryDo) FirstOrInit() (*model.CurriculumEntry, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumEntry), nil
	}
}

func (c curriculumEntryDo) FirstOrCreate() (*model.CurriculumEntry, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CurriculumEntry), nil
	}
}

func (c curriculumEntryDo) FindByPage(offset int, limit int) (result []*model.CurriculumEntry, count int64, err error) {
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

func (c curriculumEntryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c curriculumEntryDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c curriculumEntryDo) Delete(models ...*model.CurriculumEntry) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *curriculumEntryDo) withDO(do gen.Dao) *curriculumEntryDo {
	c.DO = *do.(*gen.DO)
	return c
}
