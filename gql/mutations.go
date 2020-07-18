package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/theShivaa/go-gql-crud/postgres"
)

// NewRoot returns base query type. This is where we add all the base queries
func RootMutations(db *postgres.Db) *graphql.Object {
	// Create a resolver holding our databse. Resolver can be found in resolvers.go
	resolver := Resolver{db: db}

	// Create a new Root that describes our base query set up. In this
	// example we have a user query that takes one argument called name
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Type: User,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"profession": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"friendly": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.CreateUser,
			},
			"update": &graphql.Field{
				Type: User,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.UpdateUser,
			},
			"delete": &graphql.Field{
				Type: User,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.DeleteUser,
			},
		},
	})

}
