package tests

import (
	"io"
	"net/http"

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

type ErrorJSON struct {
	Errors map[string][]string `json:"errors"`
}

var (
	DB       *gorm.DB
	fixtures *testfixtures.Context
	JWT      auth.Tokener
	users    []*models.User
	articles []*models.Article
	tags     []*models.Tag
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

	DB.Model(models.Article{}).
		Preload("User").
		Preload("Tags").
		Preload("Comments").
		Preload("Comments.User").
		Preload("Favorites").
		Preload("Favorites.User").
		Find(&articles)

	DB.Find(&tags)
}

func (t *AppTest) After() {
	println("Tear down")
}

func (t *AppTest) TestConnection() {
	t.Assert(DB != nil)
	t.Assert(fixtures != nil)
	t.AssertEqual(8, len(users))
	t.AssertEqual(5, len(articles))
	t.AssertEqual(3, len(tags))
}

func (t *AppTest) MakePostRequest(url string, body io.Reader, header http.Header) {
	request := t.PostCustom(t.BaseUrl()+url, "application/json", body)
	if header != nil {
		request.Header = header
	}

	request.Send()
}
