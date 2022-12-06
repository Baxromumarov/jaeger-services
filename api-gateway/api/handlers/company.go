package handlers

import (
	"jaeger-services/api-gateway/api/http"
	"jaeger-services/api-gateway/genproto/company_service"

	"github.com/gin-gonic/gin"
)

// @Router /company [post]
// @Summary Create company
// @Description Create company
// @Tags company
// @Accept json
// @Produce json
// @Param region body company_service.CreateCompanyRequest true "CreateCompanyRequest"
// @Success 200 {object} http.Response{data=company_service.Company} "company"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateCompany(c *gin.Context) {
	var company company_service.CreateCompanyRequest
	err := c.ShouldBindJSON(&company)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.CompanyService().CreateCompany(
		c.Request.Context(),
		&company,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Router /company/{id} [get]
// @Summary get company
// @Description get company
// @Tags company
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=company_service.Company} "company"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetCompany(c *gin.Context) {
	var company company_service.CompanyPrimaryKey

	id := c.Param("id")
	company.Id = id

	resp, err := h.services.CompanyService().GetCompany(
		c.Request.Context(),
		&company,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Get company list
// @Router /company [get]
// @Summary Get company list
// @Description Get company list
// @Tags company
// @Accept json
// @Produce json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} http.Response{data=company_service.GetCompanysListResponse} "company"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetCompanyList(c *gin.Context) {
	page, err := h.getPageParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.CompanyService().GetCompanysList(
		c.Request.Context(),
		&company_service.GetCompanysListRequest{
			Page:  int32(page),
			Limit: int32(limit),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
	}

	// resp, err = h.services.CompanyService().GetCompanysList(
	// 	c.Request.Context(),
	// 	&company_service.GetCompanysListRequest{
	// 		Page:  int32(page),
	// 		Limit: int32(limit),
	// 	},
	// )

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Delete company by id
// @ID delete_company_by_id
// @Router /company/{id} [DELETE]
// @Summary Delete company by id
// @Description Delete company by id
// @Tags company
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Empty
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteCompany(c *gin.Context) {
	var resp http.Empty

	id := c.Param("id")

	_, err := h.services.CompanyService().DeleteCompany(
		c.Request.Context(),
		&company_service.CompanyPrimaryKey{
			Id: id,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Update company
// @ID update_company
// @Router /company [PUT]
// @Summary Update company
// @Description Update company
// @Tags company
// @Accept json
// @Produce json
// @Param Company body company_service.UpdateCompanyRequest true "Request body"
// @Success 200 {object} http.Empty
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateCompany(c *gin.Context) {
	var req company_service.UpdateCompanyRequest
	var resp http.Empty

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err = h.services.CompanyService().UpdateCompany(
		c.Request.Context(),
		&req,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

