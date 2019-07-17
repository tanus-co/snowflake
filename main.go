package main

import (
	_ "github.com/tanus-co/common/config"
	"github.com/tanus-co/snowflake/rpc/server"
)

func main() {
	server := server.NewServer()
	server.Listen()
}
