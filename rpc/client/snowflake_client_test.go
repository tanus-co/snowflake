package client

import "testing"

func TestSnowflakeClient_GetId(t *testing.T) {
	client := &SnowflakeClient{
		Address: "snowflake.rpc.tanus.iooc.cc:10016",
	}
	id := client.GetIds(10)
	if len(id) != 10 {
		t.Error(len(id))
	}
	t.Log(client.GetId())

}
