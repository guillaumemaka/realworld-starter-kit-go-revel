package controllers

import (
	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/models"
	"github.com/revel/revel"
)

type GormController struct {
	*revel.Controller
	DB models.Datastorer
}

var (
	DB *models.DB
)

func InitDB() {
	dialect := revel.Config.StringDefault("db.dialect", "sqlite3")
	dbname := revel.Config.StringDefault("db.dbname", "conduit.db")

	var err error
	DB, err = models.NewDB(dialect, dbname)

	if err != nil {
		revel.INFO.Println("Database initialization error", err)
	}

	DB.InitSchema()
}

func (c *GormController) Init() revel.Result {
	c.DB = DB
	return nil
}
