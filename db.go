package CurryPoint

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"log"
	"github.com/newtechfellas/CurryPoint/entity"
)

var DBMAP *gorp.DbMap = nil

func PrepareDBMap(connectionStr string) *gorp.DbMap {
	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		log.Fatalln("Failed to connect to database " + connectionStr + ". Error is " + err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatalln("Could not ping the database " + connectionStr + ". Error is " + err.Error())
	}

	// construct a gorp DbMap
	DBMAP = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// add a table, setting the table name to 'posts' and
	(&entity.Customer{}).ConfigureDBMap(DBMAP)

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	if err = DBMAP.CreateTablesIfNotExists(); err != nil {
		log.Fatalln("Could not create tables. Exiting")
	}

	return DBMAP
}