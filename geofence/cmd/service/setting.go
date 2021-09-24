package service

import (
	"flag"
	"os"
)

var (
	mongoUri      string
	mongoDbName   string
	mongoCollName string

	redisUri string

	httpAddr string
	grpcAddr string
)

func initEnvValues() {
	var fs = flag.NewFlagSet("geofence", flag.ExitOnError)
	fs.Parse(os.Args[1:])

	if httpAddr = os.Getenv("HTTP_ADDR"); httpAddr != "" {
	} else if httpAddr = *fs.String("http-addr", ":8081", "HTTP listen address"); httpAddr != "" {
	} else {
		httpAddr = "0.0.0.0:8081"
	}

	if grpcAddr = os.Getenv("GRPC_ADDR"); grpcAddr != "" {
	} else if grpcAddr = *fs.String("grpc-addr", ":8082", "gRPC listen address"); grpcAddr != "" {
	} else {
		grpcAddr = "0.0.0.0:8082"
	}

	if mongoUri = os.Getenv("MONGO_URI"); mongoUri != "" {
	} else {
		mongoUri = "mongodb://localhost:27017"
	}

	if mongoDbName = os.Getenv("MONGO_DB"); mongoDbName != "" {
	} else {
		mongoDbName = "geo_data"
	}

	if mongoCollName = os.Getenv("MONGO_COLL"); mongoCollName != "" {
	} else {
		mongoCollName = "tehran"
	}

	if redisUri = os.Getenv("REDIS_URI"); redisUri != "" {
	} else {
		redisUri = "localhost:6379"
	}
}
