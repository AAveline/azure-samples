package model

import (
	"context"
	"encoding/json"

	dapr "github.com/dapr/go-sdk/client"
)

const daprPort = "50001"

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (p *Product) ListProducts(ctx context.Context) ([]*Product, error) {
	client, err := dapr.NewClientWithPort(daprPort)

	if err != nil {
		panic(err)
	}
	defer client.Close()

	content := &dapr.DataContent{
		ContentType: "application/json",
	}

	response, err := client.InvokeMethodWithContent(ctx, "product-service", "products", "get", content)

	if err != nil {
		panic(err)
	}

	var products []*Product

	err = json.Unmarshal(response, &products)

	if err != nil {
		panic(err)
	}

	return products, nil
}

func (p *Product) GetProduct(ctx context.Context, id string) (*Product, error) {
	client, err := dapr.NewClientWithPort(daprPort)

	if err != nil {
		panic(err)
	}

	defer client.Close()

	content := &dapr.DataContent{
		ContentType: "application/json",
		Data:        []byte(id),
	}

	response, err := client.InvokeMethodWithContent(ctx, "product-service", "products", "get", content)

	if err != nil {
		panic(err)
	}

	var product *Product

	err = json.Unmarshal(response, &product)

	if err != nil {
		panic(err)
	}

	return product, nil
}

func (p *Product) CreateProduct(ctx context.Context, input *CreateProductInput) (string, error) {
	client, err := dapr.NewClientWithPort(daprPort)

	if err != nil {
		panic(err)
	}

	defer client.Close()

	data, err := json.Marshal(input)

	if err != nil {
		panic(err)
	}

	content := &dapr.DataContent{
		ContentType: "application/json",
		Data:        data,
	}

	response, err := client.InvokeMethodWithContent(ctx, "product-service", "products", "post", content)

	if err != nil {
		panic(err)
	}

	return string(response), nil
}

func (p *Product) DeleteProduct(ctx context.Context, id string) (string, error) {
	client, err := dapr.NewClientWithPort(daprPort)

	if err != nil {
		panic(err)
	}

	defer client.Close()

	content := &dapr.DataContent{
		ContentType: "application/json",
		Data:        []byte(id),
	}

	response, err := client.InvokeMethodWithContent(ctx, "product-service", "products", "delete", content)

	if err != nil {
		panic(err)
	}

	return string(response), nil
}

func (p *Product) UpdateProduct(ctx context.Context, input UpdateProductInput) (*Product, error) {
	client, err := dapr.NewClientWithPort(daprPort)

	if err != nil {
		panic(err)
	}

	defer client.Close()

	data, err := json.Marshal(input)

	if err != nil {
		panic(err)
	}

	content := &dapr.DataContent{
		ContentType: "application/json",
		Data:        data,
	}

	response, err := client.InvokeMethodWithContent(ctx, "product-service", "products", "put", content)

	if err != nil {
		panic(err)
	}

	var product *Product

	err = json.Unmarshal(response, &product)

	if err != nil {
		panic(err)
	}

	return product, nil
}
