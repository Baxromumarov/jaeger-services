package grpc

import (
	"jaeger-services/product-service/config"
	"jaeger-services/product-service/genproto/product_service"
	"jaeger-services/product-service/grpc/client"
	"jaeger-services/product-service/grpc/service"
	"jaeger-services/product-service/pkg/logger"
	"jaeger-services/product-service/storage"

	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer())),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(opentracing.GlobalTracer())),
	)

	product_service.RegisterProductServiceServer(grpcServer, service.NewProductService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)
	return
}
