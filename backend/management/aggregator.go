package management

import (
	"errors"
	"strings"
)

type PuchasesAggregator struct {
	clients map[string]ClientPurchases
}

type ClientPurchases struct {
	ClientName  string
	CardId      int
	Purchases   []PurchaseItem
	TotalAmount float64
}

func (c *PuchasesAggregator) add(item PurchaseItem) {
	client, exists := c.clients[item.ClientName]

	if !exists {
		client = ClientPurchases{
			ClientName: item.ClientName,
			CardId:     item.CreditCardId,
		}
	}

	c.clients[item.ClientName] = *client.addPurchase(item)
}

func (c *PuchasesAggregator) ForClient(clientName string) (ClientPurchases, error) {
	clientPurchases, exists := c.clients[strings.ToUpper(clientName)]

	if !exists {
		return ClientPurchases{}, errors.New("Cliente não encontrado")
	}

	return clientPurchases, nil
}

func (c *PuchasesAggregator) ForCardId(cardId int) (ClientPurchases, error) {
	for _, clientPurchases := range c.clients {
		if clientPurchases.CardId == cardId {
			return clientPurchases, nil
		}
	}

	return ClientPurchases{}, errors.New("Cliente não encontrado")
}

func (c *PuchasesAggregator) ToJson() map[string]ClientPurchases {
	return c.clients
}

func (p *ClientPurchases) addPurchase(item PurchaseItem) *ClientPurchases {
	p.Purchases = append(p.Purchases, item)
	p.TotalAmount = p.TotalAmount + item.Amount
	return p
}

func AggregatePurchases(purchasesList []PurchaseItem) PuchasesAggregator {
	aggregator := PuchasesAggregator{
		clients: make(map[string]ClientPurchases, 0),
	}

	for _, item := range purchasesList {
		aggregator.add(item)
	}

	return aggregator
}
