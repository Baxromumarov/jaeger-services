package client

import (
	"fmt"
	"jaeger-services/api-gateway/config"
	"jaeger-services/api-gateway/genproto/company_service"

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

	fmt.Println("HERE")

	if err != nil {
		return nil, err
	}

	fmt.Println("HERE 2")

	return &grpcClients{
		companyService: company_service.NewCompanyServiceClient(connCompanyService),
	}, nil
}

func (g *grpcClients) CompanyService() company_service.CompanyServiceClient {
	return g.companyService
}
