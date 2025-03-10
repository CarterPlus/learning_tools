package handler

import (
	"sync"

	"github.com/hwholiday/learning_tools/micro_agent/model/user"
)

var (
	err    error
	server *Service
	m      sync.Mutex
)

type Service struct {
	userServer user.Service
}

func GetService() *Service {
	return server
}

func Init() {
	m.Lock()
	defer m.Unlock()
	server = new(Service)
	server.userServer, err = user.GetService()
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
