package db

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/hwholiday/learning_tools/all_packaged_library/base/config"
	"github.com/hwholiday/learning_tools/all_packaged_library/base/tool"
	"go.uber.org/zap"
)

func initRedis() {

	redisDb = redis.NewClient(
		&redis.Options{
			Addr:         fmt.Sprintf("%s:%s", config.GetRedisConfig().GetIP(), config.GetRedisConfig().GetPort()),
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			Password:     config.GetRedisConfig().GetPass(),
			PoolSize:     config.GetRedisConfig().GetMaxOpen(),
		},
	)
	err = redisDb.Ping().Err()
	if nil != err {
		tool.GetLogger().Error("ping redis err:", zap.Error(err))
		panic(err)
	}
	tool.GetLogger().Debug("redis success : " + fmt.Sprintf("%s:%s", config.GetRedisConfig().GetIP(), config.GetRedisConfig().GetPort()))

}

func closeRedis() {
	if redisDb != nil {
		_ = redisDb.Close()
	}
}
