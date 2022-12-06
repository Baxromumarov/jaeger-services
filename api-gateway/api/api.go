package api

import (
	"jaeger-services/api-gateway/api/docs"
	"jaeger-services/api-gateway/api/handlers"

	"jaeger-services/api-gateway/config"

	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// SetUpRouter godoc
// @description This is a api gateway
// @termsOfService https://udevs.io
// @version 1.0
func SetUpRouter(h handlers.Handler, cfg config.Config, tracer opentracing.Tracer) (r *gin.Engine) {
	r = gin.New()

	r.Use(gin.Logger(), gin.Recovery())
	r.Use(ginhttp.Middleware(tracer))

	docs.SwaggerInfo.Title = cfg.ServiceName
	docs.SwaggerInfo.Version = cfg.Version
	docs.SwaggerInfo.Schemes = []string{cfg.HTTPScheme}

	r.Use(customCORSMiddleware())

	// ! COMPANY
	r.POST("/company", h.CreateCompany)
	r.GET("/company/:id", h.GetCompany)
	r.GET("/company", h.GetCompanyList)
	r.DELETE("/company/:id", h.DeleteCompany)
	r.PUT("/company", h.UpdateCompany)

	// ! PRODUCT
	r.POST("/product", h.CreateProduct)


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
