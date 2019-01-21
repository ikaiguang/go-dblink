# go-dblink

package dblink

## install

`github.com/ikaiguang/go-dblink`

## demo

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

## setting

```go

os.Setenv("DatabaseOrmDebug", "true")
os.Setenv("DatabaseTablePrefix", "ikg_")
os.Setenv("DatabaseMaxOpenConn", "10")
os.Setenv("DatabaseMaxIdleConn", "10")
os.Setenv("DatabaseConnMaxLifetime", "30s")

```