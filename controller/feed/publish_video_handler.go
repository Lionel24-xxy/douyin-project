package feed

import (
	"TikTok_Project/service/video"
	"TikTok_Project/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"path/filepath"
)

var (
	videoExtension = map[string]struct{}{
		".mp4":  {},
		".avi":  {},
		".wmv":  {},
		".flv":  {},
		".mpeg": {},
		".mov":  {},
	}
)

func PublishVideoError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 1,
		"status_msg":  msg,
	})
}

func PublishVideoOk(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  msg,
	})
}

// PublishVideoHandler 发布视频，并截取视频的第一帧画面作为视频封面
func PublishVideoHandler(ctx *gin.Context) {
	//获取user_id
	rawID, _ := ctx.Get("user_id")
	userID, ok := rawID.(int64) //类型判断
	if !ok {
		PublishVideoError(ctx, "解析UserID错误！")
		return
	}
	title := ctx.PostForm("title")
	form, err := ctx.MultipartForm() //文件上传的解析
	if err != nil {
		PublishVideoError(ctx, err.Error())
		return
	}

	//支持多文件上传(?)
	files := form.File["data"]
	for _, file := range files {
		suffix := filepath.Ext(file.Filename)    //filepath.Ext函数用来获取文件后缀
		if _, ok = videoExtension[suffix]; !ok { //判断是否为我们支持的视频格式
			PublishVideoError(ctx, "不支持的视频格式类型")
			continue
		}
		VideoName := utils.NewFileName(userID) //根据UserID得到唯一的video文件名
		filename := VideoName + suffix
		//Join 将任意数量的路径元素连接到一个路径中，用操作系统特定的分隔符将它们分开。
		savePath := filepath.Join("D:\\Golang_workspace\\src\\douyin-project\\static", filename) //要更换成自己的static目录
		err = ctx.SaveUploadedFile(file, savePath)
		if err != nil {
			PublishVideoError(ctx, err.Error())
			continue
		}
		//截取第一帧画面作为封面
		pictureName := VideoName + ".jpg"
		savePath2 := filepath.Join("D:\\Golang_workspace\\src\\douyin-project\\static", pictureName) //要更换成自己的static目录
		cmd := exec.Command("ffmpeg", "-i", savePath, "-frames:v", "1", savePath2)
		err := cmd.Run()
		if err != nil {
			PublishVideoError(ctx, err.Error())
			continue
		}
		//数据库持久化
		Err := video.PostVideo(userID, filename, pictureName, title)
		if Err != nil {
			PublishVideoError(ctx, Err.Error())
			continue
		}
		fmt.Println(Err.Error())
		PublishVideoOk(ctx, filename+"上传成功!")
	}
}
