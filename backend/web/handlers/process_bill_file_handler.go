package handlers

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/c6-credit-cards-bills-manager/management"
)

func ProcessBillFileHandler(c *gin.Context) {
	receivedFile, _ := c.FormFile("bill_file")
	savedFilePath := uploadsDisk(receivedFile)

	c.SaveUploadedFile(receivedFile, savedFilePath)

	purchases := management.ReadPurchasesFile(savedFilePath)
	groupedPurchases := management.AggregatePurchases(purchases)

	c.JSON(http.StatusOK, gin.H{
		"aggregated": groupedPurchases.ToJson(),
	})
}

func uploadsDisk(destinationFileName string) string {
	return fmt.Sprintf("temp/uploads/%s", destinationFileName)
}
