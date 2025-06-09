package utils

import (
	"net/http"

	"github.com/ethanbaker/api/pkg/api_types"
	"github.com/gin-gonic/gin"
)

// Handle requests with no route
func NoRouteHandler(c *gin.Context) {
	res := api_types.NewFailResponse(http.StatusNotFound, "Route not found")
	c.JSON(res.AsGinResponse())
}
