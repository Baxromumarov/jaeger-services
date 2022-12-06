package postgres

import (
	"context"
	"database/sql"
	"jaeger-services/product-service/config"
	"jaeger-services/product-service/genproto/product_service"
	"jaeger-services/product-service/storage"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"

	"github.com/jackc/pgx/v5"
)

type productRepo struct {
	db *Pool
}

func NewProductRepo(db *Pool) storage.ProductRepoI {
	return &productRepo{
		db: db,
	}
}

type Product struct {
	Id          string
	Name        string
	ProductType string
	CreatedAt   sql.NullString
	UpdatedAt   sql.NullString
}

func (b *productRepo) Create(ctx context.Context, req *product_service.CreateProductRequest) (resp *product_service.ProductPrimaryKey, err error) {
	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Create")
	defer dbSpan.Finish()

	query := `insert into company 
				(id, 
				name, 
				product_type
				) VALUES (
					$1, 
					$2, 
					$3
				)`

	uuid, err := uuid.NewRandom()
	if err != nil {
		return resp, err
	}

	_, err = b.db.Exec(ctx, query,
		uuid.String(),
		req.Name,
		// req.ProductType
	)

	if err != nil {
		return resp, err
	}

	resp = &product_service.ProductPrimaryKey{
		Id: uuid.String(),
	}

	return
}

func (b *productRepo) Get(ctx context.Context, req *product_service.ProductPrimaryKey) (resp *product_service.Product, err error) {
	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Get")
	defer dbSpan.Finish()

	result := &Product{}
	resp = &product_service.Product{}

	query := `select 
		id, 
		name, 
		product_type,
		TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	from company 
	where id = $1`

	err = b.db.QueryRow(ctx, query, req.Id).Scan(
		&result.Id,
		&result.Name,
		&result.ProductType,
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
	// resp.ProductType = result.ProductType

	return
}

func (b *productRepo) GetList(ctx context.Context, req *product_service.GetProductsListRequest) (resp *product_service.GetProductsListResponse, err error) {
	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetList")
	defer dbSpan.Finish()

	resp = &product_service.GetProductsListResponse{}
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
				product_type,
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
		book := &product_service.Product{}
		result := &Product{}

		err = rows.Scan(
			&result.Id,
			&result.Name,
			&result.ProductType,
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
		// book.ProductType = result.ProductType

		resp.Products = append(resp.Products, book)
	}

	return
}

func (b *productRepo) Update(ctx context.Context, req *product_service.UpdateProductRequest) (rowsAffected int64, err error) {
	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Update")
	defer dbSpan.Finish()

	query := `update company SET
		name = @name,
		product_type = @product_type,
		updated_at = now()
	WHERE
		id = @id`

	params := map[string]interface{}{
		"id":           req.Product.Id,
		"name":         req.Product.Name,
		// "product_type": req.Company.ProductType,
	}

	result, err := b.db.Exec(ctx, query, pgx.NamedArgs(params))
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (b *productRepo) Delete(ctx context.Context, req *product_service.ProductPrimaryKey) (rowsAffected int64, err error) {
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
