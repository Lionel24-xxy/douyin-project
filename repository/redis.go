package repository

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
var CTX = context.Background()
var RDB *redis.Client

const (
	favor    = "favor"
	relation = "relation"
)

// 根据redis配置初始化一个客户端
func InitRedisClient() (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", 	// redis地址
		Password: "xxx",          // redis密码，没有则留空
		DB:       0,                	// 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = RDB.Ping(CTX).Result()
	if err != nil {
		return err
	}
	return nil
}

// GetVideoFavorState 得到点赞状态
func GetVideoFavorState(userId int64, videoId int64) bool {
	key := fmt.Sprintf("%s:%d", favor, userId)
	ret := RDB.SIsMember(CTX, key, videoId)
	return ret.Val()
}

// GetUserRelation 得到关注状态
func GetUserRelation(userId int64, followId int64) bool {
	key := fmt.Sprintf("%s:%d", relation, userId)
	ret := RDB.SIsMember(CTX, key, followId)
	return ret.Val()
}

// UpdateUserRelation 更新点赞状态，state:true为点关注，false为取消关注
func UpdateUserRelation(userId int64, followId int64, state bool) {
	key := fmt.Sprintf("%s:%d", relation, userId)
	if state {
		RDB.SAdd(CTX, key, followId) //添加集合元素
		return
	}
	RDB.SRem(CTX, key, followId) //删除集合元素
}
