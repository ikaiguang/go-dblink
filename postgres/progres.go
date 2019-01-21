package postgres

import (
	"fmt"
	"github.com/ikaiguang/go-dblink/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewDBConn : orm conn
func NewDBConn(authCfg *configs.AuthConfig) (*gorm.DB, error) {
	// db connection
	dbConn, err := gorm.Open("postgres", DefaultGetDSNHandler(authCfg))
	if err != nil {
		err = fmt.Errorf("gorm.Open db connection fail : %v", err)
		return nil, err
	}
	return dbConn, err
}

// DefaultGetDSNHandler : dsn = "host=myhost port=myport user=gorm dbname=gorm password=mypassword"
// DefaultGetDSNHandler : dsn = "postgresql://[user[:password]@][netloc][:port][,...][/dbname][?param1=value1&...]"
// https://www.postgresql.org/docs/11/libpq-connect.html#LIBPQ-CONNSTRING
var DefaultGetDSNHandler = func(cfg *configs.AuthConfig) string {
	var dsn string

	//dsn += "host=" + cfg.Host + " "
	//dsn += "port=" + cfg.Port + " "
	//dsn += "user=" + cfg.Username + " "
	//dsn += "dbname=" + cfg.DBName + " "
	//dsn += "password=" + cfg.Password + " "
	//dsn += cfg.Parameters

	dsn += "postgresql://" + cfg.Username + ":" + cfg.Password
	dsn += "@" + cfg.Host + ":" + cfg.Port
	dsn += "/" + cfg.DBName + "?" + cfg.Parameters

	return dsn
}
