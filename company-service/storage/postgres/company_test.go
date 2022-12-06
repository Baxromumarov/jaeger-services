package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"jaeger-services/company-service/genproto/company_service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	companyRepo := NewCompanyRepo(dbPool)

	resp, err := companyRepo.Create(context.Background(), &company_service.CreateCompanyRequest{
		Name:        "Test Company-name",
		ProductType: "test-types",
	})

	assert.NoError(t, err)

	fmt.Print("Create Book------->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))
}

func TestGet(t *testing.T) {
	companyRepo := NewCompanyRepo(dbPool)

	resp, err := companyRepo.Get(context.Background(), &company_service.CompanyPrimaryKey{
		Id: "e58ae4cb-a98a-4750-b0aa-11e867d76593",
	})

	assert.NoError(t, err)

	fmt.Print("Get company---->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))
}

func TestGetList(t *testing.T) {
	companyRepo := NewCompanyRepo(dbPool)

	resp, err := companyRepo.GetList(context.Background(), &company_service.GetCompanysListRequest{
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
	companyRepo := NewCompanyRepo(dbPool)

	resp, err := companyRepo.Update(context.Background(), &company_service.UpdateCompanyRequest{
		Company: &company_service.Company{
			Id:          "e58ae4cb-a98a-4750-b0aa-11e867d76593",
			Name:        "Test Update",
			ProductType: "test-type",
		},
	})

	assert.NoError(t, err)

	fmt.Print("Update company---->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))
}

func TestDelete(t *testing.T) {
	companyRepo := NewCompanyRepo(dbPool)

	resp, err := companyRepo.Delete(context.Background(), &company_service.CompanyPrimaryKey{
		Id: "e58ae4cb-a98a-4750-b0aa-11e867d76593",
	})

	assert.NoError(t, err)

	fmt.Print("Delete company---->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))
}
