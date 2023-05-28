package main

import "fmt"

func main() {
	purchases := ReadPurchasesFile("../Fatura_2023-05-10.csv")
	groupedPurchases := AggregatePurchases(purchases)
	client, err := groupedPurchases.ForCardId(6806)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Total para Renan Medina: ", client.TotalAmount)
}
