package mysql

import (
	"fmt"
	"github.com/ikaiguang/go-dblink/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // driver
)

// NewDBConn : orm conn
func NewDBConn() (*gorm.DB, error) {
	// db connection
	dbConn, err := gorm.Open("mysql", DefaultGetDSNHandler())
	if err != nil {
		err = fmt.Errorf("gorm.Open db connection fail : %v", err)
		return nil, err
	}
	return dbConn, err
}

// DefaultGetDSNHandler : dsn = "root:Mysql.123456@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Local"
// github.com/go-sql-driver/mysql
// mysql.Config{}.FormatDSN()
var DefaultGetDSNHandler = func() string {
	// config
	cfg := configs.GetAuthConfigHandler()

	var dsn string

	dsn += cfg.Username + ":" + cfg.Password
	dsn += "@tcp(" + cfg.Host + ":" + cfg.Port + ")"
	dsn += "/" + cfg.DBName + "?" + cfg.Parameters

	return dsn
}
