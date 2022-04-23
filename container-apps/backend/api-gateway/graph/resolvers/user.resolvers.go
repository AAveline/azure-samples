package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ecommerce-project/api-gateway/graph/model"
)

func (r *federatedMutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (string, error) {
	return r.User.CreateUser(ctx, &input)
}

func (r *federatedMutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	return r.User.UpdateUser(ctx, &input)
}

func (r *federatedMutationResolver) DeleteUser(ctx context.Context, input string) (string, error) {
	return r.User.DeleteUser(ctx, input)
}

func (r *federatedQueryResolver) ListUsers(ctx context.Context) ([]*model.User, error) {
	return r.User.ListUsers(ctx)
}

func (r *federatedQueryResolver) GetUser(ctx context.Context, email string) (*model.User, error) {
	return r.User.GetUser(ctx, email)
}
