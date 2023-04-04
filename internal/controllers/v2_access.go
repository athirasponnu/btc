package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// V1_HealthHandler
// this will work when the api version provide as v1
func (access *AccessController) V2_HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "server run with v2 version",
	})
}
