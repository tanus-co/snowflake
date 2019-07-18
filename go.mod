module github.com/tanus-co/snowflake

go 1.12

require (
	github.com/etcd-io/bbolt v1.3.3
	github.com/gitstliu/go-id-worker v0.0.0-20190201031057-62250fea083b
	github.com/golang/protobuf v1.3.2
	github.com/spf13/viper v1.4.0
	github.com/tanus-co/common v0.0.1
	google.golang.org/grpc v1.22.0
)

replace (
	cloud.google.com/go v0.26.0 => github.com/googleapis/google-cloud-go v0.26.0
	github.com/etcd-io/bbolt v1.3.2 => go.etcd.io/bbolt v1.3.2
	go.uber.org/multierr v1.1.0 => github.com/uber-go/multierr v1.1.0
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3 => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/net v0.0.0-20181220203305-927f97764cc3 => github.com/golang/net v0.0.0-20190628185345-da137c7871d7
	golang.org/x/net v0.0.0-20190311183353-d8887717615a => github.com/golang/net v0.0.0-20190628185345-da137c7871d7
	golang.org/x/net v0.0.0-20190522155817-f3200d17e092 => github.com/golang/net v0.0.0-20190628185345-da137c7871d7
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sys v0.0.0-20181107165924-66b7b1311ac8 => github.com/golang/sys v0.0.0-20190712062909-fae7ac547cb7
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a => github.com/golang/sys v0.0.0-20190712062909-fae7ac547cb7
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135 => github.com/golang/tools v0.0.0-20190717194535-128ec6dfca09
	google.golang.org/appengine v1.1.0 => github.com/allenlulu/appengine v0.0.0-20171126154002-d35ccace8f27
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 => github.com/ilisin/genproto v0.0.0-20181026194446-8b5d7a19e2d9
	google.golang.org/grpc v1.19.0 => github.com/grpc/grpc-go v1.22.0
	google.golang.org/grpc v1.21.0 => github.com/grpc/grpc-go v1.21.0
	google.golang.org/grpc v1.22.0 => github.com/grpc/grpc-go v1.19.0
)
