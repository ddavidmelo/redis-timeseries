package grpc_timeseries

import (
	"redis-timeseries/internal/grpc_timeseries/timeseries"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

type Handler struct {
	GrpcServer     *grpc.Server
	GrpcwebHandler *grpcweb.WrappedGrpcServer
}

func Grpc() *Handler {
	grpcServer := grpc.NewServer()

	s := timeseries.Server{}
	timeseries.RegisterTimeSeriesServiceServer(grpcServer, &s)

	// gRPC web code
	grpcWebServer := grpcweb.WrapServer(
		grpcServer,
		// Enable CORS
		// grpcweb.WithOriginFunc(func(origin string) bool { return false }),
		// Disable CORS
		grpcweb.WithAllowNonRootResource(true),
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)

	return &Handler{
		GrpcServer:     grpcServer,
		GrpcwebHandler: grpcWebServer,
	}
}
