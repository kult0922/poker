package repositories_test

import (
	"testing"

	"github.com/kult0922/go-react-blog/backend/models"
	"github.com/kult0922/go-react-blog/backend/repositories"
	"github.com/kult0922/go-react-blog/backend/repositories/testdata"
)

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "testest",
	}

	expectedCommentNum := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}
	if newComment.CommentID != expectedCommentNum {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentNum, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where article_id = ? and message = ?
		`
		testDB.Exec(sqlStr, comment.ArticleID, comment.Message)
	})
}

func TestSelectCommentList(t *testing.T) {
	expectedNum := len(testdata.CommentTestData)
	got, err := repositories.SelectCommentList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d comments\n", expectedNum, num)
	}
}
