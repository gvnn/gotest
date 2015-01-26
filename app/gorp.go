package app

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	"github.com/gvnn/gotest/app/models"
	_ "github.com/ziutek/mymysql/godrv"
	"log"
)

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func InitDB() *gorp.DbMap {
	db, err := sql.Open("mymysql", "tcp:localhost:3306*gotest/gotest/gotest")
	CheckErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// add a table, setting the table name to 'users' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(models.User{}, "users").SetKeys(true, "Id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	CheckErr(err, "Create tables failed")

	return dbmap
}
