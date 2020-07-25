module github.com/cedrickchee/neocargo/neocargo-service-user

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.2
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.3.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9
	google.golang.org/protobuf v1.25.0
)
