package model

import (
	"context"
	"encoding/json"

	dapr "github.com/dapr/go-sdk/client"
)

type Order struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	ProductID string `json:"productId"`
}

func (o *Order) GetOrders(ctx context.Context, userID string) ([]*Order, error) {
	client, err := dapr.NewClientWithPort(daprPort)

	if err != nil {
		panic(err)
	}
	//	defer client.Close()

	content := &dapr.DataContent{
		ContentType: "application/json",
		Data:        []byte(userID),
	}

	response, err := client.InvokeMethodWithContent(ctx, "order-service", "orders", "get", content)

	if err != nil {
		panic(err)
	}

	var orders []*Order

	err = json.Unmarshal(response, &orders)

	if err != nil {
		panic(err)
	}

	return orders, nil
}

func (o *Order) CreateOrder(ctx context.Context, input CreateOrderInput) (string, error) {
	client, err := dapr.NewClient()

	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(input)

	if err != nil {
		panic(err)
	}

	content := &dapr.DataContent{
		ContentType: "application/json",
		Data:        data,
	}

	response, err := client.InvokeMethodWithContent(ctx, "order-service", "orders", "post", content)

	if err != nil {
		panic(err)
	}

	return string(response), nil
}
