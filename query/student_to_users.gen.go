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

func newStudentToUser(db *gorm.DB, opts ...gen.DOOption) studentToUser {
	_studentToUser := studentToUser{}

	_studentToUser.studentToUserDo.UseDB(db, opts...)
	_studentToUser.studentToUserDo.UseModel(&model.StudentToUser{})

	tableName := _studentToUser.studentToUserDo.TableName()
	_studentToUser.ALL = field.NewAsterisk(tableName)
	_studentToUser.ID = field.NewField(tableName, "id")
	_studentToUser.CreatedAt = field.NewTime(tableName, "created_at")
	_studentToUser.UpdatedAt = field.NewTime(tableName, "updated_at")
	_studentToUser.DeletedAt = field.NewField(tableName, "deleted_at")
	_studentToUser.GoogleSheetUserName = field.NewString(tableName, "google_sheet_user_name")
	_studentToUser.GoogleSheetPassword = field.NewString(tableName, "google_sheet_password")
	_studentToUser.GoogleSheetSID = field.NewString(tableName, "google_sheet_sid")
	_studentToUser.Name = field.NewString(tableName, "name")
	_studentToUser.UserID = field.NewField(tableName, "user_id")
	_studentToUser.User = studentToUserBelongsToUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "model.User"),
		Role: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("User.Role", "model.Role"),
		},
	}

	_studentToUser.fillFieldMap()

	return _studentToUser
}

type studentToUser struct {
	studentToUserDo

	ALL                 field.Asterisk
	ID                  field.Field
	CreatedAt           field.Time
	UpdatedAt           field.Time
	DeletedAt           field.Field
	GoogleSheetUserName field.String
	GoogleSheetPassword field.String
	GoogleSheetSID      field.String
	Name                field.String
	UserID              field.Field
	User                studentToUserBelongsToUser

	fieldMap map[string]field.Expr
}

func (s studentToUser) Table(newTableName string) *studentToUser {
	s.studentToUserDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s studentToUser) As(alias string) *studentToUser {
	s.studentToUserDo.DO = *(s.studentToUserDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *studentToUser) updateTableName(table string) *studentToUser {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewField(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.GoogleSheetUserName = field.NewString(table, "google_sheet_user_name")
	s.GoogleSheetPassword = field.NewString(table, "google_sheet_password")
	s.GoogleSheetSID = field.NewString(table, "google_sheet_sid")
	s.Name = field.NewString(table, "name")
	s.UserID = field.NewField(table, "user_id")

	s.fillFieldMap()

	return s
}

func (s *studentToUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *studentToUser) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 10)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["google_sheet_user_name"] = s.GoogleSheetUserName
	s.fieldMap["google_sheet_password"] = s.GoogleSheetPassword
	s.fieldMap["google_sheet_sid"] = s.GoogleSheetSID
	s.fieldMap["name"] = s.Name
	s.fieldMap["user_id"] = s.UserID

}

func (s studentToUser) clone(db *gorm.DB) studentToUser {
	s.studentToUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s studentToUser) replaceDB(db *gorm.DB) studentToUser {
	s.studentToUserDo.ReplaceDB(db)
	return s
}

type studentToUserBelongsToUser struct {
	db *gorm.DB

	field.RelationField

	Role struct {
		field.RelationField
	}
}

func (a studentToUserBelongsToUser) Where(conds ...field.Expr) *studentToUserBelongsToUser {
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

func (a studentToUserBelongsToUser) WithContext(ctx context.Context) *studentToUserBelongsToUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a studentToUserBelongsToUser) Session(session *gorm.Session) *studentToUserBelongsToUser {
	a.db = a.db.Session(session)
	return &a
}

func (a studentToUserBelongsToUser) Model(m *model.StudentToUser) *studentToUserBelongsToUserTx {
	return &studentToUserBelongsToUserTx{a.db.Model(m).Association(a.Name())}
}

type studentToUserBelongsToUserTx struct{ tx *gorm.Association }

func (a studentToUserBelongsToUserTx) Find() (result *model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a studentToUserBelongsToUserTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a studentToUserBelongsToUserTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a studentToUserBelongsToUserTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a studentToUserBelongsToUserTx) Clear() error {
	return a.tx.Clear()
}

func (a studentToUserBelongsToUserTx) Count() int64 {
	return a.tx.Count()
}

type studentToUserDo struct{ gen.DO }

