module github.com/tanus-co/snowflake

go 1.12

require (
	github.com/gitstliu/go-id-worker v0.0.0-20190201031057-62250fea083b
	github.com/golang/protobuf v1.3.2
	github.com/spf13/viper v1.4.0
	github.com/tanus-co/common v0.0.1
	google.golang.org/grpc v1.22.0
)

replace (
	golang.org/x/net@v0.0.0-20190522155817-f3200d17e092 => github.com/golang/net v0.0.0-20190628185345-da137c7871d7
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	google.golang.org/grpc v1.21.0 => github.com/grpc/grpc-go v1.22.0
	google.golang.org/grpc v1.22.0 => github.com/grpc/grpc-go v1.22.0
)
