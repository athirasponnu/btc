package controllers

import (
	"backend-code-base-template/internal/entities"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (access *AccessController) GetRolesById(ctx *gin.Context) {
	var (
		id   int
		err  error
		info entities.Roles
	)
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		e := fmt.Sprintf("received invalid id path param which is not string: %v", ctx.Param("id"))
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "received invalid id",
		})
		return
	}
	// rows := env.DB.QueryRow(queries.GetRoleByIdQ, id)
	info, err = access.useCases.GetRolesById(id)
	if err != nil {
		log.Printf("error occurred while reading the database rows: %v", err)
	}
	switch err {
	case sql.ErrNoRows:
		e := "no rows records found in roles table to read"
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "no rows records found ",
		})
	case nil:
		ctx.JSON(http.StatusOK, gin.H{
			"data":   info,
			"status": "success",
		})
	default:
		e := "some internal database server error"
		log.Println(e)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error ",
		})
	}

}
