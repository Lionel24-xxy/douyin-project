package video

import (
	"TikTok_Project/repository"
	"TikTok_Project/utils"
)

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
	f.prepareParam()
	f.title = f.sensitiveCheck(f.title)
	if err := f.publish(); err != nil {
		return err
	}
	if err := f.addWorkCount(); err != nil {
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

// 增加发布作品数
func (f *PostVideoFlow) addWorkCount() error {
	return repository.NewVideoDAO().UpdateWorkCount(f.userId)
}

// 获得url
func (f *PostVideoFlow) prepareParam() {
	f.videoName = utils.GetFileUrl(f.videoName)
	f.coverName = utils.GetFileUrl(f.coverName)
}

// 敏感词检测及替换
func (f *PostVideoFlow) sensitiveCheck(text string) string {

	isContain := utils.SensitiveWordCheck(text, int(f.userId))
	if isContain {
		replaceText := utils.SensitiveWordReplace(text)
		return replaceText
	}
	return text
}
