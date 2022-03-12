package db

import (
	"api/support"
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

type PostgresDb struct {
	db *sql.DB
}

func connect() *gorm.DB {
	config := support.LoadConfig()
	db, err := gorm.Open(postgres.Open(config.DB))
	support.PanicError(err)

	return db
}

func Find(dest interface{}, conds ...interface{}) {
	db := connect()
	db.Find(dest, conds...)
}

func First(dest interface{}, conds ...interface{}) {
	db := connect()
	db.First(dest, conds...)
}

func Create(value interface{}) {
	db := connect()
	db.Create(value)
}

func Delete(value interface{}, id int) {
	db := connect()
	db.Delete(value, id)
}

func Save(value interface{}) {
	db := connect()
	db.Save(value)
}
