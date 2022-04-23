package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ecommerce-project/api-gateway/graph/model"
)

func (r *federatedMutationResolver) CreateProduct(ctx context.Context, input model.CreateProductInput) (string, error) {
	return r.Product.CreateProduct(ctx, &input)
}

func (r *federatedMutationResolver) UpdateProduct(ctx context.Context, input model.UpdateProductInput) (*model.Product, error) {
	return r.Product.UpdateProduct(ctx, input)
}

func (r *federatedMutationResolver) DeleteProduct(ctx context.Context, input *string) (string, error) {
	return r.Product.DeleteProduct(ctx, *input)
}

func (r *federatedQueryResolver) ListProducts(ctx context.Context) ([]*model.Product, error) {
	return r.Product.ListProducts(ctx)
}

func (r *federatedQueryResolver) GetProduct(ctx context.Context, id string) (*model.Product, error) {
	return r.Product.GetProduct(ctx, id)
}
