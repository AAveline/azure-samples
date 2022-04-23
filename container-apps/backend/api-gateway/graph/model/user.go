package model

import (
	"context"
	"encoding/json"

	dapr "github.com/dapr/go-sdk/client"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}

func setupDaprClient() (client dapr.Client) {
	client, err := dapr.NewClientWithPort(daprPort)

	if err != nil {
		panic(err)
	}

	return
}

func (u *User) ListUsers(ctx context.Context) ([]*User, error) {
	client := setupDaprClient()

	defer client.Close()

	content := &dapr.DataContent{
		ContentType: "application/json",
	}

	response, err := client.InvokeMethodWithContent(ctx, "user-service", "users", "get", content)

	if err != nil {
		panic(err)
	}

	var users []*User

	err = json.Unmarshal(response, &users)

	if err != nil {
		panic(err)
	}

	return users, nil
}

func (u *User) GetUser(ctx context.Context, email string) (*User, error) {
	client := setupDaprClient()

	defer client.Close()

	content := &dapr.DataContent{
		ContentType: "application/json",
		Data:        []byte(email),
	}

	response, err := client.InvokeMethodWithContent(ctx, "user-service", "users", "get", content)

	if err != nil {
		panic(err)
	}

	var user *User

	err = json.Unmarshal(response, &user)

	if err != nil {
		panic(err)
	}

	return user, nil
}

func (u *User) CreateUser(ctx context.Context, input *CreateUserInput) (string, error) {
	client := setupDaprClient()

	defer client.Close()

	data, err := json.Marshal(input)

	if err != nil {
		panic(err)
	}

	content := &dapr.DataContent{
		ContentType: "application/json",
		Data:        data,
	}

	response, err := client.InvokeMethodWithContent(ctx, "user-service", "users", "post", content)

	if err != nil {
		panic(err)
	}

	return string(response), nil
}

func (u *User) DeleteUser(ctx context.Context, id string) (string, error) {
	client := setupDaprClient()

	defer client.Close()

	content := &dapr.DataContent{
		ContentType: "application/json",
		Data:        []byte(id),
	}

	response, err := client.InvokeMethodWithContent(ctx, "user-service", "users", "delete", content)

	if err != nil {
		panic(err)
	}

	return string(response), nil
}

func (u *User) UpdateUser(ctx context.Context, input *UpdateUserInput) (*User, error) {
	client := setupDaprClient()

	defer client.Close()

	data, err := json.Marshal(input)

	if err != nil {
		panic(err)
	}

	content := &dapr.DataContent{
		ContentType: "application/json",
		Data:        data,
	}

	response, err := client.InvokeMethodWithContent(ctx, "user-service", "users", "put", content)

	if err != nil {
		panic(err)
	}

	var user *User

	err = json.Unmarshal(response, &user)

	if err != nil {
		panic(err)
	}

	return user, nil
}
