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

func newParentUserActivating(db *gorm.DB, opts ...gen.DOOption) parentUserActivating {
	_parentUserActivating := parentUserActivating{}

	_parentUserActivating.parentUserActivatingDo.UseDB(db, opts...)
	_parentUserActivating.parentUserActivatingDo.UseModel(&model.ParentUserActivating{})

	tableName := _parentUserActivating.parentUserActivatingDo.TableName()
	_parentUserActivating.ALL = field.NewAsterisk(tableName)
	_parentUserActivating.ID = field.NewField(tableName, "id")
	_parentUserActivating.CreatedAt = field.NewTime(tableName, "created_at")
	_parentUserActivating.UpdatedAt = field.NewTime(tableName, "updated_at")
	_parentUserActivating.DeletedAt = field.NewField(tableName, "deleted_at")
	_parentUserActivating.FullName = field.NewString(tableName, "full_name")
	_parentUserActivating.UserName = field.NewString(tableName, "user_name")
	_parentUserActivating.Password = field.NewString(tableName, "password")
	_parentUserActivating.ContactNumber = field.NewString(tableName, "contact_number")
	_parentUserActivating.Email = field.NewString(tableName, "email")
	_parentUserActivating.ActivationKey = field.NewString(tableName, "activation_key")

	_parentUserActivating.fillFieldMap()

	return _parentUserActivating
}

type parentUserActivating struct {
	parentUserActivatingDo

	ALL           field.Asterisk
	ID            field.Field
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	FullName      field.String
	UserName      field.String
	Password      field.String
	ContactNumber field.String
	Email         field.String
	ActivationKey field.String

	fieldMap map[string]field.Expr
}

func (p parentUserActivating) Table(newTableName string) *parentUserActivating {
	p.parentUserActivatingDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p parentUserActivating) As(alias string) *parentUserActivating {
	p.parentUserActivatingDo.DO = *(p.parentUserActivatingDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *parentUserActivating) updateTableName(table string) *parentUserActivating {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewField(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.FullName = field.NewString(table, "full_name")
	p.UserName = field.NewString(table, "user_name")
	p.Password = field.NewString(table, "password")
	p.ContactNumber = field.NewString(table, "contact_number")
	p.Email = field.NewString(table, "email")
	p.ActivationKey = field.NewString(table, "activation_key")

	p.fillFieldMap()

	return p
}

func (p *parentUserActivating) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *parentUserActivating) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 10)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["full_name"] = p.FullName
	p.fieldMap["user_name"] = p.UserName
	p.fieldMap["password"] = p.Password
	p.fieldMap["contact_number"] = p.ContactNumber
	p.fieldMap["email"] = p.Email
	p.fieldMap["activation_key"] = p.ActivationKey
}

func (p parentUserActivating) clone(db *gorm.DB) parentUserActivating {
	p.parentUserActivatingDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p parentUserActivating) replaceDB(db *gorm.DB) parentUserActivating {
	p.parentUserActivatingDo.ReplaceDB(db)
	return p
}

type parentUserActivatingDo struct{ gen.DO }

type IParentUserActivatingDo interface {
	gen.SubQuery
	Debug() IParentUserActivatingDo
	WithContext(ctx context.Context) IParentUserActivatingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IParentUserActivatingDo
	WriteDB() IParentUserActivatingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IParentUserActivatingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IParentUserActivatingDo
	Not(conds ...gen.Condition) IParentUserActivatingDo
	Or(conds ...gen.Condition) IParentUserActivatingDo
	Select(conds ...field.Expr) IParentUserActivatingDo
	Where(conds ...gen.Condition) IParentUserActivatingDo
	Order(conds ...field.Expr) IParentUserActivatingDo
	Distinct(cols ...field.Expr) IParentUserActivatingDo
	Omit(cols ...field.Expr) IParentUserActivatingDo
	Join(table schema.Tabler, on ...field.Expr) IParentUserActivatingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IParentUserActivatingDo
	RightJoin(table schema.Tabler, on ...field.Expr) IParentUserActivatingDo
	Group(cols ...field.Expr) IParentUserActivatingDo
	Having(conds ...gen.Condition) IParentUserActivatingDo
	Limit(limit int) IParentUserActivatingDo
	Offset(offset int) IParentUserActivatingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IParentUserActivatingDo
	Unscoped() IParentUserActivatingDo
	Create(values ...*model.ParentUserActivating) error
	CreateInBatches(values []*model.ParentUserActivating, batchSize int) error
	Save(values ...*model.ParentUserActivating) error
	First() (*model.ParentUserActivating, error)
	Take() (*model.ParentUserActivating, error)
	Last() (*model.ParentUserActivating, error)
	Find() ([]*model.ParentUserActivating, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ParentUserActivating, err error)
	FindInBatches(result *[]*model.ParentUserActivating, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ParentUserActivating) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IParentUserActivatingDo
	Assign(attrs ...field.AssignExpr) IParentUserActivatingDo
	Joins(fields ...field.RelationField) IParentUserActivatingDo
	Preload(fields ...field.RelationField) IParentUserActivatingDo
	FirstOrInit() (*model.ParentUserActivating, error)
	FirstOrCreate() (*model.ParentUserActivating, error)
	FindByPage(offset int, limit int) (result []*model.ParentUserActivating, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IParentUserActivatingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FilterWithNameAndRole(name string, role string) (result []model.ParentUserActivating, err error)
}

// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
func (p parentUserActivatingDo) FilterWithNameAndRole(name string, role string) (result []model.ParentUserActivating, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	generateSQL.WriteString("SELECT * FROM parent_user_activatings WHERE name = ? ")
	if role != "" {
		params = append(params, role)
		generateSQL.WriteString("AND role = ? ")
	}

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p parentUserActivatingDo) Debug() IParentUserActivatingDo {
	return p.withDO(p.DO.Debug())
}

func (p parentUserActivatingDo) WithContext(ctx context.Context) IParentUserActivatingDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p parentUserActivatingDo) ReadDB() IParentUserActivatingDo {
	return p.Clauses(dbresolver.Read)
}

