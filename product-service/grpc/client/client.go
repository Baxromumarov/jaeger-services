package client

import (
	"jaeger-services/product-service/config"
	"jaeger-services/product-service/genproto/product_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	ProductService() product_service.ProductServiceClient
}

type grpcClients struct {
	productService product_service.ProductServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connProductService, err := grpc.Dial(
		cfg.ServiceHost+cfg.ServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		productService: product_service.NewProductServiceClient(connProductService),
	}, nil
}

func (g *grpcClients) ProductService() product_service.ProductServiceClient {
	return g.productService
}
