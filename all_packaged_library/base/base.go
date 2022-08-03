package base

import (
	"github.com/hwholiday/learning_tools/all_packaged_library/base/config"
	"github.com/hwholiday/learning_tools/all_packaged_library/base/db"
	"github.com/hwholiday/learning_tools/all_packaged_library/base/tool"
)

//配置文件的目录
func Init(path string) {
	config.Init(path)
	tool.Init()
	db.Init()
}
