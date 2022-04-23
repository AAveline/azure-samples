package resolvers

import "ecommerce-project/api-gateway/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	*model.Product
	*model.Order
	*model.User
}
