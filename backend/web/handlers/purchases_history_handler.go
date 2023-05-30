package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PurchasesHistoryHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": make([]interface{}, 0),
	})
}
