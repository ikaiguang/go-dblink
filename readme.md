# go-dblink

package dblink

## install

```bash

go get github.com/ikaiguang/go-dblink

```

## demo

> mysql

```go

package dblink

import (
	"github.com/jinzhu/gorm"
	"os"
	"testing"
)

func TestNewDBConn(t *testing.T) {
	// option
	os.Setenv("DatabaseOrmDebug", "true")
	os.Setenv("DatabaseTablePrefix", "ikg_")
	os.Setenv("DatabaseMaxOpenConn", "10")
	os.Setenv("DatabaseMaxIdleConn", "10")
	os.Setenv("DatabaseConnMaxLifetime", "30s")

	var db *gorm.DB

	// mysql
	t.Logf("test mysql ... \n")
	// auth
	// root:Mysql.123456@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Local
	os.Setenv("DatabaseDriver", "mysql")
	os.Setenv("DatabaseUsername", "root")
	os.Setenv("DatabasePassword", "Mysql.123456")
	os.Setenv("DatabaseHost", "127.0.0.1")
	os.Setenv("DatabasePort", "3306")
	os.Setenv("DatabaseDBName", "test")
	os.Setenv("DatabaseParameters", "charset=utf8&timeout=60s&loc=Local&autocommit=true")
	db = NewDBConn()
	t.Logf("%v \n", db)

	// postgres
	t.Logf("test postgres ... \n")
	// auth
	// host=myhost port=myport user=gorm dbname=gorm password=mypassword
	// postgresql://[user[:password]@][netloc][:port][,...][/dbname][?param1=value1&...]
	os.Setenv("DatabaseDriver", "postgres")
	os.Setenv("DatabaseUsername", "postgres")
	os.Setenv("DatabasePassword", "Postgres.123456")
	os.Setenv("DatabaseHost", "127.0.0.1")
	os.Setenv("DatabasePort", "5432")
	os.Setenv("DatabaseDBName", "postgres")
	//os.Setenv("DatabaseParameters", "sslmode=disable connect_timeout=20")
	os.Setenv("DatabaseParameters", "connect_timeout=20&sslmode=disable")
	db = NewDBConn()
	t.Logf("%v \n", db)

	//db = NewDBConn(mssql.NewDBConn)
	//t.Logf("%v", db)
}

```