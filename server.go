package main

import (
	"github.com/go-martini/martini"
	"github.com/gvnn/gotest/app"
	"github.com/gvnn/gotest/app/models"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})

	// initialize the DbMap
	dbmap := app.InitDB()
	defer dbmap.Db.Close()

	// delete any existing rows
	err := dbmap.TruncateTables()
	app.CheckErr(err, "TruncateTables failed")

	// create one user
	u1 := models.User{
		Email: "info@gvnn.it",
	}

	// insert rows
	err = dbmap.Insert(&u1)
	app.CheckErr(err, "Insert failed")

	m.Run()
}
