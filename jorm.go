package jorm

import (
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/smacker/opentracing-gorm"
)

// DB is a wrapper struct around a *gorm.DB
type DB struct {
	db *gorm.DB
}

// NewDB returns a new interface wrapper around the given *gorm.DB
func NewDB(db *gorm.DB) *DB {
	return &DB{db: db}
}

// WithContext returns a clone of the current db with parent Span ID set to that of the context
func (db *DB) WithContext(ctx context.Context) Interface {
	return &DB{db: otgorm.SetSpanToGorm(ctx, db.db)}
}

// Value is a wrapper function for the Value field
func (db *DB) Value() interface{} {
	return db.db.Value
}

// Error is a wrapper function for the Error field
func (db *DB) Error() error {
	return db.db.Error
}

// RowsAffected is a wrapper function for the RowsAffected field
func (db *DB) RowsAffected() int64 {
	return db.db.RowsAffected
}

// New clone a new db connection without search conditions
func (db *DB) New() Interface {
	return &DB{db: db.db.New()}
}

// Close close current db connection.  If database connection is not an io.Closer, returns an error.
func (db *DB) Close() error {
	return db.db.Close()
}

// DB get `*sql.DB` from current connection
// If the underlying database connection is not a *sql.DB, returns nil
func (db *DB) DB() *sql.DB {
	return db.db.DB()
}

// CommonDB return the underlying `*sql.DB` or `*sql.Tx` instance, mainly intended to allow coexistence with legacy non-GORM code.
func (db *DB) CommonDB() gorm.SQLCommon {
	return db.db.CommonDB()
}

// Dialect get dialect
func (db *DB) Dialect() gorm.Dialect {
	return db.db.Dialect()
}

// Callback return `Callbacks` container, you could add/change/delete callbacks with it
//     db.Callback().Create().Register("update_created_at", updateCreated)
// Refer https://jinzhu.github.io/gorm/development.html#callbacks
func (db *DB) Callback() *gorm.Callback {
	return db.db.Callback()
}

// SetLogger replace default logger
func (db *DB) SetLogger(log mysql.Logger) {
	db.db.SetLogger(log)
}

// LogMode set log mode, `true` for detailed logs, `false` for no log, default, will only print error logs
func (db *DB) LogMode(enable bool) Interface {
	return &DB{db: db.db.LogMode(enable)}
}

// BlockGlobalUpdate if true, generates an error on update/delete without where clause.
// This is to prevent eventual error with empty objects updates/deletions
func (db *DB) BlockGlobalUpdate(enable bool) Interface {
	return &DB{db: db.db.BlockGlobalUpdate(enable)}
}

// HasBlockGlobalUpdate return state of block
func (db *DB) HasBlockGlobalUpdate() bool {
	return db.db.HasBlockGlobalUpdate()
}

// SingularTable use singular table by default
func (db *DB) SingularTable(enable bool) {
	db.db.SingularTable(enable)
}

// Where return a new relation, filter records with given conditions, accepts `map`, `struct` or `string` as conditions, refer http://jinzhu.github.io/gorm/crud.html#query
func (db *DB) Where(query interface{}, args ...interface{}) Interface {
	return &DB{db: db.db.Where(query, args...)}
}

// Or filter records that match before conditions or this one, similar to `Where`
func (db *DB) Or(query interface{}, args ...interface{}) Interface {
	return &DB{db: db.db.Or(query, args...)}
}

// Not filter records that don't match current conditions, similar to `Where`
func (db *DB) Not(query interface{}, args ...interface{}) Interface {
	return &DB{db: db.db.Not(query, args...)}
}

// Limit specify the number of records to be retrieved
func (db *DB) Limit(limit interface{}) Interface {
	return &DB{db: db.db.Limit(limit)}
}

// Offset specify the number of records to skip before starting to return the records
func (db *DB) Offset(offset interface{}) Interface {
	return &DB{db: db.db.Offset(offset)}
}

