// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"bench/dal/model"
)

func newFollowCount(db *gorm.DB, opts ...gen.DOOption) followCount {
	_followCount := followCount{}

	_followCount.followCountDo.UseDB(db, opts...)
	_followCount.followCountDo.UseModel(&model.FollowCount{})

	tableName := _followCount.followCountDo.TableName()
	_followCount.ALL = field.NewAsterisk(tableName)
	_followCount.ID = field.NewInt64(tableName, "id")
	_followCount.UID = field.NewInt64(tableName, "uid")
	_followCount.FollowCount = field.NewInt64(tableName, "follow_count")
	_followCount.FollowerCount = field.NewInt64(tableName, "follower_count")
	_followCount.Ctime = field.NewTime(tableName, "ctime")
	_followCount.Mtime = field.NewTime(tableName, "mtime")

	_followCount.fillFieldMap()

	return _followCount
}

type followCount struct {
	followCountDo followCountDo

	ALL           field.Asterisk
	ID            field.Int64
	UID           field.Int64
	FollowCount   field.Int64
	FollowerCount field.Int64
	Ctime         field.Time
	Mtime         field.Time

	fieldMap map[string]field.Expr
}

func (f followCount) Table(newTableName string) *followCount {
	f.followCountDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f followCount) As(alias string) *followCount {
	f.followCountDo.DO = *(f.followCountDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *followCount) updateTableName(table string) *followCount {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.UID = field.NewInt64(table, "uid")
	f.FollowCount = field.NewInt64(table, "follow_count")
	f.FollowerCount = field.NewInt64(table, "follower_count")
	f.Ctime = field.NewTime(table, "ctime")
	f.Mtime = field.NewTime(table, "mtime")

	f.fillFieldMap()

	return f
}

func (f *followCount) WithContext(ctx context.Context) *followCountDo {
	return f.followCountDo.WithContext(ctx)
}

func (f followCount) TableName() string { return f.followCountDo.TableName() }

func (f followCount) Alias() string { return f.followCountDo.Alias() }

func (f *followCount) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *followCount) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 6)
	f.fieldMap["id"] = f.ID
	f.fieldMap["uid"] = f.UID
	f.fieldMap["follow_count"] = f.FollowCount
	f.fieldMap["follower_count"] = f.FollowerCount
	f.fieldMap["ctime"] = f.Ctime
	f.fieldMap["mtime"] = f.Mtime
}

func (f followCount) clone(db *gorm.DB) followCount {
	f.followCountDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f followCount) replaceDB(db *gorm.DB) followCount {
	f.followCountDo.ReplaceDB(db)
	return f
}

type followCountDo struct{ gen.DO }

// sql(insert into @@table (uid, follow_count) values (@uid, 1) on duplicate key update follow_count=follow_count+1)
func (f followCountDo) IncrFollowCount(uid int64) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, uid)
	generateSQL.WriteString("insert into follow_count (uid, follow_count) values (?, 1) on duplicate key update follow_count=follow_count+1 ")

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(update @@table set follow_count=follow_count-1 where uid=@uid limit 1)
func (f followCountDo) DecrFollowCount(uid int64) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, uid)
	generateSQL.WriteString("update follow_count set follow_count=follow_count-1 where uid=? limit 1 ")

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(insert into @@table (uid, follower_count) values (@uid, @cnt) on duplicate key update follower_count=follower_count+@cnt)
func (f followCountDo) IncrByFollowerCount(uid int64, cnt int64) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, uid)
	params = append(params, cnt)
	params = append(params, cnt)
	generateSQL.WriteString("insert into follow_count (uid, follower_count) values (?, ?) on duplicate key update follower_count=follower_count+? ")

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(update @@table set follower_count=follower_count-@cnt where uid=@uid limit @cnt)
func (f followCountDo) DecrByFollowerCount(uid int64, cnt int64) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, cnt)
	params = append(params, uid)
	params = append(params, cnt)
	generateSQL.WriteString("update follow_count set follower_count=follower_count-? where uid=? limit ? ")

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(select uid, follow_count, follower_count from @@table where uid in (@uids))
func (f followCountDo) FindUsersRelationCount(uids []int64) (result []*model.FollowCount, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, uids)
	generateSQL.WriteString("select uid, follow_count, follower_count from follow_count where uid in (?) ")

	var executeSQL *gorm.DB
	executeSQL = f.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (f followCountDo) Debug() *followCountDo {
	return f.withDO(f.DO.Debug())
}

