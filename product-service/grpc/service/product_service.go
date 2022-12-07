package service

import (
	"context"
	"jaeger-services/product-service/config"
	"jaeger-services/product-service/genproto/company_service"
	"jaeger-services/product-service/genproto/product_service"
	"jaeger-services/product-service/grpc/client"
	"jaeger-services/product-service/pkg/logger"
	"jaeger-services/product-service/storage"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	product_service.UnimplementedProductServiceServer
}

func NewProductService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *ProductService {
	return &ProductService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (b *ProductService) CreateProduct(ctx context.Context, req *product_service.CreateProductRequest) (resp *product_service.Product, err error) {
	b.log.Info("---CreateProduct--->", logger.Any("req", req))

	company, err := b.services.CompanyService().GetCompanyWithName(ctx, &company_service.CompanyName{Name: req.CompanyName})
	if err != nil {
		b.log.Error("!!!Getcompany with name--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req.CompanyId = company.Id

	pKey, err := b.strg.Product().Create(ctx, req)
	if err != nil {
		b.log.Error("!!!CreateProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return b.strg.Product().Get(ctx, pKey)
}

func (b *ProductService) GetProduct(ctx context.Context, req *product_service.ProductPrimaryKey) (resp *product_service.Product, err error) {
	b.log.Info("---GetProduct--->", logger.Any("req", req))

	resp, err = b.strg.Product().Get(ctx, req)

	if err != nil {
		b.log.Error("!!!GetProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *ProductService) GetProductsList(ctx context.Context, req *product_service.GetProductsListRequest) (resp *product_service.GetProductsListResponse, err error) {
	b.log.Info("---GetProductsList--->", logger.Any("req", req))

	resp, err = b.strg.Product().GetList(ctx, req)

	if err != nil {
		b.log.Error("!!!GetProductsList--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *ProductService) UpdateProduct(ctx context.Context, req *product_service.UpdateProductRequest) (resp *product_service.Product, err error) {
	b.log.Info("---UpdateProduct--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Product().Update(ctx, req)

	if err != nil {
		b.log.Error("!!!UpdateProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Product().Get(ctx, &product_service.ProductPrimaryKey{Id: req.Product.Id})
	if err != nil {
		b.log.Error("!!!UpdateProduct--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (b *ProductService) DeleteProduct(ctx context.Context, req *product_service.ProductPrimaryKey) (resp *empty.Empty, err error) {
	b.log.Info("---DeleteProduct--->", logger.Any("req", req))

	resp = &empty.Empty{}

	rowsAffected, err := b.strg.Product().Delete(ctx, req)

	if err != nil {
		b.log.Error("!!!DeleteProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return resp, err
}

func (b *ProductService) GetCompany(ctx context.Context, req *product_service.CompanyPrimaryKey) (resp *product_service.Product, err error) {
	b.log.Info("---GetCompany--->", logger.Any("req", req))

	resp, err = b.strg.Product().GetCompany(ctx, req)

	if err != nil {
		b.log.Error("!!!GetCompany--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}
