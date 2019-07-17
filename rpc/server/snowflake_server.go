package server

import (
	"context"
	"errors"
	"github.com/tanus-co/snowflake/rpc/proto"
	"github.com/tanus-co/snowflake/service"
)

//ID Server
type SnowflakeServer struct {
}

//获取ID的RPC服务
func (s *SnowflakeServer) GetId(ctx context.Context, request *proto.SnowflakeRequest) (*proto.SnowflakeResponse, error) {
	response := new(proto.SnowflakeResponse)
	if request.Count > 0 {
		if ids, err := service.GetIds(int(request.Count)); err == nil {
			response.Ids = ids
		} else {
			return nil, err
		}
	} else {
		return nil, errors.New("count must greater than 0")
	}
	return response, nil
}