func (f followCountDo) WithContext(ctx context.Context) *followCountDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f followCountDo) ReadDB() *followCountDo {
	return f.Clauses(dbresolver.Read)
}

func (f followCountDo) WriteDB() *followCountDo {
	return f.Clauses(dbresolver.Write)
}

func (f followCountDo) Session(config *gorm.Session) *followCountDo {
	return f.withDO(f.DO.Session(config))
}

func (f followCountDo) Clauses(conds ...clause.Expression) *followCountDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f followCountDo) Returning(value interface{}, columns ...string) *followCountDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f followCountDo) Not(conds ...gen.Condition) *followCountDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f followCountDo) Or(conds ...gen.Condition) *followCountDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f followCountDo) Select(conds ...field.Expr) *followCountDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f followCountDo) Where(conds ...gen.Condition) *followCountDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f followCountDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *followCountDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f followCountDo) Order(conds ...field.Expr) *followCountDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f followCountDo) Distinct(cols ...field.Expr) *followCountDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f followCountDo) Omit(cols ...field.Expr) *followCountDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f followCountDo) Join(table schema.Tabler, on ...field.Expr) *followCountDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f followCountDo) LeftJoin(table schema.Tabler, on ...field.Expr) *followCountDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f followCountDo) RightJoin(table schema.Tabler, on ...field.Expr) *followCountDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f followCountDo) Group(cols ...field.Expr) *followCountDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f followCountDo) Having(conds ...gen.Condition) *followCountDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f followCountDo) Limit(limit int) *followCountDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f followCountDo) Offset(offset int) *followCountDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f followCountDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *followCountDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f followCountDo) Unscoped() *followCountDo {
	return f.withDO(f.DO.Unscoped())
}

func (f followCountDo) Create(values ...*model.FollowCount) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f followCountDo) CreateInBatches(values []*model.FollowCount, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f followCountDo) Save(values ...*model.FollowCount) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f followCountDo) First() (*model.FollowCount, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowCount), nil
	}
}

func (f followCountDo) Take() (*model.FollowCount, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowCount), nil
	}
}

func (f followCountDo) Last() (*model.FollowCount, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowCount), nil
	}
}

func (f followCountDo) Find() ([]*model.FollowCount, error) {
	result, err := f.DO.Find()
	return result.([]*model.FollowCount), err
}

func (f followCountDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FollowCount, err error) {
	buf := make([]*model.FollowCount, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f followCountDo) FindInBatches(result *[]*model.FollowCount, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f followCountDo) Attrs(attrs ...field.AssignExpr) *followCountDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f followCountDo) Assign(attrs ...field.AssignExpr) *followCountDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f followCountDo) Joins(fields ...field.RelationField) *followCountDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f followCountDo) Preload(fields ...field.RelationField) *followCountDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f followCountDo) FirstOrInit() (*model.FollowCount, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowCount), nil
	}
}

func (f followCountDo) FirstOrCreate() (*model.FollowCount, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FollowCount), nil
	}
}

func (f followCountDo) FindByPage(offset int, limit int) (result []*model.FollowCount, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f followCountDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f followCountDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f followCountDo) Delete(models ...*model.FollowCount) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *followCountDo) withDO(do gen.Dao) *followCountDo {
	f.DO = *do.(*gen.DO)
	return f
}
