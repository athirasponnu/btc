package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// get all roles
func (access *AccessController) GetRoles(ctx *gin.Context) {
	page := ctx.Query("page")
	if page == "" {
		e := "missing query param: page"
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing query param: page ",
		})
		return
	}

	limit := ctx.Query("limit")
	if limit == "" {
		e := "missing query param: limit"
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing query param: limit ",
		})
		return
	}

	Limit, err := strconv.Atoi(limit)
	if err != nil {
		e := fmt.Sprintf("received invalid page query param which is not integer : %v", page)
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "received invalid page query param which is not integer  ",
		})
		return
	}

	if Limit > 25 {
		e := fmt.Sprintf("we agreed to fetch less than %d records but we received request for %d", 25, Limit)
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "exceed limit ",
		})
		return
	}

	Page, err := strconv.Atoi(page)
	if err != nil {
		e := fmt.Sprintf("received invalid offset query param which is not integer : %v", page)
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "received invalid offset query param which is not integer ",
		})
		return
	}

	// check if page or limit is a negative value
	if Page < 0 || Limit < 0 {
		e := " query param cannot be negative"
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": " query param cannot be negative",
		})
		return
	}
	offset := Limit * (Page - 1)
	result, err := access.useCases.GetRoles(Limit, offset)
	switch err {
	case sql.ErrNoRows:
		e := "no rows records found in roles table to read"
		log.Println(e)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "no rows records found ",
		})
	case nil:
		ctx.JSON(http.StatusOK, gin.H{

			"status": "success",
			"data":   result,
		})
	default:
		e := "some internal database server error"
		log.Println(e)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error ",
		})
	}
}
