package service

import (
	"context"
	"jaeger-services/company-service/config"
	"jaeger-services/company-service/genproto/company_service"
	"jaeger-services/company-service/grpc/client"
	"jaeger-services/company-service/pkg/logger"
	"jaeger-services/company-service/storage"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CompanyService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	company_service.UnimplementedCompanyServiceServer
}

func NewCompanyService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *CompanyService {
	return &CompanyService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (b *CompanyService) CreateCompany(ctx context.Context, req *company_service.CreateCompanyRequest) (resp *company_service.Company, err error) {
	b.log.Info("---CreateCompany--->", logger.Any("req", req))

	pKey, err := b.strg.Company().Create(ctx, req)

	if err != nil {
		b.log.Error("!!!CreateCompany--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return b.strg.Company().Get(ctx, pKey)
}

func (b *CompanyService) GetCompany(ctx context.Context, req *company_service.CompanyPrimaryKey) (resp *company_service.Company, err error) {
	b.log.Info("---GetCompany--->", logger.Any("req", req))

	resp, err = b.strg.Company().Get(ctx, req)

	if err != nil {
		b.log.Error("!!!GetCompany--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *CompanyService) GetCompanysList(ctx context.Context, req *company_service.GetCompanysListRequest) (resp *company_service.GetCompanysListResponse, err error) {
	b.log.Info("---GetCompanysList--->", logger.Any("req", req))

	resp, err = b.strg.Company().GetList(ctx, req)

	if err != nil {
		b.log.Error("!!!GetCompanysList--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *CompanyService) UpdateCompany(ctx context.Context, req *company_service.UpdateCompanyRequest) (resp *company_service.Company, err error) {
	b.log.Info("---UpdateCompany--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Company().Update(ctx, req)

	if err != nil {
		b.log.Error("!!!UpdateCompany--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Company().Get(ctx, &company_service.CompanyPrimaryKey{Id: req.Company.Id})
	if err != nil {
		b.log.Error("!!!UpdateCompany--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}



func (b *CompanyService) DeleteCompany(ctx context.Context, req *company_service.CompanyPrimaryKey) (resp *empty.Empty, err error) {
	b.log.Info("---DeleteCompany--->", logger.Any("req", req))

	resp = &empty.Empty{}

	rowsAffected, err := b.strg.Company().Delete(ctx, req)

	if err != nil {
		b.log.Error("!!!DeleteCompany--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return resp, err
}
