package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"jaeger-services/product-service/genproto/product_service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	productRepo := NewProductRepo(dbPool)

	resp, err := productRepo.Create(context.Background(), &product_service.CreateProductRequest{
		ProductName: "Test Product-name",
		// ProductType: "test-types",
	})

	assert.NoError(t, err)

	fmt.Print("Create Book------->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))
}

func TestGet(t *testing.T) {
	productRepo := NewProductRepo(dbPool)

	resp, err := productRepo.Get(context.Background(), &product_service.ProductPrimaryKey{
		Id: "e58ae4cb-a98a-4750-b0aa-11e867d76593",
	})

	assert.NoError(t, err)

	fmt.Print("Get company---->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))
}

func TestGetList(t *testing.T) {
	productRepo := NewProductRepo(dbPool)

	resp, err := productRepo.GetList(context.Background(), &product_service.GetProductsListRequest{
		Page:  1,
		Limit: 10,
	})

	assert.NoError(t, err)

	fmt.Print("Get List---->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))
}

func TestUpdate(t *testing.T) {
	productRepo := NewProductRepo(dbPool)

	resp, err := productRepo.Update(context.Background(), &product_service.UpdateProductRequest{
		Product: &product_service.Product{
			Id:   "e58ae4cb-a98a-4750-b0aa-11e867d76593",
			Name: "Test Update",
			// ProductType: "test-type",
		},
	})

	assert.NoError(t, err)

	fmt.Print("Update company---->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))
}

func TestDelete(t *testing.T) {
	productRepo := NewProductRepo(dbPool)

	resp, err := productRepo.Delete(context.Background(), &product_service.ProductPrimaryKey{
		Id: "e58ae4cb-a98a-4750-b0aa-11e867d76593",
	})

	assert.NoError(t, err)

	fmt.Print("Delete company---->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))
}
