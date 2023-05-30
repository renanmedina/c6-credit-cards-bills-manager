package main

import (
	"github.com/renanmedina/c6-credit-cards-bills-manager/web/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	webserver := gin.Default()
	webserver.MaxMultipartMemory = 20 << 20 // 20 MiB
	webserver.POST("/process-file", handlers.ProcessBillFileHandler)
	webserver.GET("/purchases", handlers.PurchasesHistoryHandler)
	webserver.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