type IStudentToUserDo interface {
	gen.SubQuery
	Debug() IStudentToUserDo
	WithContext(ctx context.Context) IStudentToUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStudentToUserDo
	WriteDB() IStudentToUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStudentToUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStudentToUserDo
	Not(conds ...gen.Condition) IStudentToUserDo
	Or(conds ...gen.Condition) IStudentToUserDo
	Select(conds ...field.Expr) IStudentToUserDo
	Where(conds ...gen.Condition) IStudentToUserDo
	Order(conds ...field.Expr) IStudentToUserDo
	Distinct(cols ...field.Expr) IStudentToUserDo
	Omit(cols ...field.Expr) IStudentToUserDo
	Join(table schema.Tabler, on ...field.Expr) IStudentToUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStudentToUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStudentToUserDo
	Group(cols ...field.Expr) IStudentToUserDo
	Having(conds ...gen.Condition) IStudentToUserDo
	Limit(limit int) IStudentToUserDo
	Offset(offset int) IStudentToUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStudentToUserDo
	Unscoped() IStudentToUserDo
	Create(values ...*model.StudentToUser) error
	CreateInBatches(values []*model.StudentToUser, batchSize int) error
	Save(values ...*model.StudentToUser) error
	First() (*model.StudentToUser, error)
	Take() (*model.StudentToUser, error)
	Last() (*model.StudentToUser, error)
	Find() ([]*model.StudentToUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StudentToUser, err error)
	FindInBatches(result *[]*model.StudentToUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StudentToUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStudentToUserDo
	Assign(attrs ...field.AssignExpr) IStudentToUserDo
	Joins(fields ...field.RelationField) IStudentToUserDo
	Preload(fields ...field.RelationField) IStudentToUserDo
	FirstOrInit() (*model.StudentToUser, error)
	FirstOrCreate() (*model.StudentToUser, error)
	FindByPage(offset int, limit int) (result []*model.StudentToUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStudentToUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s studentToUserDo) Debug() IStudentToUserDo {
	return s.withDO(s.DO.Debug())
}

func (s studentToUserDo) WithContext(ctx context.Context) IStudentToUserDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s studentToUserDo) ReadDB() IStudentToUserDo {
	return s.Clauses(dbresolver.Read)
}

func (s studentToUserDo) WriteDB() IStudentToUserDo {
	return s.Clauses(dbresolver.Write)
}

func (s studentToUserDo) Session(config *gorm.Session) IStudentToUserDo {
	return s.withDO(s.DO.Session(config))
}

func (s studentToUserDo) Clauses(conds ...clause.Expression) IStudentToUserDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s studentToUserDo) Returning(value interface{}, columns ...string) IStudentToUserDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s studentToUserDo) Not(conds ...gen.Condition) IStudentToUserDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s studentToUserDo) Or(conds ...gen.Condition) IStudentToUserDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s studentToUserDo) Select(conds ...field.Expr) IStudentToUserDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s studentToUserDo) Where(conds ...gen.Condition) IStudentToUserDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s studentToUserDo) Order(conds ...field.Expr) IStudentToUserDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s studentToUserDo) Distinct(cols ...field.Expr) IStudentToUserDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s studentToUserDo) Omit(cols ...field.Expr) IStudentToUserDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s studentToUserDo) Join(table schema.Tabler, on ...field.Expr) IStudentToUserDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s studentToUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStudentToUserDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s studentToUserDo) RightJoin(table schema.Tabler, on ...field.Expr) IStudentToUserDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s studentToUserDo) Group(cols ...field.Expr) IStudentToUserDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s studentToUserDo) Having(conds ...gen.Condition) IStudentToUserDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s studentToUserDo) Limit(limit int) IStudentToUserDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s studentToUserDo) Offset(offset int) IStudentToUserDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s studentToUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStudentToUserDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s studentToUserDo) Unscoped() IStudentToUserDo {
	return s.withDO(s.DO.Unscoped())
}

func (s studentToUserDo) Create(values ...*model.StudentToUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s studentToUserDo) CreateInBatches(values []*model.StudentToUser, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s studentToUserDo) Save(values ...*model.StudentToUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s studentToUserDo) First() (*model.StudentToUser, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StudentToUser), nil
	}
}

func (s studentToUserDo) Take() (*model.StudentToUser, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StudentToUser), nil
	}
}

func (s studentToUserDo) Last() (*model.StudentToUser, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StudentToUser), nil
	}
}

func (s studentToUserDo) Find() ([]*model.StudentToUser, error) {
	result, err := s.DO.Find()
	return result.([]*model.StudentToUser), err
}

func (s studentToUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StudentToUser, err error) {
	buf := make([]*model.StudentToUser, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s studentToUserDo) FindInBatches(result *[]*model.StudentToUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s studentToUserDo) Attrs(attrs ...field.AssignExpr) IStudentToUserDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s studentToUserDo) Assign(attrs ...field.AssignExpr) IStudentToUserDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s studentToUserDo) Joins(fields ...field.RelationField) IStudentToUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s studentToUserDo) Preload(fields ...field.RelationField) IStudentToUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s studentToUserDo) FirstOrInit() (*model.StudentToUser, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StudentToUser), nil
	}
}

func (s studentToUserDo) FirstOrCreate() (*model.StudentToUser, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StudentToUser), nil
	}
}

func (s studentToUserDo) FindByPage(offset int, limit int) (result []*model.StudentToUser, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s studentToUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s studentToUserDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s studentToUserDo) Delete(models ...*model.StudentToUser) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *studentToUserDo) withDO(do gen.Dao) *studentToUserDo {
	s.DO = *do.(*gen.DO)
	return s
}
