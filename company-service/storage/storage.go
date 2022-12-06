package storage

import (
	"jaeger-services/company-service/genproto/company_service"
	"context"
)

type StorageI interface {
	CloseDB()
	Company() CompanyRepoI
}

type CompanyRepoI interface {
	Create(ctx context.Context, req *company_service.CreateCompanyRequest) (resp *company_service.CompanyPrimaryKey, err error)
	Get(ctx context.Context, req *company_service.CompanyPrimaryKey) (resp *company_service.Company, err error)
	GetList(ctx context.Context, req *company_service.GetCompanysListRequest) (resp *company_service.GetCompanysListResponse, err error)
	Update(ctx context.Context, req *company_service.UpdateCompanyRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *company_service.CompanyPrimaryKey) (rowsAffected int64, err error)
}
