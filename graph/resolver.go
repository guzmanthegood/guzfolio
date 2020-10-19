package graph

import "guzfolio/datastore"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DS datastore.DataStore
}