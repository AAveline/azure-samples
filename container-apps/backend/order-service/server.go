package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

const database = "ecommerce-orders"
const collection = "store"
const port = "3001"

var client *mongo.Client

type Order struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductID string
	UserID    string
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

func getOrders(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	filter := bson.M{
		"userId": string(in.Data),
	}

	result, err := client.Database(database).Collection(collection).Find(context.Background(), filter)

	if err != nil {
		return
	}

	var orders []Order

	result.Decode(&orders)

	ordersToJson, err := json.Marshal(orders)

	if err != nil {
		return
	}

	out = &common.Content{
		Data:        ordersToJson,
		ContentType: in.ContentType,
	}
	return
}

func createOrder(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	var order Order

	json.Unmarshal(in.Data, &order)

	doc := bson.M{
		"userId":    order.UserID,
		"productId": order.ProductID,
	}

	a, err := client.Database(database).Collection(collection).InsertOne(ctx, doc)

	fmt.Print(a.InsertedID)
	if err != nil {
		return
	}

	out = &common.Content{
		Data:        []byte("Success"),
		ContentType: in.ContentType,
	}
	return
}

func ordersHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}

	switch in.Verb {
	case "GET":
		out, err = getOrders(ctx, in)
	case "POST":
		out, err = createOrder(ctx, in)
	}

	return
}

func main() {
	s := daprd.NewService(":" + port)

	if err := connect(); err != nil {
		log.Fatal(err)
	}

	if err := s.AddServiceInvocationHandler("orders", ordersHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error: %v", err)
	}
}
