package lock

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func TestRedis(t *testing.T) {
	GlobalClient := redis.NewClient(
		&redis.Options{
			Addr:         "172.12.12.165:6379",
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			Password:     "",
			PoolSize:     10,
			DB:           0,
		},
	)
	ping, err := GlobalClient.Ping().Result()
	if nil != err {
		panic(err)
	}
	fmt.Println("ping", ping)
	redisLock := NewRedisLock(GlobalClient, "test", "1", time.Second*3)
	InitRedis(redisLock)
	select {}
}
func InitRedis(lock RedisLockServer) {
	go func() {
		for {
			time.Sleep(time.Second)
			err := lock.TryLock()
			if err != nil {
				fmt.Println("获取锁失败")
			} else {
				fmt.Println("获取锁成功")
			}
		}
	}()
}
