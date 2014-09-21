package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/macococo/go-gamereviews/conf"
	"github.com/macococo/go-gamereviews/utils"
	"log"
	"os"
)

var (
	DbMap *gorp.DbMap
)

func init() {
	dataSource := conf.Config.Datasource
	maxIdleConns := conf.Config.MaxIdleConns
	maxOpenConns := conf.Config.MaxOpenConns

	log.Println("mysql datasource:", dataSource, ", maxIdleConns", maxIdleConns, ", maxOpenConns", maxOpenConns)

	db, err := sql.Open("mysql", dataSource)
	utils.HandleError(err)

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)

	DbMap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// Trace query log
	if conf.IsDev() {
		DbMap.TraceOn("", log.New(os.Stdout, "", log.Lmicroseconds))
	}

	utils.HandleError(DbMap.CreateTablesIfNotExists())
}
