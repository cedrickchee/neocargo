module github.com/haxorbit/shippy/shippy-cli-user

go 1.14

replace github.com/haxorbit/shippy/shippy-service-user => ../shippy-service-user

// Fix etcd dependency error
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
)
