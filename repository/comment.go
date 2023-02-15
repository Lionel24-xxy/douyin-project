package repository

import "time"

type Comment struct {
	Id         		int64     	`json:"id"`
	UserId 			int64     	`json:"-"` //用于一对多关系的id
	VideoId    		int64     	`json:"-"` //一对多，视频对评论
	User       		User      	`json:"user" gorm:"-"`
	Content    		string    	`json:"content"`
	CreateDate 		string    	`json:"create_date" gorm:"-"`
	CreatedAt  		time.Time 	`json:"-"`
}