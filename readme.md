# go-dblink

package ormlink

## install

`github.com/ikaiguang/go-dblink`

## use

> mysql

```go

import (
	orm "github.com/ikaiguang/go-dblink/mysql"
)

func GetDBConn() {
	dbConn := orm.NewDBConn()

	_ = dbConn
}

```

## test

```go

package mysql

import (
	"os"
	"testing"
)

func TestNewDBConn(t *testing.T) {
	os.Setenv("MysqlDebug", "true")
	os.Setenv("MysqlTablePrefix", "ikg_")
	os.Setenv("MysqlUser", "root")
	os.Setenv("MysqlPassword", "Mysql.123456")
	os.Setenv("MysqlNet", "tcp")
	os.Setenv("MysqlAddr", "127.0.0.1:3306")
	os.Setenv("MysqlName", "test")
	os.Setenv("MysqlParameters", "charset=utf8mb4&collation=utf8_general_ci&timeout=60s&loc=Local&autocommit=true")
	os.Setenv("MysqlCollation", "utf8_general_ci")
	os.Setenv("MysqlMaxOpenConn", "0")
	os.Setenv("MysqlMaxIdleConn", "0")
	os.Setenv("MysqlConnMaxLifetime", "0s")

	if err := NewDBConn().DB().Ping(); err != nil {
		t.Errorf("testing : NewDBConn().DB().Ping() error : %v", err)
	}
}


```