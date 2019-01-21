package mssql

import "github.com/ikaiguang/go-dblink/config"

//import (
//	"fmt"
//	"github.com/ikaiguang/go-dblink/config"
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/mssql"
//)
//
//// NewDBConn : orm conn
//func NewDBConn() (*gorm.DB, error) {
//	// db connection
//	dbConn, err := gorm.Open("mssql", DefaultGetDSNHandler())
//	if err != nil {
//		err = fmt.Errorf("gorm.Open db connection fail : %v", err)
//		return nil, err
//	}
//	return dbConn, err
//}

// DefaultGetDSNHandler : dsn = "sqlserver://username:password@localhost:1433?database=dbname"
var DefaultGetDSNHandler = func() string {
	// config
	cfg := configs.GetAuthConfigHandler()

	var dsn string

	dsn += "sqlserver://"
	dsn += cfg.Username + ":" + cfg.Password
	dsn += "@" + cfg.Host + ":" + cfg.Port + "?"
	dsn += "database=" + cfg.DBName

	return dsn
}
