# 一、项目介绍

本项目实现了抖声服务端的基础功能（视频流，投稿视频，用户注册登录等），同时拓展了互动模块和社交模块（点赞，评论，关注等）。

## 1.1 项目仓库和Apk文件

Github项目地址：https://github.com/Lionel24-xxy/douyin-project

抖声app的apk地址：[极简抖音App使用说明 - 青训营版](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7) 

## 1.2 项目环境配置说明

### 1.2.1 环境配置

1. Go版本==1.19.4
2. 数据库：MySQL 8.0.23
3. Redis：3.0.504

### 1.2.2 项目使用

1. 手机安装抖声的客户端
2. 将手机和电脑服务端共用一个`wifi`（也可以采用手机热点，电脑连接的方式）
3. 使用`win+R`输入`cmd`，在终端输入`ipconfig`，找到对应的ipv4的地址
4. 打开项目，在`main.go`文件中修改`r.Run()`，输入ipv4和默认端口8080
5. 再去`utils`的`video.go`文件中修改全局变量：`defaultIP和defaultPort`
6. MySQL数据库只需要在本地创建数据库`tiktok`即可，运行之后自动会生成对应的数据表，注意需要在`mysql_init.go`中修改为你的`username`和`password`
7. 启动`Redis`（必须）,同时在`redis.go`中修改redis密码（如果没有redis的密码那就不必修改）
8. 由于你需要存储视频和视频封面，所以你需要创建一个`static目录`
9. 安装依赖。在 `douyin-project `目录下运行`go mod tidy`
10. 运行。打开`Goland`的终端（就是是douyin-project目录下的终端）,输入`go build main.go`命令，之后运行生成的`main.exe`文件

### 1.2.3 项目说明

1. 数据库部署在本地
2. `Redis`是启动项目所必须的，缺省时会缺少限制频率的功能
3. 视频模块中采用本地存储方式
4. 采用`ffmpeg`获取视频封面，`ffmpeg.exe`已同步上传项目，但如果出现`erro`r，建议安装`ffmpeg`然后添加环境变量

# 二、项目分工

| **团队成员** | **主要贡献**                                                 |
| ------------ | ------------------------------------------------------------ |
| 周幸         | 视频流实现，视频投稿功能实现，关注功能和关注列表的实现       |
| 许昕宇       | 架构设计，数据库设计，用户注册登陆及用户信息模块，发布列表，评论及评论列表 |
| 刘涛         | 粉丝列表和好友列表，相关文档的撰写                           |
| 吴思         | 点赞功能和喜欢列表，相关文档的撰写                           |

# 三、项目实现

### 3.1 技术选型与相关开发文档

#### 3.1.1  Web 框架选型

本项目选择 Gin 作为基本的 Web 框架，Gin 是一个标准的 Web 服务框架，它封装比较好,API友好,源码注释比较明确,并且遵循 Restful API 接口规范，其路由库是基于 httproute 实现的，具有快速灵活,容错方便等特点。

#### 3.1.2 数据库选型

本项目选择 Gorm 框架用于操作 MySQL 数据库，Gorm 是 Golang 语言中一款性能极好的 ORM 库。同时选择 go-redis 框架操作 Redis 数据库。

#### 3.1.3 中间件介绍

本项目包括三个中间件，第一个是用于 JWT 鉴权的`JWTMiddleWare()`函数，用于验证 Token 并将 user_id 存入上下文；第二个是用于限制恶意请求的`RateMiddleware()`函数，用于限制同一 IP 在短时间多次访问服务器，避免网站负载升高或者造成网站带宽阻塞而拒绝或无法响应正常用户的请求，第三个是`NoAuthToGetUserid`函数，它是在我们不需要认证的时候，用于将user_id存入上下文的。

#### 3.1.4 安全问题

1. 密码安全

本项目通过正则表达式的方法对密码进行校验，要求用户的密码必须同时包括数字和大小写字母，也可以包含特殊字符，提高了密码的安全等级。

同时使用 SHA-1 算法对用户密码进行加密存储，保证用户的密码安全。

1. 敏感词检测及屏蔽

本项目利用 github.com/feiin/sensitivewords 库评论和发布视频请求中的文本信息实现敏感词的检测，敏感词类型涉及黄赌毒、党政、欺骗消费者等方面。检测到敏感词后使用 * 进行替代，确保绿色健康的网络环境。

#### 3.1.5 雪花算法

雪花算法是推特开源的分布式ID生成算法，用于在不同的机器上生成唯一的ID的算法。该算法生成一个64bit的数字作为分布式ID，保证这个ID自增并且全局唯一。我们使用雪花算法生成uid和user_id拼接成视频名称和封面名称，保证了名称的不可重复性。

