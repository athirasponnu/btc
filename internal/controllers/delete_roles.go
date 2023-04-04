package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (access *AccessController) DeleteRoles(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		e := fmt.Sprintf("received invalid id path param which is not string: %v", ctx.Param("id"))
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "received invalid id",
		})
		return
	}
	result, err := access.useCases.DeleteRoles(id)
	if err != nil {
		e := fmt.Sprintf(
			"error occurred while deleting the record with id: %d and error is: %v", id, err)
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": e,
		})
		return
	}

	// checking the number of rows affected
	n, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf(
			"error occurred while checking the returned result from database after deletion: %v", err)
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": e,
		})
		return
	}

	// if no record was deleted, let us inform that there might be no
	// records to delete for this given data.
	if n == 0 {
		e := "could not delete the record, there might be no records for the given ID"
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": e,
		})
		return
	}

	m := "successfully deleted the record"
	log.Println(m)
	ctx.JSON(http.StatusOK, gin.H{
		"message": m,
	})
}
