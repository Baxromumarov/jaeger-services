package client

import (
	"jaeger-services/product-service/config"
	"jaeger-services/product-service/genproto/company_service"

	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	CompanyService() company_service.CompanyServiceClient
}

type grpcClients struct {
	companyService company_service.CompanyServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connCompanyService, err := grpc.Dial(
		cfg.CompanyServiceHost+cfg.CompanyGRPCPort,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		companyService: company_service.NewCompanyServiceClient(connCompanyService),
	}, nil
}

func (g *grpcClients) CompanyService() company_service.CompanyServiceClient {
	return g.companyService
}
