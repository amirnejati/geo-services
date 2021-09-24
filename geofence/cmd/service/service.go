package service

import (
	"fmt"
	"github.com/amirnejati/map-services/geofence/pkg/endpoint"
	grpctransport "github.com/amirnejati/map-services/geofence/pkg/grpc"
	"github.com/amirnejati/map-services/geofence/pkg/grpc/pb"
	httptransport "github.com/amirnejati/map-services/geofence/pkg/http"
	"github.com/amirnejati/map-services/geofence/pkg/service"
	kitendpoint "github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/oklog/oklog/pkg/group"
	googlegrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var logger kitlog.Logger

func Run() {
	initEnvValues()

	// Create a single logger, which we'll use and give to other components.
	logger = kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp)
	logger = kitlog.With(logger, "caller", kitlog.DefaultCaller)

	repoMongo := initMongo()
	svc := service.NewGeofenceServiceMiddlewareEnabled(repoMongo, logger, getServiceMiddleware(logger))
	eps := endpoint.NewEndpointsSet(svc, getEndpointMiddleware(logger))
	g := createService(eps)
	initCancelInterrupt(g)
	logger.Log("exit", g.Run())
}

func getServiceMiddleware(logger kitlog.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}
	mw = addDefaultServiceMiddleware(logger, mw)
	// Append your middleware here
	return
}

func getEndpointMiddleware(logger kitlog.Logger) (mw map[string][]kitendpoint.Middleware) {
	mw = map[string][]kitendpoint.Middleware{}
	addDefaultEndpointMiddleware(logger, mw)
	// Add you endpoint middleware here
	return
}

func initHttpHandler(endpoints endpoint.EndpointsSet, g *group.Group) {
	options := defaultHttpOptions(logger)
	// Add your httptransport options here

	httpHandler := httptransport.NewHTTPHandler(endpoints, options)
	httpListener, err := net.Listen("tcp", httpAddr)
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
		os.Exit(1)
	}
	g.Add(func() error {
		logger.Log("transport", "HTTP", "addr", httpAddr)
		return http.Serve(httpListener, httpHandler)
	}, func(error) {
		httpListener.Close()
	})
}

func initGRPCHandler(endpoints endpoint.EndpointsSet, g *group.Group) {
	//options := defaultGRPCOptions(logger)
	// Add your grpctransport options here

	grpcServer := grpctransport.NewGRPCServer(endpoints, nil)
	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}
	g.Add(func() error {
		logger.Log("transport", "gRPC", "addr", grpcAddr)
		baseServer := googlegrpc.NewServer(googlegrpc.UnaryInterceptor(kitgrpc.Interceptor))
		pb.RegisterGeofenceServer(baseServer, grpcServer)
		reflection.Register(baseServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})
}

func initMongo() service.RepositoryMongo {
	GetMongoDB()
	repo, err := service.NewRepoMongo(dbMongo, mongoCollName, logger)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}
	return repo
}

func initCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}
