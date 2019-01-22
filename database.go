package dblink

import (
	"fmt"
	"github.com/ikaiguang/go-dblink/config"
	"github.com/ikaiguang/go-dblink/mysql"
	"github.com/ikaiguang/go-dblink/postgres"
	"github.com/jinzhu/gorm"
	"strings"
)

// database driver
const (
	DriverMysql    = "mysql"    // mysql
	DriverPostgres = "postgres" // postgres
)

// NewDBConn : db conn
func NewDBConn() *gorm.DB {
	var dbConn *gorm.DB
	var err error

	// auth config
	authCfg := configs.GetAuthConfigHandler()

	// open db connection
	switch authCfg.Driver {
	case DriverMysql:
		// mysql
		dbConn, err = mysql.NewDBConn(authCfg)

	case DriverPostgres:
		// postgres
		dbConn, err = postgres.NewDBConn(authCfg)

	default:
		panic(fmt.Errorf("invalid database driver"))
	}
	// db connection error
	if err != nil {
		panic(fmt.Errorf("dblink new connection error : %v", err))
	}

	// ping
	if err := dbConn.DB().Ping(); err != nil {
		panic(fmt.Errorf("ping database connection fail : %v", err))
	}
	return SetDBConnOptions(dbConn)
}

// SetDBConnOptions set dbConn options
func SetDBConnOptions(dbConn *gorm.DB) *gorm.DB {
	// option config
	optionCfg := configs.GetOptionConfigHandler()

	// set table prefix
	SetTablePrefix()

	// set max open conn
	SetMaxOpenConn(dbConn, optionCfg)

	// set max idle conn
	SetMaxIdleConn(dbConn, optionCfg)

	// set conn max lifetime
	SetConnMaxLifetime(dbConn, optionCfg)

	// debug
	return SetOrmDebug(dbConn, optionCfg)
}

// TablePrefix table prefix
func TablePrefix() string {
	return configs.GetTablePrefixHandler()
}

// SetTablePrefix : set table prefix
func SetTablePrefix() {
	tablePrefix := TablePrefix()
	// empty
	if len(tablePrefix) == 0 {
		return
	}

	// rewrite handler
	gorm.DefaultTableNameHandler = func(db *gorm.DB, tableName string) string {
		if !strings.HasPrefix(tableName, tablePrefix) {
			return tablePrefix + tableName
		}
		return tableName
	}
}

// SetOrmDebug : print debug
// LogMode set log mode, `true` for detailed logs, `false` for no log,
// default, will only print error logs
func SetOrmDebug(dbConn *gorm.DB, cfg *configs.OptionConfig) *gorm.DB {
	return dbConn.LogMode(cfg.Debug)
}

// SetMaxOpenConn : set max open conn
// SetMaxOpenConns sets the maximum number of open connections to the database.
//
// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
// MaxIdleConns, then MaxIdleConns will be reduced to match the new
// MaxOpenConns limit.
//
// If n <= 0, then there is no limit on the number of open connections.
// The default is 0 (unlimited).
func SetMaxOpenConn(dbConn *gorm.DB, cfg *configs.OptionConfig) {
	dbConn.DB().SetMaxOpenConns(cfg.MaxOpen)
}

// SetMaxIdleConn : set max idle conn
// SetMaxIdleConns sets the maximum number of connections in the idle
// connection pool.
//
// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns,
// then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.
//
// If n <= 0, no idle connections are retained.
//
// The default max idle connections is currently 2. This may change in
// a future release.
func SetMaxIdleConn(dbConn *gorm.DB, cfg *configs.OptionConfig) {
	dbConn.DB().SetMaxIdleConns(cfg.MaxIdle)
}

// SetConnMaxLifetime : conn max lifetime
// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
//
// Expired connections may be closed lazily before reuse.
//
// If d <= 0, connections are reused forever.
func SetConnMaxLifetime(dbConn *gorm.DB, cfg *configs.OptionConfig) {
	dbConn.DB().SetConnMaxLifetime(cfg.MaxLifetime)
}
