package postgres

import (
	"context"
	"database/sql"
	"jaeger-services/company-service/config"
	"jaeger-services/company-service/genproto/company_service"
	"jaeger-services/company-service/storage"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"

	"github.com/jackc/pgx/v5"
)

type companyRepo struct {
	db *Pool
}

func NewCompanyRepo(db *Pool) storage.CompanyRepoI {
	return &companyRepo{
		db: db,
	}
}

type Company struct {
	Id        string
	Name      string
	CreatedAt sql.NullString
	UpdatedAt sql.NullString
}

func (b *companyRepo) Create(ctx context.Context, req *company_service.CreateCompanyRequest) (resp *company_service.CompanyPrimaryKey, err error) {
	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Create")
	defer dbSpan.Finish()

	query := `insert into company 
				(id, 
				name
				) VALUES (
					$1, 
					$2
				)`

	uuid, err := uuid.NewRandom()
	if err != nil {
		return resp, err
	}

	_, err = b.db.Exec(ctx, query,
		uuid.String(),
		req.Name,
	)

	if err != nil {
		return resp, err
	}

	resp = &company_service.CompanyPrimaryKey{
		Id: uuid.String(),
	}

	return
}

func (b *companyRepo) Get(ctx context.Context, req *company_service.CompanyPrimaryKey) (resp *company_service.Company, err error) {
	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Get")
	defer dbSpan.Finish()

	result := &Company{}
	resp = &company_service.Company{}

	query := `select 
		id, 
		name, 
		TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	from company 
	where id = $1`

	err = b.db.QueryRow(ctx, query, req.Id).Scan(
		&result.Id,
		&result.Name,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		return resp, err
	}

	if result.CreatedAt.Valid {
		resp.CreatedAt = result.CreatedAt.String
	}

	if result.UpdatedAt.Valid {
		resp.UpdatedAt = result.UpdatedAt.String
	}

	resp.Id = result.Id
	resp.Name = result.Name

	return
}

func (b *companyRepo) GetList(ctx context.Context, req *company_service.GetCompanysListRequest) (resp *company_service.GetCompanysListResponse, err error) {
	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetList")
	defer dbSpan.Finish()

	resp = &company_service.GetCompanysListResponse{}
	var (
		params      (map[string]interface{})
		filter      string
		order       string
		arrangement string
		offset      string
		limit       string
		q           string
	)

	params = map[string]interface{}{}

	query := `select 
				id, 
				name, 
				created_at,
				updated_at
			from company`
	filter = " WHERE true"
	order = " ORDER BY created_at"
	arrangement = " DESC"
	offset = " OFFSET 0"
	limit = " LIMIT 10"

	if req.Page > 0 {
		req.Page = (req.Page - 1) * req.Limit
		params["offset"] = req.Page
		offset = " OFFSET @offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT @limit"
	}

	cQ := `SELECT count(1) FROM company` + filter

	err = b.db.QueryRow(ctx, cQ, pgx.NamedArgs(params)).Scan(
		&resp.Count,
	)

	if err != nil {
		return resp, err
	}

	q = query + filter + order + arrangement + offset + limit

	rows, err := b.db.Query(ctx, q, pgx.NamedArgs(params))
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		book := &company_service.Company{}
		result := &Company{}

		err = rows.Scan(
			&result.Id,
			&result.Name,
			&result.CreatedAt,
			&result.UpdatedAt,
		)

		if err != nil {
			return resp, err
		}

		if result.CreatedAt.Valid {
			book.CreatedAt = result.CreatedAt.String
		}

		if result.UpdatedAt.Valid {
			book.UpdatedAt = result.UpdatedAt.String
		}

		book.Id = result.Id
		book.Name = result.Name

		resp.Companys = append(resp.Companys, book)
	}

	return
}

func (b *companyRepo) Update(ctx context.Context, req *company_service.UpdateCompanyRequest) (rowsAffected int64, err error) {
	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Update")
	defer dbSpan.Finish()

	query := `update company SET
		name = @name,
		updated_at = now()
	WHERE
		id = @id`

	params := map[string]interface{}{
		"id":   req.Company.Id,
		"name": req.Company.Name,
	}

	result, err := b.db.Exec(ctx, query, pgx.NamedArgs(params))
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (b *companyRepo) Delete(ctx context.Context, req *company_service.CompanyPrimaryKey) (rowsAffected int64, err error) {
	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Delete")
	defer dbSpan.Finish()

	query := `delete from company where id = $1`

	result, err := b.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (b *companyRepo) GetWithName(ctx context.Context, req *company_service.CompanyName) (resp *company_service.Company, err error) {
	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Get")
	defer dbSpan.Finish()

	result := &Company{}
	resp = &company_service.Company{}

	query := `select 
		id, 
		name, 
		TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	from company 
	where name = $1`

	err = b.db.QueryRow(ctx, query, req.Name).Scan(
		&result.Id,
		&result.Name,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		return resp, err
	}

	if result.CreatedAt.Valid {
		resp.CreatedAt = result.CreatedAt.String
	}

	if result.UpdatedAt.Valid {
		resp.UpdatedAt = result.UpdatedAt.String
	}

	resp.Id = result.Id
	resp.Name = result.Name

	return
}
