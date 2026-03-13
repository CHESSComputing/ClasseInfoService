package main

import (
	"net/http"

	services "github.com/CHESSComputing/golib/services"
	"github.com/gin-gonic/gin"
)

// UserParams represents user-based parameters passed to service
type UserParams struct {
	Id   string `form:"id"`
	Name string `form:"name"`
}

// GetHandler handles queries via GET requests
func GetHandler(c *gin.Context) {
	var params UserParams
	err := c.Bind(&params)
	if err != nil {
		rec := services.Response("ClassID", http.StatusBadRequest, services.BindError, err)
		c.JSON(http.StatusBadRequest, rec)
		return
	}
	var records []map[string]any
	c.JSON(http.StatusOK, records)
}
