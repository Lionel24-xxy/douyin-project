# 极简抖音项目

## 项目结构

```bash
│ main.go 为主文件，运行时`go run`即可
├─controller
│  ├─comment
│  │      comment_action_handler.go
│  │      comment_list_handler.go
│  │
│  ├─feed
│  │      feed_list_handler.go
│  │      ffmpeg.exe
│  │      publish_video_handler.go
│  │      query_follower_handler.go
│  │      query_publish_list_handler.go
│  │
│  └─user
│          user_info_handler.go
│          user_login_handler.go
│          user_register_handler.go
│
├─middleware
│      AuthMiddleWare.go
│
├─repository
│      comment.go
│      mysql_init.go
│      redis.go
│      redis_test.go
│      user.go
│      video.go
│
├─router
│      router.go
│
├─service
│  ├─comment
│  │      get_comment_list.go
│  │      publish_comment.go
│  │
│  ├─user
│  │      common.go
│  │      common_test.go
│  │      query_follower_list.go
│  │      user_login.go
│  │      user_register.go
│  │
│  └─video
│          feed_video.go
│          post_video.go
│          query_videolist.go
│
├─static
│
└─utils
        comment.go
        jwt.go
        password.go
        snowflakes.go
        video.go
```

