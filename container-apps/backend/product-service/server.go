package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const database = "ecommerce-products"
const collection = "store"
const port = "3000"

var client *mongo.Client

type Product struct {
	ID   *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string
}

func connect() (err error) {
	if err = godotenv.Load(); err != nil {
		return
	}

	uri := os.Getenv("MONGODB_CS")
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	c, err := mongo.Connect(ctx, clientOptions)

	client = c

	return
}

func listProducts(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	cursor, _ := client.Database(database).Collection(collection).Find(context.Background(), bson.D{}, nil)

	defer cursor.Close(ctx)

	var results []Product

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	productsToJson, _ := json.Marshal(results)

	out = &common.Content{
		Data:        productsToJson,
		ContentType: in.ContentType,
	}
	return
}

func getProduct(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	id, err := primitive.ObjectIDFromHex(string(in.Data))

	if err != nil {
		return
	}

	filter := bson.M{
		"_id": id,
	}

	result := client.Database(database).Collection(collection).FindOne(context.Background(), filter)

	var product Product
	result.Decode(&product)

	productToJson, err := json.Marshal(product)

	if err != nil {
		return
	}

	out = &common.Content{
		Data:        productToJson,
		ContentType: in.ContentType,
	}
	return
}

func createProduct(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	var product Product

	err = json.Unmarshal(in.Data, &product)

	if err != nil {
		return
	}

	doc := bson.M{
		"name": product.Name,
	}

	_, err = client.Database(database).Collection(collection).InsertOne(ctx, doc)

	if err != nil {
		return
	}

	out = &common.Content{
		Data:        []byte("Success"),
		ContentType: in.ContentType,
	}

	return out, err
}

func updateProduct(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	var product Product

	err = json.Unmarshal(in.Data, &product)

	if err != nil {
		return
	}

	doc := bson.D{
		{"$set", bson.D{
			{"name", product.Name},
		}},
	}

	filter := bson.M{
		"_id": product.ID,
	}
	result := client.Database(database).Collection(collection).FindOneAndUpdate(ctx, filter, doc)

	if err = result.Decode(&product); err != nil {
		return
	}

	result.Decode(product)

	productToJson, err := json.Marshal(product)

	if err != nil {
		return
	}

	out = &common.Content{
		Data:        productToJson,
		ContentType: in.ContentType,
	}

	return out, err
}

func deleteProduct(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	id, err := primitive.ObjectIDFromHex(string(in.Data))

	if err != nil {
		return
	}

	filter := bson.M{
		"_id": id,
	}

	_, err = client.Database(database).Collection(collection).DeleteOne(context.TODO(), filter, nil)

	if err != nil {
		return
	}

	out = &common.Content{
		Data:        []byte("Deleted"),
		ContentType: in.ContentType,
	}

	return
}

func productsHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}

	switch in.Verb {
	case "GET":
		if len(string(in.Data)) == 0 {
			out, err = listProducts(ctx, in)
		} else {
			out, err = getProduct(ctx, in)
		}
	case "POST":
		out, err = createProduct(ctx, in)
	case "DELETE":
		out, err = deleteProduct(ctx, in)
	case "PUT":
		out, err = updateProduct(ctx, in)
	}

	return
}

func main() {
	s := daprd.NewService(":" + port)

	if err := connect(); err != nil {
		log.Fatal(err)
	}

	if err := s.AddServiceInvocationHandler("products", productsHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error: %v", err)
	}
}
