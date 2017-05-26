package tests

import (
	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/lib/auth"
	"github.com/guillaumemaka/realworld-starter-kit-go-revel/app/models"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"github.com/revel/revel/testing"
	"gopkg.in/testfixtures.v2"
)

type AppTest struct {
	testing.TestSuite
}

var (
	DB       *gorm.DB
	fixtures *testfixtures.Context
	JWT      auth.Tokener
	users    []models.User
)

func (t *AppTest) Before() {
	dialect, _ := revel.Config.String("db.dialect")
	dbname, _ := revel.Config.String("db.dbname")

	db, err := models.NewDB(dialect, dbname)

	if err != nil {
		revel.ERROR.Fatal(err)
	}

	DB = db.DB
	JWT = auth.NewJWT()

	fixtures, err = testfixtures.NewFolder(DB.DB(), &testfixtures.SQLite{}, "tests/testdata/fixtures")

	if err != nil {
		revel.ERROR.Fatal(err)
	}

	err = fixtures.Load()

	if err != nil {
		revel.ERROR.Fatal(err)
	}

	DB.Find(&users)
}

func (t *AppTest) After() {
	println("Tear down")
}

func (t *AppTest) TestConnection() {
	t.Assert(DB != nil)
	t.Assert(fixtures != nil)
	t.AssertEqual(8, len(users))
}