// Order specify order when retrieve records from database, set reorder to `true` to overwrite defined conditions
//     db.Order("name DESC")
//     db.Order("name DESC", true) // reorder
//     db.Order(gorm.Expr("name = ? DESC", "first")) // sql expression
func (db *DB) Order(value interface{}, reorder ...bool) Interface {
	return &DB{db: db.db.Order(value, reorder...)}
}

// Select specify fields that you want to retrieve from database when querying, by default, will select all fields;
// When creating/updating, specify fields that you want to save to database
func (db *DB) Select(query interface{}, args ...interface{}) Interface {
	return &DB{db: db.db.Select(query, args...)}
}

// Omit specify fields that you want to ignore when saving to database for creating, updating
func (db *DB) Omit(columns ...string) Interface {
	return &DB{db: db.db.Omit(columns...)}
}

// Group specify the group method on the find
func (db *DB) Group(query string) Interface {
	return &DB{db: db.db.Group(query)}
}

// Having specify HAVING conditions for GROUP BY
func (db *DB) Having(query interface{}, values ...interface{}) Interface {
	return &DB{db: db.db.Having(query, values...)}
}

// Joins specify Joins conditions
//     db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
func (db *DB) Joins(query string, args ...interface{}) Interface {
	return &DB{db: db.db.Joins(	query, args...)}
}

// Scopes pass current database connection to arguments `func(*DB) *DB`, which could be used to add conditions dynamically
//     func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
//         return db.Where("amount > ?", 1000)
//     }
//
//     func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
//         return func (db *gorm.DB) *gorm.DB {
//             return db.Scopes(AmountGreaterThan1000).Where("status in (?)", status)
//         }
//     }
//
//     db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
// Refer https://jinzhu.github.io/gorm/crud.html#scopes
func (db *DB) Scopes(funcs ...func(*gorm.DB) *gorm.DB) Interface {
	return &DB{db: db.db.Scopes(funcs...)}
}

// Unscoped return all record including deleted record, refer Soft Delete https://jinzhu.github.io/gorm/crud.html#soft-delete
func (db *DB) Unscoped() Interface {
	return &DB{db: db.db.Unscoped()}
}

// Attrs initialize struct with argument if record not found with `FirstOrInit` https://jinzhu.github.io/gorm/crud.html#firstorinit or `FirstOrCreate` https://jinzhu.github.io/gorm/crud.html#firstorcreate
func (db *DB) Attrs(attrs ...interface{}) Interface {
	return &DB{db: db.db.Attrs(attrs...)}
}

// Assign assign result with argument regardless it is found or not with `FirstOrInit` https://jinzhu.github.io/gorm/crud.html#firstorinit or `FirstOrCreate` https://jinzhu.github.io/gorm/crud.html#firstorcreate
func (db *DB) Assign(attrs ...interface{}) Interface {
	return &DB{db: db.db.Assign(attrs...)}
}

