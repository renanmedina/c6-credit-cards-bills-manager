package main

import (
	"github.com/renanmedina/c6-credit-cards-bills-manager/web/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	webserver := gin.Default()
	webserver.MaxMultipartMemory = 20 << 20 // 20 MiB

	webserver.Use(CORSMiddleware())
	webserver.POST("/process-file", handlers.ProcessBillFileHandler)
	webserver.GET("/purchases", handlers.PurchasesHistoryHandler)
	webserver.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
