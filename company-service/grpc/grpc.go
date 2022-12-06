package grpc

import (
	"jaeger-services/company-service/config"
	"jaeger-services/company-service/genproto/company_service"
	"jaeger-services/company-service/genproto/product_service"
	"jaeger-services/company-service/grpc/client"
	"jaeger-services/company-service/grpc/service"
	"jaeger-services/company-service/pkg/logger"
	"jaeger-services/company-service/storage"

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

	company_service.RegisterCompanyServiceServer(grpcServer, service.NewCompanyService(cfg, log, strg, svcs))
	product_service.RegisterProductServiceServer(grpcServer, service.NewProductService(cfg, log, strg, svcs))


	reflection.Register(grpcServer)
	return
}

