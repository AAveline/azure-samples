package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ecommerce-project/api-gateway/graph/model"
)

func (r *federatedMutationResolver) CreateOrder(ctx context.Context, input *model.CreateOrderInput) (string, error) {
	return r.Order.CreateOrder(ctx, *input)
}

func (r *federatedQueryResolver) GetOrders(ctx context.Context, userID string) ([]*model.Order, error) {
	return r.Order.GetOrders(ctx, userID)
}
