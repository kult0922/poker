package services

import (
	"github.com/kult0922/go-react-blog/backend/apperrors"
	"github.com/kult0922/go-react-blog/backend/models"
	"github.com/kult0922/go-react-blog/backend/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {

	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
