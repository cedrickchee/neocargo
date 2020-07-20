module github.com/haxorbit/shippy/shippy-service-consignment

go 1.14

// replace github.com/haxorbit/shippy/shippy-service-vessel => ../shippy-service-vessel

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.2-0.20200717144925-cce4976773cf // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.6.3 // indirect
	go.mongodb.org/mongo-driver v1.3.5
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200715011427-11fb19a81f2c // indirect
	google.golang.org/grpc v1.30.0 // indirect
	google.golang.org/grpc/examples v0.0.0-20200716233830-6dc7938fe875 // indirect
	google.golang.org/protobuf v1.25.0
)
