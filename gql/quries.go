package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/theShivaa/go-gql-crud/postgres"
)

// Root holds a pointer to a graphql object

// NewRoot returns base query type. This is where we add all the base queries
func RootQueries(db *postgres.Db) *graphql.Object {
	// Create a resolver holding our databse. Resolver can be found in resolvers.go
	resolver := Resolver{db: db}

	// Create a new Root that describes our base query set up. In this
	// example we have a user query that takes one argument called name

	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"users": &graphql.Field{
					// Slice of User type which can be found in types.go
					Type: graphql.NewList(User),
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: resolver.GetUserByName,
				},
				"user": &graphql.Field{
					Type: User,
					Args: graphql.FieldConfigArgument{
						"userid": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: resolver.GetUserByUserID,
				},
				"allusers": &graphql.Field{
					Type:    graphql.NewList(User),
					Resolve: resolver.GetAllUsers,
				},
			},
		},
	)

}
