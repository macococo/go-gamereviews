package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/macococo/go-gamereviews/utils"
	"log"
	"os"
)

var (
	DbMap *gorp.DbMap
)

func init() {
	log.Println("database init.")

	db, err := sql.Open("mysql", "gamereviews:gamereviews@tcp(127.0.0.1:3306)/gamereviews?charset=utf8")
	utils.HandleError(err)

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	DbMap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// Debug
	DbMap.TraceOn("", log.New(os.Stdout, "", log.Lmicroseconds))

	utils.HandleError(DbMap.CreateTablesIfNotExists())
}
