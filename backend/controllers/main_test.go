package controllers_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kult0922/go-react-blog/backend/controllers"
	"github.com/kult0922/go-react-blog/backend/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)
	m.Run()
}
