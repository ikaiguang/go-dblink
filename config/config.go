package configs

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// AuthConfig auth cfg
type AuthConfig struct {
	Driver     string // os.Setenv("DatabaseDriver", "mysql")
	Username   string // os.Setenv("DatabaseUsername", "username")
	Password   string // os.Setenv("DatabasePassword", "password")
	Host       string // os.Setenv("DatabaseHost", "127.0.0.1")
	Port       string // os.Setenv("DatabasePort", "3306")
	DBName     string // os.Setenv("DatabaseDBName", "test")
	Parameters string // os.Setenv("DatabaseParameters", "charset=utf8&timeout=60s&loc=Local&autocommit=true")
}

// GetAuthConfigHandler get database config
var GetAuthConfigHandler = func() *AuthConfig {
	var cfg AuthConfig

	cfg.Driver = strings.TrimSpace(os.Getenv("DatabaseDriver"))
	cfg.Username = strings.TrimSpace(os.Getenv("DatabaseUsername"))
	cfg.Password = strings.TrimSpace(os.Getenv("DatabasePassword"))
	cfg.Host = strings.TrimSpace(os.Getenv("DatabaseHost"))
	cfg.Port = strings.TrimSpace(os.Getenv("DatabasePort"))
	cfg.DBName = strings.TrimSpace(os.Getenv("DatabaseDBName"))
	cfg.Parameters = strings.TrimSpace(os.Getenv("DatabaseParameters"))

	return &cfg
}

// tablePrefix : database table prefix
var tablePrefix = strings.TrimSpace(os.Getenv("DatabaseTablePrefix"))

// GetTablePrefixHandler get table prefix
var GetTablePrefixHandler = func() string {
	return tablePrefix
}

// OptionConfig option cfg
type OptionConfig struct {
	// sets orm LogMode
	Debug bool // os.Setenv("DatabaseOrmDebug", "true")

	// sets table prefix
	Prefix string // os.Setenv("DatabaseTablePrefix", "ikg_")

	// sets the maximum number of open connections to the database.
	MaxOpen int // os.Setenv("DatabaseMaxOpenConn", "10")

	// sets the maximum number of connections in the idle
	MaxIdle int // os.Setenv("DatabaseMaxIdleConn", "10")

	// sets the maximum amount of time a connection may be reused.
	MaxLifetime time.Duration // os.Setenv("DatabaseConnMaxLifetime", "30s")
}

// GetOptionConfigHandler get database config
var GetOptionConfigHandler = func() *OptionConfig {
	var cfg OptionConfig
	var err error

	// debug
	debugString := strings.TrimSpace(os.Getenv("DatabaseOrmDebug"))
	//cfg.Debug = strings.ToLower(debugString) == "true"
	cfg.Debug, _ = strconv.ParseBool(debugString)

	// prefix
	cfg.Prefix = GetTablePrefixHandler()

	// max open
	openString := strings.TrimSpace(os.Getenv("DatabaseMaxOpenConn"))
	cfg.MaxOpen, err = strconv.Atoi(openString)
	if err != nil {
		log.Println("strconv.Atoi(DatabaseMaxOpenConn) error : ", err)
	}

	// max idle
	idleString := strings.TrimSpace(os.Getenv("DatabaseMaxIdleConn"))
	cfg.MaxIdle, err = strconv.Atoi(idleString)
	if err != nil {
		log.Println("strconv.Atoi(DatabaseMaxIdleConn) error : ", err)
	}

	// max lifetime
	lifetimeString := strings.TrimSpace(os.Getenv("DatabaseConnMaxLifetime"))
	cfg.MaxLifetime, err = time.ParseDuration(lifetimeString)
	if err != nil {
		log.Println("time.ParseDuration(DatabaseConnMaxLifetime) error : ", err)
	}
	return &cfg
}
