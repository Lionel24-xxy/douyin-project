package video

import "TikTok_Project/repository"

func PostVideo(userID int64, videoName, coverName, title string) error {
	return NewPostVideoFlow(userID, videoName, coverName, title).Do()
}

func NewPostVideoFlow(userId int64, videoName, coverName, title string) *PostVideoFlow {
	return &PostVideoFlow{
		videoName: videoName,
		coverName: coverName,
		userId:    userId,
		title:     title,
	}
}

type PostVideoFlow struct {
	videoName string
	coverName string
	title     string
	userId    int64

	video *repository.Video
}

func (f *PostVideoFlow) Do() error {
	if err := f.publish(); err != nil {
		return err
	}
	return nil
}

// 组合并添加到数据库
func (f *PostVideoFlow) publish() error {
	video := &repository.Video{
		UserId:   f.userId,
		PlayUrl:  f.videoName,
		CoverUrl: f.coverName,
		Title:    f.title,
	}
	return repository.NewVideoDAO().AddVideo(video)
}
