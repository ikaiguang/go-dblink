package dblink

import (
	"fmt"
	"github.com/ikaiguang/go-dblink/config"
	"github.com/jinzhu/gorm"
	"strings"
)

// ConnHandler new db connection func
type ConnHandler func() (*gorm.DB, error)

// NewDBConn : orm conn
func NewDBConn(f ConnHandler) *gorm.DB {
	dbConn, err := f()
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
	// database config
	cfg := configs.GetOptionConfigHandler()

	// set table prefix
	SetTablePrefix()

	// set max open conn
	SetMaxOpenConn(dbConn, cfg)

	// set max idle conn
	SetMaxIdleConn(dbConn, cfg)

	// set conn max lifetime
	SetConnMaxLifetime(dbConn, cfg)

	// debug
	return SetOrmDebug(dbConn, cfg)
}

// SetTablePrefix : set table prefix
func SetTablePrefix() {
	// empty
	if len(configs.TablePrefix) == 0 {
		return
	}

	// rewrite handler
	gorm.DefaultTableNameHandler = func(db *gorm.DB, tableName string) string {
		if !strings.HasPrefix(tableName, configs.TablePrefix) {
			return configs.TablePrefix + tableName
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
