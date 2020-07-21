module github.com/haxorbit/shippy/shippy-cli-consignment

go 1.14

// replace github.com/haxorbit/shippy/shippy-service-consignment => ../shippy-service-consignment

// Fix etcd dependency error
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/grpc v1.30.0 // indirect
	google.golang.org/grpc/examples v0.0.0-20200716233830-6dc7938fe875 // indirect
)
