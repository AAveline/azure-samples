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

const database = "ecommerce-users"
const collection = "store"

var client *mongo.Client

const port = "3002"

type User struct {
	ID       *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email    string
	FullName string
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

func listUsers(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	cursor, _ := client.Database(database).Collection(collection).Find(context.Background(), bson.D{}, nil)

	defer cursor.Close(ctx)

	var results []User

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	usersToJson, _ := json.Marshal(results)

	out = &common.Content{
		Data:        usersToJson,
		ContentType: in.ContentType,
	}
	return
}

func getUser(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	filter := bson.M{
		"email": string(in.Data),
	}

	result := client.Database(database).Collection(collection).FindOne(context.Background(), filter)

	var user User

	result.Decode(&user)

	userToJson, err := json.Marshal(user)

	if err != nil {
		return
	}

	out = &common.Content{
		Data:        userToJson,
		ContentType: in.ContentType,
	}
	return
}

func createUser(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	var user User

	err = json.Unmarshal(in.Data, &user)

	if err != nil {
		return
	}

	doc := bson.M{
		"fullName": user.FullName,
		"email":    user.Email,
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

func updateUser(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	var user User

	err = json.Unmarshal(in.Data, &user)

	if err != nil {
		return
	}

	doc := bson.D{
		{"$set", bson.D{
			{"fullName", user.FullName},
			{"email", user.Email},
		}},
	}

	filter := bson.M{
		"_id": user.ID,
	}
	result := client.Database(database).Collection(collection).FindOneAndUpdate(ctx, filter, doc)

	if err = result.Decode(&user); err != nil {
		return
	}

	result.Decode(user)

	userToJson, err := json.Marshal(user)

	if err != nil {
		return
	}

	out = &common.Content{
		Data:        userToJson,
		ContentType: in.ContentType,
	}

	return out, err
}

func deleteUser(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
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

func usersHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}

	switch in.Verb {
	case "GET":
		if len(string(in.Data)) == 0 {
			out, err = listUsers(ctx, in)
		} else {
			out, err = getUser(ctx, in)
		}
	case "POST":
		out, err = createUser(ctx, in)
	case "DELETE":
		out, err = deleteUser(ctx, in)
	case "PUT":
		out, err = updateUser(ctx, in)
	}

	return
}

func main() {
	s := daprd.NewService(":" + port)

	if err := connect(); err != nil {
		log.Fatal(err)
	}

	if err := s.AddServiceInvocationHandler("users", usersHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error: %v", err)
	}
}
