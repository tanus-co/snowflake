//Snowflake rpc 客户端
package client

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tanus-co/common/util/rpc"
	"github.com/tanus-co/snowflake/rpc/proto"
	"google.golang.org/grpc"
)

const (
	IdWorkerServerPortKey = "snowflake.server-port"
	IdWorkerServerHostKey = "snowflake.server-host"
)

//RPC客户端
type SnowflakeClient struct {
	Address string
}

//创建RPC客户端
func NewSnowflakeClient() *SnowflakeClient {
	port := viper.GetInt(IdWorkerServerPortKey)
	address := viper.GetString(IdWorkerServerHostKey)
	return &SnowflakeClient{
		Address: fmt.Sprintf("%s:%d", address, port),
	}
}

//获取一个id
func (c *SnowflakeClient) GetId() int64 {
	response, err := rpc.RpcRequest(c.Address, func(conn *grpc.ClientConn) (interface{}, error) {
		client := proto.NewGreeterClient(conn)
		response, err := client.GetId(context.Background(), &proto.SnowflakeRequest{Count: 1})
		return response, err
	})
	if err != nil {
		return 0
	}
	snowflakeResponse := response.(*proto.SnowflakeResponse)
	if snowflakeResponse != nil && len(snowflakeResponse.Ids) > 0 {
		return snowflakeResponse.Ids[0]
	} else {
		return 0
	}
}

//获取多个id
func (c *SnowflakeClient) GetIds(count int32) []int64 {
	response, err := rpc.RpcRequest(c.Address, func(conn *grpc.ClientConn) (interface{}, error) {
		client := proto.NewGreeterClient(conn)
		response, err := client.GetId(context.Background(), &proto.SnowflakeRequest{Count: count})
		return response, err
	})
	if err != nil {
		return nil
	}
	snowflakeResponse := response.(*proto.SnowflakeResponse)
	if snowflakeResponse != nil && len(snowflakeResponse.Ids) > 0 {
		return snowflakeResponse.Ids
	} else {
		return nil
	}
}
