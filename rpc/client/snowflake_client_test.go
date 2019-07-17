package client

import "testing"

func TestSnowflakeClient_GetId(t *testing.T) {
	client := &SnowflakeClient{
		Address: "127.0.0.1:50051",
	}
	id := client.GetIds(10)
	if len(id) != 10 {
		t.Error(len(id))
	}
	t.Log(client.GetId())

}