#### 3.1.6 代码管理

团队采用 Git 进行代码管理和版本控制，版本迭代清晰，团队分工明确，提升开发效率。

### 3.2 架构设计

#### 3.2.1 数据库表结构设计

![img](https://jdxhj9jomk.feishu.cn/space/api/box/stream/download/asynccode/?code=MzhjNmFhMmYwZjA4NzYwMzM1YjRkMDgxMDczN2E2MzRfQ2FCdVR1aVZzUVlSZHpoYUdaVzliZUFWbUxManIzMVNfVG9rZW46Ym94Y25xdmVvd3B3eTRsdVBqZmNyMG5udElkXzE2NzcxNzY4Nzg6MTY3NzE4MDQ3OF9WNA)

主要有三个表：User、Video、Comment，利用 Gorm 中`Many to Many`另外生成两个表：user_relations、user_favorite 用于存放不同用户之间的关注信息和用户与点赞视频之间的关联信息。

1. User 表

- 主键为ID
- 通过 Gorm 中 `has many` 与 Video 表建立了一对多的连接，表示一个用户有多个投稿视频。
- 通过 Gorm 中 `has many` 与 Comment 表建立了一对多的连接，表示一个用户有多个评论。
- 通过 Gorm 中 `Many to Many` 与 自己建立了多对多的连接，表示用户之间是多对多关系。`Many to Many`会在模型中添加一张连接表 user_relations。
- 通过 Gorm 中 `Many to Many` 与 Video 表建立了多对多的连接，表示用户与点赞视频是多对多关系。`Many to Many`会在模型中添加一张连接表 user_favorite。

```Go
type User struct {
    ID              int64      `json:"id" gorm:"id,omitempty;primaryKey;AUTO_INCREMENT"`
    Username        string     `json:"username" gorm:"username,omitempty"`
    Password        string     `json:"password" gorm:"size:200;notnull"`
    FollowCount     int64      `json:"follow_count" gorm:"follow_count,omitempty"`
    FollowerCount   int64      `json:"follower_count" gorm:"follower_count,omitempty"`
    IsFollow        bool       `json:"is_follow" gorm:"is_follow,omitempty"`
    Avatar          string     `json:"avatar" gorm:"avatar,omitempty"`
    BackgroundImage string     `json:"background_image" gorm:"background_image,omitempty"`
    Signature       string     `json:"signature" gorm:"signature,omitempty"`
    TotalFavorited  int64      `json:"total_favorited" gorm:"total_favorited,omitempty"`
    WorkCount       int64      `json:"work_count" gorm:"work_count,omitempty"`
    FavoriteCount   int64      `json:"favorite_count" gorm:"favorite_count,omitempty"`
    Relations       []*User    `json:"-" gorm:"many2many:user_relations;association_jointable_foreignkey:follow_id"` //用户之间的多对多
    Videos          []*Video   `json:"-"`                                                                            //用户与投稿视频的一对多
    Favorite        []*Video   `json:"-" gorm:"many2many:user_favorite;"`                                            //用户与点赞视频之间的多对多
    Comments        []*Comment `json:"-"`                                                                            //用户与评论的一对多
}
```

1. Video 表

- 通过 Gorm 中 `has many` 与 Comment 表建立了一对多的连接，表示一个视频有多个评论。

```Go
type Video struct {
    Id            int64      `json:"id,omitempty"`
    UserId        int64      `json:"-"`
    Author        User       `json:"author,omitempty" gorm:"-"` 
    PlayUrl       string     `json:"play_url,omitempty"`
    CoverUrl      string     `json:"cover_url,omitempty"`
    FavoriteCount int64      `json:"favorite_count,omitempty"`
    CommentCount  int64      `json:"comment_count,omitempty"`
    IsFavorite    bool       `json:"is_favorite,omitempty"`
    Title         string     `json:"title,omitempty"`
    Users         []*User    `json:"-" gorm:"many2many:user_favorite;"`
    Comments      []*Comment `json:"-"`
    CreatedAt     time.Time  `json:"-"`
    UpdatedAt     time.Time  `json:"-"`
}
```

1. Commet 表

```Go
type Comment struct {
    Id         int64     `json:"id"`
    UserId     int64     `json:"-"` //用于一对多关系的id
    VideoId    int64     `json:"-"` //一对多，视频对评论
    User       User      `json:"user" gorm:"-"`
    Content    string    `json:"content"`
    CreateDate string    `json:"create_date" gorm:"-"`
    CreatedAt  time.Time `json:"-"`
}
```

#### 3.2.2 项目架构设计

本项目采用多级架构设计，router 层用于分发路由，middleware 层包含所有中间件，controller 层用于接收参数、编写不同路由逻辑以及返回参数，service 层用于实现主要的功能代码，repository 层用于完成数据库相关操作，utils 包里包括工具函数（如：密码加密，敏感词检测，雪花算法等）

项目代码目录：

```Go
│  main.go 为主文件，运行时`go run`即可
│
│
├─controller 控制层，用于接收参数，编写逻辑，返回参数
│  ├─comment
│  │      comment_action_handler.go 发布评论及删除评论处理器
│  │      comment_list_handler.go  评论列表处理器
│  │
│  ├─feed
│  │      feed_list_handler.go 视频流处理器
│  │      ffmpeg.exe
│  │      publish_video_handler.go 投稿视频处理器
│  │      query_publish_list_handler.go  发布列表处理器
│  │
│  ├─follow
│  │      follow_action_handler.go 关注处理器
│  │      follow_list_handler.go 关注列表处理器
│  │      query_follower_handler.go  粉丝列表处理器
│  │
│  ├─user
│          user_info_handler.go  用户信息处理器
│          user_login_handler.go 用户登录处理器
│          user_register_handler.go 用户注册处理器
│  │
│  └─video
│          post_favor_hander.go 点赞和取消赞处理器
│          query_favor_videolist_hander.go 喜欢列表处理器
│
├─middleware 中间件，包括用户鉴权功能（验证`Token`同时将`Token`中携带的`user_id`信息写入上下文）
│      AuthMiddleWare.go  用户鉴权
│      normal.go                    不需要认证时
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
│  ├─follow
│  │      follow_list.go 关注列表
│  │      post_follow_action.go 关注功能
│  │
│  ├─user 
│  │      common.go  密码强度检测
│  │      common_test.go 
│  │      user_login.go 用户登录
│  │      user_register.go  用户注册
│  │
│  └─video
│          post_favor_state.go 点赞功能
│          query_favor_videolist.go 喜欢列表
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
        sensitiveWords.go 敏感词检测
        sensitiveWords.txt 敏感词文件
        sensitiveWords_test.go 测试文件
        snowflakes.go 雪花算法
        video.go  雪花id和userid生成独一无二视频名，视频url拼接   
```

### 3.3 项目代码介绍

#### 3.3.1 用户模块

1. 注册操作 /douyin/user/register/
   1. 检测用户名是否已存在
   2. 校验密码强度，要求密码大于8位，至少包含一位数字和大小写字母，可以使用特殊字符
   3. 对密码进行 SHA-1 加密，写入数据库中
   4. 将 user_id 作为 claims 通过 JWT 生成 Token
   5. 将 user_id 和 Token 封装后返回
2. 登录操作 /douyin/user/login/
   1. 对密码进行同样的加密，检测数据库中是否存在该记录
   2. 将 user_id 作为 claims 通过 JWT 生成 Token
   3. 将 Token 和 user_id 进行封装并返回
3. 获取用户信息 /douyin/user/
   1. 使用 JWT 中间件对 Token 进行验证
      - 验证失败：返回失败原因，阻止向下运行
      - 验证成功：将 user_id 存入上下文中
   2. 查询用户信息封装后返回

#### 3.3.2 视频模块

1. 获取视频流 /douyin/feed/
   1. 使用JWT中间件对Token进行验证
      - 验证失败：不进行操作，继续进行
      - 验证成功：将user_id存入上下文中
   2. SQL语句查询最新发布的三十个视频及作者信息
      - 上下文中含有user_id：查询十个视频是否被该用户点赞，作者是否被该用户关注
      - 上下文中不含user_id：默认为视频未点赞，作者未被关注
   3. 封装数据返回信息
2. 发布视频 /douyin/publish/action/
   1. 使用JWT中间件对Token进行验证
      - 验证失败：返回失败原因，阻止向下运行
      - 验证成功：将user_id存入上下文中
   2. 从参数中获取视频及视频信息，从上下文中获取user_id
   3. 采用雪花算法生成uid和user_id拼接成视频名称
   4. 使用`ffmpeg`截取视频第一帧作为封面
   5. 将封面、视频保存至本地的static目录下
   6. 将视频名称、封面名称及视频和作者的相关信息存入数据库中
   7. 封装发布完成情况并返回
3. 查看已发布视频 /douyin/publish/list/ 
   1. 使用 JWT 中间件对 Token 进行验证
      - 验证失败：返回失败原因，阻止向下运行
      - 验证成功：将 user_id 存入上下文中
   2. 从上下文中获取 user_id
   3. 数据库查询已发布的所有视频及本人信息
   4. 封装数据返回信息

#### 3.3.3 关注模块

1. 关注操作 /douyin/relation/action/
   1. 使用JWT中间件对Token进行验证
      - 验证失败：返回失败原因，阻止向下运行
      - 验证成功：将user_id存入上下文中
   2. 对参数进行验证
   3. 根据所传入信息对数据库进行更新（已经关注过），如果数据库没有改条数据，则创建该数据（没有关注过）
   4. 根据查询信息进行数据封装后返回
2.  获取关注列表 /douyin/relation/list/
   1. 使用JWT中间件对Token进行验证
      - 验证失败：返回失败原因，阻止向下运行
      - 验证成功：将user_id存入上下文中
   2. 对参数进行验证
   3. 根据登录用户id使用联合索引获取关注用户id列表，并查询关注用户名和关注总数和粉丝总数
   4. 根据查询信息进行数据封装后返回

#### 3.3.4 评论模块

1. 发布评论 /douyin/comment/action/
   1. 使用 JWT 中间件对 Token 进行验证
      - 验证失败：返回失败原因，阻止向下运行
      - 验证成功：将 user_id 存入上下文中
   2. 根据所传入信息，若 action_type 为1且评论内容存在，数据库中创建该数据，若 action_type 为2且视频 id 存在对数据库进行更新
   3. 根据查询信息进行数据封装后返回
2. 视频评论列表 /douyin/comment/list/
   1. 使用 JWT 中间件对 Token 进行验证
      - 验证失败：返回失败原因，阻止向下运行
      - 验证成功：将 user_id 存入上下文中
   2. 根据 video_id 使用联合查询获取评论信息列表
   3. 根据 user_id 使用联合查询获取用户信息
   4. 根据查询信息进行数据封装后返回

#### 3.3.5 点赞模块

1. 点赞操作 /douyin/favorite/action/
   1. 使用JWT中间件对Token进行验证
      - 验证失败：返回失败原因，阻止向下运行
      - 验证成功：将user_id存入上下文中
   2. 对参数进行验证
   3. 根据所传入信息对数据库进行更新（已经点赞过，取消点赞），如果数据库没有该条数据，则创建该数据（没有点赞过，进行点赞）
   4. 根据查询信息进行数据封装后返回
2. 获取点赞列表 /douyin/favorite/list/
   1. 使用JWT中间件对Token进行验证
      - 验证失败：返回失败原因，阻止向下运行
      - 验证成功：将user_id存入上下文中
   2. 对参数进行验证
   3. 根据登录用户id使用联合查询获取点赞视频信息列表（包含视频作者部分信息），然后查询视频作者关注总数、粉丝总数和是否已关注
   4. 根据查询信息进行数据封装后返回

#### 3.3.6 粉丝，好友列表模块

1.  获取粉丝列表 /douyin/relation/follower/list/
   1. 使用JWT中间件对Token进行验证
      - 验证失败：返回失败原因，阻止向下运行
      - 验证成功：将user_id存入上下文中
   2. 对参数进行验证
   3. 根据登录用户id使用联合索引获取粉丝用户id列表，并查询粉丝用户名和关注总数和粉丝总数
   4. 根据查询信息进行数据封装后返回
2. 获取好友列表 /douyin/relation/friend/list/

# 四、测试结果

> 建议从功能测试和性能测试两部分分析，其中功能测试补充测试用例，性能测试补充性能分析报告、可优化点等内容。

***功能测试为必填**

# 五、Demo 演示视频 （必填）

# 六、项目总结与反思

1. ### 目前仍存在的问题

- 关注列表的人，还显示可以关注
- 发布视频之后发布列表不能直接作品+1,必须重新登录才行

1. ### 已识别出的优化项

- 使用腾讯云cos进行视频和封面的存储
- 使用docker配置数据库
- 采用消息队列
- 使用微服务

1. ### 架构演进的可能性

- 我们采用的是自上而下的架构模式，后期的重构开发存在很大风险和难度，因此我们打算之后将该项目重构，使用微服务框架KiteX,进行二次开发。微服务架构独立性高，而且是进程隔离的，组件之间还可以独立治理，方便我们进行维护。


1. ### 项目过程中的反思与总结

- 我们曾经遇到过一个项目回滚的经历，这让我们团队意识到上传项目并不是自己模块完成，测完就行，一定要看看其他功能是否未受到影响。
- 项目中期遇到过空指针问题导致抖声app崩溃，这让我们明白处理空指针要谨慎小心。
- 抖声项目完成是整个团队协作的功劳，每个人都是第一个使用go语言，做出来实属不易。很幸运和大家相遇，我们高处再见！

# 七、鸣谢

[字节青训营](https://youthcamp.bytedance.com/)
