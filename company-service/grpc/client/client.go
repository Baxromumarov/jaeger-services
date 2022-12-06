package client

import (
	"fmt"
	"jaeger-services/company-service/config"
	"jaeger-services/company-service/genproto/company_service"
	"jaeger-services/company-service/genproto/product_service"

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
		cfg.ServiceHost+cfg.ServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	connProductService, err := grpc.Dial(
		cfg.ProductServiceHost+cfg.ProductServiceHost,
		grpc.WithInsecure(),
	)
	if err != nil {
		fmt.Println("here")
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
