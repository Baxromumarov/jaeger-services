package handlers

import (
	"jaeger-services/api-gateway/api/http"
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

	resp, err := h.services.ProductService().CreateProduct(
		c.Request.Context(),
		&product,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @Router /product/{id} [get]
// @Summary get product
// @Description get product
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=product_service.Product} "product"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetProduct(c *gin.Context) {
	var product product_service.ProductPrimaryKey

	id := c.Param("id")
	product.Id = id

	resp, err := h.services.ProductService().GetProduct(
		c.Request.Context(),
		&product,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Get product list
// @Router /product [get]
// @Summary Get product list
// @Description Get product list
// @Tags product
// @Accept json
// @Produce json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} http.Response{data=product_service.GetProductsListResponse} "product"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetProductList(c *gin.Context) {
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

	resp, err := h.services.ProductService().GetProductsList(
		c.Request.Context(),
		&product_service.GetProductsListRequest{
			Page:  int32(page),
			Limit: int32(limit),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
	}

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Delete product by id
// @ID delete_product_by_id
// @Router /product/{id} [DELETE]
// @Summary Delete product by id
// @Description Delete product by id
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Empty
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteProduct(c *gin.Context) {
	var resp http.Empty

	id := c.Param("id")

	_, err := h.services.ProductService().DeleteProduct(
		c.Request.Context(),
		&product_service.ProductPrimaryKey{
			Id: id,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Update product
// @ID update_product
// @Router /product [PUT]
// @Summary Update product
// @Description Update product
// @Tags product
// @Accept json
// @Produce json
// @Param product body product_service.UpdateProductRequest true "Request body"
// @Success 200 {object} http.Empty
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateProduct(c *gin.Context) {
	var (
		req  product_service.UpdateProductRequest
		resp http.Empty
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err = h.services.ProductService().UpdateProduct(
		c.Request.Context(),
		&req,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
