package repository

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Comment struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"-"` //用于一对多关系的id
	VideoId    int64     `json:"-"` //一对多，视频对评论
	User       User      `json:"user" gorm:"-"`
	Content    string    `json:"content"`
	CreateDate string    `json:"create_date" gorm:"-"`
	CreatedAt  time.Time `json:"-"`
}

type CommentDAO struct {
}

var (
	commentDao CommentDAO
)

func InitCommentDAO() *CommentDAO {
	return &commentDao
}

func (c *CommentDAO) AddComment(comment *Comment) error {
	if comment == nil {
		return errors.New("AddCommentAndUpdateCount comment空指针")
	}
	// 执行事务
	return DB.Transaction(func(tx *gorm.DB) error {
		// 添加评论数据
		if err := tx.Create(comment).Error; err != nil {
			return err
		}
		// 增加count
		if err := tx.Exec("UPDATE videos v SET v.comment_count = v.comment_count+1 WHERE v.id=?", comment.VideoId).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

func (c *CommentDAO) DeleteComment(commentId, videoId int64) error {
	//执行事务
	return DB.Transaction(func(tx *gorm.DB) error {
		//删除评论
		if err := tx.Exec("DELETE FROM comments WHERE id = ?", commentId).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		//减少count
		if err := tx.Exec("UPDATE videos v SET v.comment_count = v.comment_count-1 WHERE v.id=? AND v.comment_count>0", videoId).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

func (c *CommentDAO) QueryCommentById(id int64, comment *Comment) error {
	if comment == nil {
		return errors.New("QueryCommentById comment 空指针")
	}
	return DB.Where("id=?", id).First(comment).Error
}

func (c *CommentDAO) QueryCommentListByVideoId(videoId int64, commentList *[]*Comment) error {
	if commentList == nil {
		return errors.New("QueryCommentListByVideoId commentList 空指针")
	}

	if err := DB.Where("video_id = ?", videoId).Find(commentList).Error; err != nil {
		return err
	}
	return nil
}
