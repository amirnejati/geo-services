module github.com/amirnejati/map-services

go 1.16

require (
	github.com/go-kit/kit v0.10.0
	github.com/go-redis/redis/v8 v8.9.0
	github.com/oklog/oklog v0.3.2
	go.mongodb.org/mongo-driver v1.5.2
	google.golang.org/grpc v1.26.0
	google.golang.org/protobuf v1.26.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.27.0
