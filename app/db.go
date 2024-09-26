package app

import (
	"database/sql"
	"time"

	"github.com/Rakhulsr/go-rest-api-pzn/helper"
	"github.com/Rakhulsr/go-rest-api-pzn/utils"
	_ "github.com/go-sql-driver/mysql"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", utils.Envs.Dburl)
	helper.PanicError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
