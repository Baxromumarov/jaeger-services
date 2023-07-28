package client


import (
	"jaeger-services/api-gateway/config"
	"jaeger-services/api-gateway/genproto/company_service"
	"jaeger-services/api-gateway/genproto/product_service"

	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	CompanyService() company_service.CompanyServiceClient
	ProductService() product_service.ProductServiceClient
}

type grpcClients struct {
	companyService company_service.CompanyServiceClient
	productService product_service.ProductServiceClient
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

	connProductService, err := grpc.Dial(
		cfg.ProductServiceHost+cfg.ProductServicePort,
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
		productService: product_service.NewProductServiceClient(connProductService),
	}, nil
}

func (g *grpcClients) CompanyService() company_service.CompanyServiceClient {
	return g.companyService
}

func (g *grpcClients) ProductService() product_service.ProductServiceClient {
	return g.productService
}
