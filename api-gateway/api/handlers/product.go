package handlers

import (
	"jaeger-services/api-gateway/api/http"
	"jaeger-services/api-gateway/genproto/company_service"
	"jaeger-services/api-gateway/genproto/product_service"

	"github.com/gin-gonic/gin"
)

// @Router /product [post]
// @Summary Create product
// @Description Create product
// @Tags product
// @Accept json
// @Produce json
// @Param region body product_service.CreateProductRequest true "CreateProductRequest"
// @Success 200 {object} http.Response{data=product_service.Product} "product"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateProduct(c *gin.Context) {
	var product product_service.CreateProductRequest
	err := c.ShouldBindJSON(&product)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services()(
		c.Request.Context(),
		&product,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
