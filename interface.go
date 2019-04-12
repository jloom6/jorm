package jorm

//go:generate retool do mockgen -destination=mocks/jorm.go -package=mocks github.com/jloom6/jorm Interface,Row,Rows

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Interface contains all of the funcs a *gorm.DB has
type Interface interface {
	// Adds context propagation
	WithContext(ctx context.Context) Interface
	// Allow field values to be accessed via function calls
	Value() interface{}
	Error() error
	RowsAffected() int64
	// Wrap the functions
	New() Interface
	Close() error
	DB() *sql.DB
	CommonDB() gorm.SQLCommon
	Dialect() gorm.Dialect
	Callback() *gorm.Callback
	SetLogger(log mysql.Logger)
	LogMode(enable bool) Interface
	BlockGlobalUpdate(enable bool) Interface
	HasBlockGlobalUpdate() bool
	SingularTable(enable bool)
	Where(query interface{}, args ...interface{}) Interface
	Or(query interface{}, args ...interface{}) Interface
	Not(query interface{}, args ...interface{}) Interface
	Limit(limit interface{}) Interface
	Offset(offset interface{}) Interface
	Order(value interface{}, reorder ...bool) Interface
	Select(query interface{}, args ...interface{}) Interface
	Omit(columns ...string) Interface
	Group(query string) Interface
	Having(query interface{}, values ...interface{}) Interface
	Joins(query string, args ...interface{}) Interface
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) Interface
	Unscoped() Interface
	Assign(attrs ...interface{}) Interface
	Attrs(attrs ...interface{}) Interface
	First(out interface{}, where ...interface{}) Interface
	Take(out interface{}, where ...interface{}) Interface
	Last(out interface{}, where ...interface{}) Interface
	Find(out interface{}, where ...interface{}) Interface
	Scan(dest interface{}) Interface
	Row() Row
	Rows() (Rows, error)
	ScanRows(rows *sql.Rows, result interface{}) error
	Pluck(column string, value interface{}) Interface
	Count(value interface{}) Interface
	Related(value interface{}, foreignKeys ...string) Interface
	FirstOrInit(out interface{}, where ...interface{}) Interface
	FirstOrCreate(out interface{}, where ...interface{}) Interface
	Update(attrs ...interface{}) Interface
	Updates(values interface{}, ignoreProtectedAttrs ...bool) Interface
	UpdateColumn(attrs ...interface{}) Interface
	UpdateColumns(values interface{}) Interface
	Save(value interface{}) Interface
	Create(value interface{}) Interface
	Delete(value interface{}, where ...interface{}) Interface
	Raw(sql string, values ...interface{}) Interface
	Exec(sql string, values ...interface{}) Interface
	Model(value interface{}) Interface
	Table(name string) Interface
	Debug() Interface
	Begin() Interface
	Commit() Interface
	Rollback() Interface
	NewRecord(value interface{}) bool
	RecordNotFound() bool
	CreateTable(models ...interface{}) Interface
	DropTable(values ...interface{}) Interface
	DropTableIfExists(values ...interface{}) Interface
	HasTable(value interface{}) bool
	AutoMigrate(values ...interface{}) Interface
	ModifyColumn(column string, typ string) Interface
	DropColumn(column string) Interface
	AddIndex(indexName string, columns ...string) Interface
	AddUniqueIndex(indexName string, columns ...string) Interface
	RemoveIndex(indexName string) Interface
	AddForeignKey(field string, dest string, onDelete string, onUpdate string) Interface
	RemoveForeignKey(field string, dest string) Interface
	Association(column string) *gorm.Association
	Preload(column string, conditions ...interface{}) Interface
	Set(name string, value interface{}) Interface
	InstantSet(name string, value interface{}) Interface
	Get(name string) (value interface{}, ok bool)
	SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface)
	AddError(err error) error
	GetErrors() []error
}

// Row is an interface wrapper for sql.Row
type Row interface {
	Scan(dest ...interface{}) error
}

// Rows is an interface wrapper for sql.Rows
type Rows interface {
	Next() bool
	NextResultSet() bool
	Err() error
	Columns() ([]string, error)
	ColumnTypes() ([]*sql.ColumnType, error)
	Scan(dest ...interface{}) error
	Close() error
}
