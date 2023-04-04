package controllers

import (
	"backend-code-base-template/internal/entities"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (access *AccessController) UpdateRoles(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		e := fmt.Sprintf("received invalid id path param which is not string: %v", ctx.Param("id"))
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "received invalid id",
		})
		return
	}
	var updateDetails entities.Role
	if err := ctx.BindJSON(&updateDetails); err != nil {
		log.Printf("invalid JSON body: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid JSON body",
		})
		return
	}

	result, err := access.useCases.UpdateRoles(updateDetails, id)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating roles record ", err)
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error occurred while updating roles ",
		})
		return
	}
	// checking the number of rows affected
	n, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred while checking the returned result from database after updation: %v", err)
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error occurred while checking the returned result from database after updation ",
		})
	}
	// if no record was updated
	if n == 0 {
		e := "could not update the record, please try again after sometime"
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not update the record, please try again after sometime",
		})
		return
	}

	m := "successfully updated the record"
	log.Println(m)
	ctx.JSON(http.StatusOK, gin.H{
		"message": m,
	})

}
