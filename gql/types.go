package gql

import "github.com/graphql-go/graphql"

// User describes a graphql object containing a User
var User = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"userid": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"age": &graphql.Field{
				Type: graphql.Int,
			},
			"profession": &graphql.Field{
				Type: graphql.String,
			},
			"friendly": &graphql.Field{
				Type: graphql.Boolean,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
