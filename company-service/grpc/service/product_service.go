package service

import (
	"context"
	"jaeger-services/company-service/config"
	"jaeger-services/company-service/genproto/product_service"
	"jaeger-services/company-service/grpc/client"
	"jaeger-services/company-service/pkg/logger"
	"jaeger-services/company-service/storage"

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

	pKey, err := b.services.ProductService().CreateProduct(ctx, req)

	if err != nil {
		b.log.Error("!!!CreateProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return b.services.ProductService().GetProduct(ctx,&product_service.ProductPrimaryKey{Id: pKey.Id})
}

// func (b *ProductService) GetProduct(ctx context.Context, req *product_service.CompanyPrimaryKey) (resp *product_service.Company, err error) {
// 	b.log.Info("---GetCompany--->", logger.Any("req", req))

// 	resp, err = b.strg.Company().Get(ctx, req)

// 	if err != nil {
// 		b.log.Error("!!!GetCompany--->", logger.Error(err))
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}

// 	return resp, err
// }

// func (b *ProductService) GetCompanysList(ctx context.Context, req *product_service.GetCompanysListRequest) (resp *product_service.GetCompanysListResponse, err error) {
// 	b.log.Info("---GetCompanysList--->", logger.Any("req", req))

// 	resp, err = b.strg.Company().GetList(ctx, req)

// 	if err != nil {
// 		b.log.Error("!!!GetCompanysList--->", logger.Error(err))
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}

// 	return resp, err
// }

// func (b *ProductService) UpdateCompany(ctx context.Context, req *product_service.UpdateCompanyRequest) (resp *product_service.Company, err error) {
// 	b.log.Info("---UpdateCompany--->", logger.Any("req", req))

// 	rowsAffected, err := b.strg.Company().Update(ctx, req)

// 	if err != nil {
// 		b.log.Error("!!!UpdateCompany--->", logger.Error(err))
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}

// 	if rowsAffected <= 0 {
// 		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
// 	}

// 	resp, err = b.strg.Company().Get(ctx, &product_service.CompanyPrimaryKey{Id: req.Company.Id})
// 	if err != nil {
// 		b.log.Error("!!!UpdateCompany--->", logger.Error(err))
// 		return nil, status.Error(codes.NotFound, err.Error())
// 	}

// 	return resp, err
// }

// func (b *ProductService) DeleteCompany(ctx context.Context, req *product_service.CompanyPrimaryKey) (resp *empty.Empty, err error) {
// 	b.log.Info("---DeleteCompany--->", logger.Any("req", req))

// 	resp = &empty.Empty{}

// 	rowsAffected, err := b.strg.Company().Delete(ctx, req)

// 	if err != nil {
// 		b.log.Error("!!!DeleteCompany--->", logger.Error(err))
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}

// 	if rowsAffected <= 0 {
// 		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
// 	}

// 	return resp, err
// }
