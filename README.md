# 极简抖音项目

## 项目结构

```
│  main.go 为主文件，运行时`go run`即可
│
├─controller 控制层，用于接收参数，编写逻辑，返回参数
│  ├─comment
│  │      comment_action_handler.go 发布评论及删除评论
│  │      comment_list_handler.go 评论列表
│  │
│  ├─feed
│  │      feedList_handler.go 视频流处理器
│  │      ffmpeg.exe
│  │      publish_video_handler.go 投稿视频处理器
│  │      query_publish_list_handler.go  发布列表处理器
│  │
│  └─user 
│          user_info_handler.go  用户信息处理器
│          user_login_handler.go 用户登录处理器
│          user_register_handler.go 用户注册处理器
│
├─middleware 中间件，包括用户鉴权功能（验证`Token`同时将`Token`中携带的`user_id`信息写入上下文）
│      AuthMiddleWare.go  用户鉴权
│
├─repository 模型层，用于数据库初始化以及完成和数据库有关的操作
│      comment.go 评论模型
│      mysql_init.go Mysql初始化
│      redis.go  Redis初始化
│      redis_test.go 
│      user.go 用户模型
│      video.go 视频模型
│
├─router 路由层，包含路由组，项目不同功能的入口
│      router.go 路由存放
│
├─service 服务层，用于实现主要业务函数
│  ├─comment
│  │      get_comment_list.go 获取视频评论列表
│  │      publish_comment.go 发布评论及删除评论
│  │
│  ├─user 
│  │      common.go  密码强度检测
│  │      common_test.go 
│  │      user_login.go 用户登录
│  │      user_register.go  用户注册
│  │
│  └─video
│          feed_video.go 视频流
│          post_video.go 投稿视频
│          query_videolist.go 发布列表
│
├─static 放视频和封面的文件夹
│      
└─utils 工具包，包括`jwt`令牌功能、密码加密功能等
        comment.go 修改评论时间为前端要求格式
        jwt.go   jwt令牌
        password.go  sha1加密
        snowflakes.go 雪花算法
        video.go  雪花id和userid生成独一无二视频名，视频url拼接
```


首先你要在自己的文件夹创建static目录，其次修改mysql和redis的密码（需要自己创建tiktok数据库），更改视频url的address，然后在main.go同样修改address，然后运行main.go。
