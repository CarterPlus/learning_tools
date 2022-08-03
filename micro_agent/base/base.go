package base

import (
	"github.com/hwholiday/learning_tools/micro_agent/base/config"
	"github.com/hwholiday/learning_tools/micro_agent/base/db"
	"github.com/hwholiday/learning_tools/micro_agent/base/tool"
)

func Init(path string) {
	config.Init(path)
	db.Init()
	tool.Init()
}
