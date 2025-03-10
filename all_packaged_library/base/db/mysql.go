package db

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/hwholiday/learning_tools/all_packaged_library/base/config"
	"github.com/hwholiday/learning_tools/all_packaged_library/base/tool"
	"go.uber.org/zap"
)

//mysql连接池
func initMysql() {
	var err error
	sql := fmt.Sprintf("%s:%s@(%s:%d)/%s", config.GetMysqlConfig().GetUser(), config.GetMysqlConfig().GetPwd(),
		config.GetMysqlConfig().GetIp(), config.GetMysqlConfig().GetPort(), config.GetMysqlConfig().GetDbName())
	tool.GetLogger().Debug("[initMysql] " + sql)
	mysqlEngine, err = xorm.NewEngine("mysql", sql)
	if err != nil {
		tool.GetLogger().Error("[initMysql] "+sql, zap.Error(err))
		os.Exit(0)
	}
	mysqlEngine.SetMaxOpenConns(config.GetMysqlConfig().GetPoolSize())
	mysqlEngine.SetMaxIdleConns(config.GetMysqlConfig().GetPoolSize())
	if err = mysqlEngine.Ping(); err != nil {
		panic(err)
	}
}

func CloseMysqlConnection() {
	_ = mysqlEngine.Close()
}
