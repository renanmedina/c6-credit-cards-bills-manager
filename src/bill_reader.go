package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type PurchaseItem struct {
	Date         string
	ClientName   string
	CreditCardId int
	Category     string
	Description  string
	Installment  string
	Amount       float64
}

func NewPurchaseItem(date string, client string, cardId int, category string, description string, installment string, amount float64) PurchaseItem {
	return PurchaseItem{
		Date:         date,
		ClientName:   client,
		CreditCardId: cardId,
		Category:     category,
		Description:  description,
		Installment:  installment,
		Amount:       amount,
	}
}

func createFromRecordString(record []string) PurchaseItem {
	cardId, _ := strconv.Atoi(record[2])
	amountParsed := strings.ReplaceAll(record[8], ",", ".")
	amount, _ := strconv.ParseFloat(amountParsed, 64)

	return NewPurchaseItem(
		record[0],
		record[1],
		cardId,
		record[3],
		record[4],
		record[5],
		amount,
	)
}

func ReadPurchasesFile(filepath string) []PurchaseItem {
	fileOpened, err := os.Open(filepath)
	defer fileOpened.Close()

	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}

	fileReader := csv.NewReader(fileOpened)
	fileReader.Comma = ';'

	var purchases []PurchaseItem
	isHeader := false

	for {
		record, err := fileReader.Read()

		if err == io.EOF {
			break
		}

		if !isHeader {
			isHeader = true
			continue
		}

		purchase := createFromRecordString(record)
		purchases = append(purchases, purchase)
	}

	return purchases
}
