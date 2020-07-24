module github.com/haxorbit/neocargo/neocargo-cli-user

go 1.14

// replace github.com/haxorbit/neocargo/neocargo-service-user => ../neocargo-service-user

// Fix etcd dependency error
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
)
