package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Options struct {
	DbType   string
	DbFile   string
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Cfg      *gorm.Config
}

// builder模式
func NewDbOptions(user, password, host, port, database string) *Options {
	return &Options{
		DbType:   "mysql",
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
		Cfg:      &gorm.Config{},
	}
}

func (o *Options) Type(dbtype string) *Options {
	o.DbType = dbtype
	return o
}

func (o *Options) File(file string) *Options {
	o.DbFile = file
	return o
}

func (o *Options) Config(cfg *gorm.Config) *Options {
	o.Cfg = cfg
	return o
}

func (o *Options) NewDB() (*gorm.DB, error) {
	switch o.DbType {
	case "sqlite":
		return NewSqliteDb(o.DbFile, o.Cfg)
	case "mysql":
		return NewMysqlDb(o.User, o.Password, o.Host, o.Port, o.Database, o.Cfg)
	case "postgresl":
		return NewPostgresqlDb(o.User, o.Password, o.Host, o.Port, o.Database, o.Cfg)
	default:
		return nil, fmt.Errorf("wrong database type")
	}
}

func NewSqliteDb(file string, cfg *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(file), cfg)
	return db, err
}

func NewMysqlDb(user, pass, host, port, database string, cfg *gorm.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, database)
	fmt.Println(">>>>>>>>" + dsn + "<<<<<<<<<")
	db, err := gorm.Open(mysql.Open(dsn), cfg)
	return db, err
}

func NewPostgresqlDb(user, pass, host, port, database string, cfg *gorm.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		user,
		pass,
		database,
		port)
	db, err := gorm.Open(postgres.Open(dsn), cfg)
	return db, err
}
