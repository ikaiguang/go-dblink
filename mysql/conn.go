package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // driver
	"os"
	"strconv"
	"strings"
	"time"
)

//os.Setenv("MysqlDebug", "true")
//os.Setenv("MysqlTablePrefix", "ikg_")
//os.Setenv("MysqlUser", "root")
//os.Setenv("MysqlPassword", "Mysql.123456")
//os.Setenv("MysqlNet", "tcp")
//os.Setenv("MysqlAddr", "127.0.0.1:3306")
//os.Setenv("MysqlName", "test")
//os.Setenv("MysqlParameters", "charset=utf8mb4&collation=utf8_general_ci&timeout=60s&loc=Local&autocommit=true")
//os.Setenv("MysqlCollation", "utf8_general_ci")
//os.Setenv("MysqlMaxOpenConn", "0")
//os.Setenv("MysqlMaxIdleConn", "0")
//os.Setenv("MysqlConnMaxLifetime", "0s")

// TablePrefix : database table prefix
var TablePrefix = os.Getenv("MysqlTablePrefix")

// NewDBConn : orm conn
func NewDBConn() *gorm.DB {
	dbConn, err := newDBConn()
	if err != nil {
		panic(fmt.Errorf("dblink.newDBConn() error : %v", err))
	}
	return setDBConnOptions(dbConn)
}

// newDBConn : orm conn
func newDBConn() (*gorm.DB, error) {
	// db connection
	dbConn, err := gorm.Open("mysql", generateDSN())
	if err != nil {
		err = fmt.Errorf("gorm.Open db connection fail : %v", err)
		return nil, err
	}

	// ping
	if err := dbConn.DB().Ping(); err != nil {
		err = fmt.Errorf("ping database connection fail : %v", err)
		return nil, err
	}
	return dbConn, err
}

// set dbConn options
func setDBConnOptions(dbConn *gorm.DB) *gorm.DB {
	// debug
	dbConn = setDebug(dbConn)

	// set table prefix
	setTablePrefix()

	// set max open conn
	if err := setMaxOpenConn(dbConn); err != nil {
		fmt.Println(err)
	}

	// set max idle conn
	if err := setMaxIdleConn(dbConn); err != nil {
		fmt.Println(err)
	}

	// set conn max lifetime
	if err := setConnMaxLifetime(dbConn); err != nil {
		fmt.Println(err)
	}
	return dbConn
}

// setTablePrefix : set table prefix
func setTablePrefix() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, tableName string) string {
		if TablePrefix != "" && !strings.HasPrefix(tableName, TablePrefix) {
			return TablePrefix + tableName
		}
		return tableName
	}
}

// setDebug : print debug
// LogMode set log mode, `true` for detailed logs, `false` for no log,
// default, will only print error logs
func setDebug(dbConn *gorm.DB) *gorm.DB {
	debug := os.Getenv("MysqlDebug")
	if debug != "" && strings.ToLower(debug) == "true" {
		dbConn.LogMode(true)
	} else {
		dbConn.LogMode(false)
	}
	return dbConn
}

// setMaxOpenConn : set max open conn
// SetMaxOpenConns sets the maximum number of open connections to the database.
//
// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
// MaxIdleConns, then MaxIdleConns will be reduced to match the new
// MaxOpenConns limit.
//
// If n <= 0, then there is no limit on the number of open connections.
// The default is 0 (unlimited).
func setMaxOpenConn(dbConn *gorm.DB) error {
	openNumberString := os.Getenv("MysqlMaxOpenConn")
	if openNumberString == "" {
		return nil
	}
	openNumber, err := strconv.Atoi(openNumberString)
	if err != nil {
		return fmt.Errorf("setMaxOpenConn strconv.Atoi error : %v", err)
	}

	// set
	dbConn.DB().SetMaxOpenConns(openNumber)

	return nil
}

// setMaxIdleConn : set max idle conn
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
func setMaxIdleConn(dbConn *gorm.DB) error {
	idleNumberString := os.Getenv("MysqlMaxIdleConn")
	if idleNumberString == "" {
		return nil
	}
	idleNumber, err := strconv.Atoi(idleNumberString)
	if err != nil {
		return fmt.Errorf("setMaxIdleConn strconv.Atoi error : %v", err)
	}

	// set
	dbConn.DB().SetMaxIdleConns(idleNumber)

	return nil
}

// setConnMaxLifetime : conn max lifetime
// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
//
// Expired connections may be closed lazily before reuse.
//
// If d <= 0, connections are reused forever.
func setConnMaxLifetime(dbConn *gorm.DB) error {
	lifetime := os.Getenv("MysqlConnMaxLifetime")
	if lifetime == "" {
		return nil
	}

	// duration
	d, err := time.ParseDuration(lifetime)
	if err != nil {
		return fmt.Errorf("setConnMaxLifetime time.ParseDuration error : %v", err)
	}

	// set
	dbConn.DB().SetConnMaxLifetime(d)

	return nil
}

// generateDSN : dsn = "root:Mysql.123456@tcp(127.0.0;1:3306)/test?charset=utf8&loc=Local"
// github.com/go-sql-driver/mysql
// mysql.Config{}.FormatDSN()
func generateDSN() string {
	var dsn string
	dsn += os.Getenv("MysqlUser") + ":" + os.Getenv("MysqlPassword")
	dsn += "@" + os.Getenv("MysqlNet")
	dsn += "(" + os.Getenv("MysqlAddr") + ")"
	dsn += "/" + os.Getenv("MysqlName") + "?"
	dsn += os.Getenv("MysqlParameters")

	return dsn
}
