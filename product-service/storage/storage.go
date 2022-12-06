package storage

import (
	"context"
	"jaeger-services/product-service/genproto/product_service"
)

type StorageI interface {
	CloseDB()
	Product() ProductRepoI
}

type ProductRepoI interface {
	Create(ctx context.Context, req *product_service.CreateProductRequest) (resp *product_service.ProductPrimaryKey, err error)
	Get(ctx context.Context, req *product_service.ProductPrimaryKey) (resp *product_service.Product, err error)
	GetList(ctx context.Context, req *product_service.GetProductsListRequest) (resp *product_service.GetProductsListResponse, err error)
	Update(ctx context.Context, req *product_service.UpdateProductRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *product_service.ProductPrimaryKey) (rowsAffected int64, err error)
}
