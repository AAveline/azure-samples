package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ecommerce-project/api-gateway/graph/generated"
	"fmt"
)

func (r *federatedMutationResolver) Default(ctx context.Context, ok *bool) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *federatedQueryResolver) Default(ctx context.Context) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// FederatedMutation returns generated.FederatedMutationResolver implementation.
func (r *Resolver) FederatedMutation() generated.FederatedMutationResolver {
	return &federatedMutationResolver{r}
}

// FederatedQuery returns generated.FederatedQueryResolver implementation.
func (r *Resolver) FederatedQuery() generated.FederatedQueryResolver {
	return &federatedQueryResolver{r}
}

type federatedMutationResolver struct{ *Resolver }
type federatedQueryResolver struct{ *Resolver }
