module go-learning

go 1.14

require (
	github.com/astaxie/beego v1.12.1
	github.com/emicklei/go-restful/v3 v3.2.0
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.0
	github.com/micro/go-plugins/broker/kafka/v2 v2.8.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/registry/consul/v2 v2.8.0
	github.com/micro/go-plugins/v2 v2.0.0 // indirect
	github.com/micro/micro/v2 v2.9.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200620020550-bd6e04640131 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.24.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