// First find first record that match given conditions, order by primary key
func (db *DB) First(out interface{}, where ...interface{}) Interface {
	return &DB{db: db.db.First(out, where...)}
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (db *DB) Take(out interface{}, where ...interface{}) Interface {
	return &DB{db: db.db.Take(out, where...)}
}

// Last find last record that match given conditions, order by primary key
func (db *DB) Last(out interface{}, where ...interface{}) Interface {
	return &DB{db: db.db.Last(out, where...)}
}

// Find find records that match given conditions
func (db *DB) Find(out interface{}, where ...interface{}) Interface {
	return &DB{db: db.db.Find(out, where...)}
}

// Scan scan value to a struct
func (db *DB) Scan(dest interface{}) Interface {
	return &DB{db: db.db.Scan(dest)}
}

// Row return `*sql.Row` with given conditions
func (db *DB) Row() Row {
	return db.db.Row()
}

// Rows return `*sql.Rows` with given conditions
func (db *DB) Rows() (Rows, error) {
	return db.db.Rows()
}

// ScanRows scan `*sql.Rows` to give struct
func (db *DB) ScanRows(rows *sql.Rows, result interface{}) error {
	return db.db.ScanRows(rows, result)
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Find(&users).Pluck("age", &ages)
func (db *DB) Pluck(column string, value interface{}) Interface {
	return &DB{db: db.db.Pluck(column, value)}
}

// Count get how many records for a model
func (db *DB) Count(value interface{}) Interface {
	return &DB{db: db.db.Count(value)}
}

// Related get related associations
func (db *DB) Related(value interface{}, foreignKeys ...string) Interface {
	return &DB{db: db.db.Related(value, foreignKeys...)}
}

// FirstOrInit find first matched record or initialize a new one with given conditions (only works with struct, map conditions)
// https://jinzhu.github.io/gorm/crud.html#firstorinit
func (db *DB) FirstOrInit(out interface{}, where ...interface{}) Interface {
	return &DB{db: db.db.FirstOrInit(out, where...)}
}

// FirstOrCreate find first matched record or create a new one with given conditions (only works with struct, map conditions)
// https://jinzhu.github.io/gorm/crud.html#firstorcreate
func (db *DB) FirstOrCreate(out interface{}, where ...interface{}) Interface {
	return &DB{db: db.db.FirstOrCreate(out, where...)}
}

// Update update attributes with callbacks, refer: https://jinzhu.github.io/gorm/crud.html#update
func (db *DB) Update(attrs ...interface{}) Interface {
	return &DB{db: db.db.Update(attrs...)}
}

// Updates update attributes with callbacks, refer: https://jinzhu.github.io/gorm/crud.html#update
func (db *DB) Updates(values interface{}, ignoreProtectedAttrs ...bool) Interface {
	return &DB{db: db.db.Updates(values, ignoreProtectedAttrs...)}
}

// UpdateColumn update attributes without callbacks, refer: https://jinzhu.github.io/gorm/crud.html#update
func (db *DB) UpdateColumn(attrs ...interface{}) Interface {
	return &DB{db: db.db.UpdateColumn(attrs...)}
}

// UpdateColumns update attributes without callbacks, refer: https://jinzhu.github.io/gorm/crud.html#update
func (db *DB) UpdateColumns(values interface{}) Interface {
	return &DB{db: db.db.UpdateColumns(values)}
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (db *DB) Save(value interface{}) Interface {
	return &DB{db: db.db.Save(value)}
}

// Create insert the value into database
func (db *DB) Create(value interface{}) Interface {
	return &DB{db: db.db.Create(value)}
}

// Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition
func (db *DB) Delete(value interface{}, where ...interface{}) Interface {
	return &DB{db: db.db.Delete(value, where...)}
}

// Raw use raw sql as conditions, won't run it unless invoked by other methods
//    db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)
func (db *DB) Raw(sql string, values ...interface{}) Interface {
	return &DB{db: db.db.Raw(sql, values...)}
}

// Exec execute raw sql
func (db *DB) Exec(sql string, values ...interface{}) Interface {
	return &DB{db: db.db.Exec(sql, values...)}
}

// Model specify the model you would like to run db operations
//    // update all users's name to `hello`
//    db.Model(&User{}).Update("name", "hello")
//    // if user's primary key is non-blank, will use it as condition, then will only update the user's name to `hello`
//    db.Model(&user).Update("name", "hello")
func (db *DB) Model(value interface{}) Interface {
	return &DB{db: db.db.Model(value)}
}

// Table specify the table you would like to run db operations
func (db *DB) Table(name string) Interface {
	return &DB{db: db.db.Table(name)}
}

// Debug start debug mode
func (db *DB) Debug() Interface {
	return &DB{db: db.db.Debug()}
}

// Begin begin a transaction
func (db *DB) Begin() Interface {
	return &DB{db: db.db.Begin()}
}

// Commit commit a transaction
func (db *DB) Commit() Interface {
	return &DB{db: db.db.Commit()}
}

// Rollback rollback a transaction
func (db *DB) Rollback() Interface {
	return &DB{db: db.db.Rollback()}
}

// NewRecord check if value's primary key is blank
func (db *DB) NewRecord(value interface{}) bool {
	return db.db.NewRecord(value)
}

// RecordNotFound check if returning ErrRecordNotFound error
func (db *DB) RecordNotFound() bool {
	return db.db.RecordNotFound()
}

// CreateTable create table for models
func (db *DB) CreateTable(models ...interface{}) Interface {
	return &DB{db: db.db.CreateTable(models...)}
}

// DropTable drop table for models
func (db *DB) DropTable(values ...interface{}) Interface {
	return &DB{db: db.db.DropTable(values...)}
}

// DropTableIfExists drop table if it is exist
func (db *DB) DropTableIfExists(values ...interface{}) Interface {
	return &DB{db: db.db.DropTableIfExists(values...)}
}

// HasTable check has table or not
func (db *DB) HasTable(value interface{}) bool {
	return db.db.HasTable(value)
}

// AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data
func (db *DB) AutoMigrate(values ...interface{}) Interface {
	return &DB{db: db.db.AutoMigrate(values)}
}

// ModifyColumn modify column to type
func (db *DB) ModifyColumn(column string, typ string) Interface {
	return &DB{db: db.db.ModifyColumn(column, typ)}
}

// DropColumn drop a column
func (db *DB) DropColumn(column string) Interface {
	return &DB{db: db.db.DropColumn(column)}
}

// AddIndex add index for columns with given name
func (db *DB) AddIndex(indexName string, columns ...string) Interface {
	return &DB{db: db.db.AddIndex(indexName, columns...)}
}

// AddUniqueIndex add unique index for columns with given name
func (db *DB) AddUniqueIndex(indexName string, columns ...string) Interface {
	return &DB{db: db.db.AddUniqueIndex(indexName, columns...)}
}

// RemoveIndex remove index with name
func (db *DB) RemoveIndex(indexName string) Interface {
	return &DB{db: db.db.RemoveIndex(indexName)}
}

// AddForeignKey Add foreign key to the given scope, e.g:
//     db.Model(&User{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")
func (db *DB) AddForeignKey(field string, dest string, onDelete string, onUpdate string) Interface {
	return &DB{db: db.db.AddForeignKey(field, dest, onDelete, onUpdate)}
}

// RemoveForeignKey Remove foreign key from the given scope, e.g:
//     db.Model(&User{}).RemoveForeignKey("city_id", "cities(id)")
func (db *DB) RemoveForeignKey(field string, dest string) Interface {
	return &DB{db: db.db.RemoveForeignKey(field, dest)}
}

// Association start `Association Mode` to handler relations things easir in that mode, refer: https://jinzhu.github.io/gorm/associations.html#association-mode
func (db *DB) Association(column string) *gorm.Association {
	return db.db.Association(column)
}

// Preload preload associations with given conditions
//    db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
func (db *DB) Preload(column string, conditions ...interface{}) Interface {
	return &DB{db: db.db.Preload(column, conditions...)}
}

// Set set setting by name, which could be used in callbacks, will clone a new db, and update its setting
func (db *DB) Set(name string, value interface{}) Interface {
	return &DB{db: db.db.Set(name, value)}
}

// InstantSet instant set setting, will affect current db
func (db *DB) InstantSet(name string, value interface{}) Interface {
	return &DB{db: db.db.InstantSet(name, value)}
}

// Get get setting by name
func (db *DB) Get(name string) (value interface{}, ok bool) {
	return db.db.Get(name)
}

// SetJoinTableHandler set a model's join table handler for a relation
func (db *DB) SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface) {
	db.db.SetJoinTableHandler(source, column, handler)
}

// AddError add error to the db
func (db *DB) AddError(err error) error {
	return db.db.AddError(err)
}

// GetErrors get happened errors from the db
func (db *DB) GetErrors() []error {
	return db.db.GetErrors()
}
