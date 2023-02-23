package utils

import (
	"TikTok_Project/repository"
	"errors"
)

func FillCommentFields(comment *repository.Comment) error {
	if comment == nil {
		return errors.New("FillCommentFields comments为空")
	}
	comment.CreateDate = comment.CreatedAt.Format("1-2") //转为前端要求的日期格式
	return nil
}

func FillCommentListFields(commentList *[]*repository.Comment) error {
	if commentList == nil {
		return errors.New("FillCommentListFields commentList 空指针")
	}

	for _, l := range *commentList {
		l.CreateDate = l.CreatedAt.Format("1-2")
	}
	return nil
}