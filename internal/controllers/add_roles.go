package controllers

import (
	"backend-code-base-template/internal/entities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (access *AccessController) InsertRoles(ctx *gin.Context) {
	var details entities.Role
	if err := ctx.BindJSON(&details); err != nil {
		log.Printf("invalid JSON body: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid JSON body",
		})
		return
	}
	id, err := access.useCases.InsertRoles(details)

	if err != nil {
		log.Printf("error occurred while inserting new record into roles table: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error while inserting",
		})
		return
	}

	// checking the number of rows affected
	// n, err := result.RowsAffected()
	// if err != nil {
	// 	log.Printf("error occurred while checking the returned result from database after insertion: %v", err)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "error occurred while checking the returned result",
	// 	})
	// 	return
	// }

	// if no record was inserted
	if id == 0 {
		e := "could not insert the record, please try again after sometime"
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "could not insert the record",
		})
		return
	}

	m := "successfully created the record"
	log.Println(m)
	ctx.JSON(http.StatusOK, gin.H{
		"message": m,
	})

}