func (p parentUserActivatingDo) WriteDB() IParentUserActivatingDo {
	return p.Clauses(dbresolver.Write)
}

func (p parentUserActivatingDo) Session(config *gorm.Session) IParentUserActivatingDo {
	return p.withDO(p.DO.Session(config))
}

func (p parentUserActivatingDo) Clauses(conds ...clause.Expression) IParentUserActivatingDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p parentUserActivatingDo) Returning(value interface{}, columns ...string) IParentUserActivatingDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p parentUserActivatingDo) Not(conds ...gen.Condition) IParentUserActivatingDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p parentUserActivatingDo) Or(conds ...gen.Condition) IParentUserActivatingDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p parentUserActivatingDo) Select(conds ...field.Expr) IParentUserActivatingDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p parentUserActivatingDo) Where(conds ...gen.Condition) IParentUserActivatingDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p parentUserActivatingDo) Order(conds ...field.Expr) IParentUserActivatingDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p parentUserActivatingDo) Distinct(cols ...field.Expr) IParentUserActivatingDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p parentUserActivatingDo) Omit(cols ...field.Expr) IParentUserActivatingDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p parentUserActivatingDo) Join(table schema.Tabler, on ...field.Expr) IParentUserActivatingDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p parentUserActivatingDo) LeftJoin(table schema.Tabler, on ...field.Expr) IParentUserActivatingDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p parentUserActivatingDo) RightJoin(table schema.Tabler, on ...field.Expr) IParentUserActivatingDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p parentUserActivatingDo) Group(cols ...field.Expr) IParentUserActivatingDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p parentUserActivatingDo) Having(conds ...gen.Condition) IParentUserActivatingDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p parentUserActivatingDo) Limit(limit int) IParentUserActivatingDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p parentUserActivatingDo) Offset(offset int) IParentUserActivatingDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p parentUserActivatingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IParentUserActivatingDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p parentUserActivatingDo) Unscoped() IParentUserActivatingDo {
	return p.withDO(p.DO.Unscoped())
}

func (p parentUserActivatingDo) Create(values ...*model.ParentUserActivating) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p parentUserActivatingDo) CreateInBatches(values []*model.ParentUserActivating, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p parentUserActivatingDo) Save(values ...*model.ParentUserActivating) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p parentUserActivatingDo) First() (*model.ParentUserActivating, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ParentUserActivating), nil
	}
}

func (p parentUserActivatingDo) Take() (*model.ParentUserActivating, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ParentUserActivating), nil
	}
}

func (p parentUserActivatingDo) Last() (*model.ParentUserActivating, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ParentUserActivating), nil
	}
}

func (p parentUserActivatingDo) Find() ([]*model.ParentUserActivating, error) {
	result, err := p.DO.Find()
	return result.([]*model.ParentUserActivating), err
}

func (p parentUserActivatingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ParentUserActivating, err error) {
	buf := make([]*model.ParentUserActivating, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p parentUserActivatingDo) FindInBatches(result *[]*model.ParentUserActivating, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p parentUserActivatingDo) Attrs(attrs ...field.AssignExpr) IParentUserActivatingDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p parentUserActivatingDo) Assign(attrs ...field.AssignExpr) IParentUserActivatingDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p parentUserActivatingDo) Joins(fields ...field.RelationField) IParentUserActivatingDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p parentUserActivatingDo) Preload(fields ...field.RelationField) IParentUserActivatingDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p parentUserActivatingDo) FirstOrInit() (*model.ParentUserActivating, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ParentUserActivating), nil
	}
}

func (p parentUserActivatingDo) FirstOrCreate() (*model.ParentUserActivating, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ParentUserActivating), nil
	}
}

func (p parentUserActivatingDo) FindByPage(offset int, limit int) (result []*model.ParentUserActivating, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p parentUserActivatingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p parentUserActivatingDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p parentUserActivatingDo) Delete(models ...*model.ParentUserActivating) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *parentUserActivatingDo) withDO(do gen.Dao) *parentUserActivatingDo {
	p.DO = *do.(*gen.DO)
	return p
}
