package handler

import (
	"context"
	"fmt"

	user_agent "github.com/hwholiday/learning_tools/micro_agent/proto/user"
)

func (s *Service) RpcUserInfo(ctx context.Context, req *user_agent.ReqMsg, res *user_agent.ResMsg) error {
	fmt.Println(s.userServer.UserInfo(nil))
	return nil
}
