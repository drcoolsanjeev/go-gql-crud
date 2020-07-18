package gql

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/theShivaa/go-gql-crud/postgres"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	db *postgres.Db
}

// GetUserByName resolves our user query through a db call to GetUserByName
func (r *Resolver) GetUserByName(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	name, ok := p.Args["name"].(string)
	if ok {
		users := r.db.GetUsersByName(name)
		return users, nil
	}

	return nil, nil
}

func (r *Resolver) GetAllUsers(p graphql.ResolveParams) (interface{}, error) {
	return r.db.GetAllUsers(), nil
}

func (r *Resolver) GetUserByUserID(p graphql.ResolveParams) (interface{}, error) {
	userid, ok := p.Args["userid"].(string)
	if ok {
		user := r.db.GetUsersByUserID(userid)
		return user, nil
	}
	return nil, nil
}

func (r *Resolver) CreateUser(p graphql.ResolveParams) (interface{}, error) {
	userid, ok := p.Args["userid"].(string)
	if !ok {
		return nil, errors.New("userid is mandatory")
	}
	name, ok := p.Args["name"].(string)
	if !ok {
		return nil, errors.New("name is mandatory")
	}
	age, ok := p.Args["age"].(int)
	if !ok {
		return nil, errors.New("age is mandatory")
	}
	profession, ok := p.Args["profession"].(string)
	if !ok {
		return nil, errors.New("profession is mandatory")
	}
	friendly, ok := p.Args["friendly"].(bool)
	if !ok {
		return nil, errors.New("friendly is mandatory")
	}
	password, ok := p.Args["password"].(string)
	if !ok {
		return nil, errors.New("password is mandatory")
	}

	user := postgres.User{
		UserID:     userid,
		Name:       name,
		Age:        age,
		Profession: profession,
		Friendly:   friendly,
		Password:   password,
	}
	return r.db.CreateUser(user), nil

}

func (r *Resolver) UpdateUser(p graphql.ResolveParams) (interface{}, error) {
	userid, useridOk := p.Args["userid"].(string)
	name, nameOk := p.Args["name"].(string)

	if useridOk {
		user := postgres.User{UserID: userid}
		if nameOk {
			user.Name = name
		}

		return r.db.UpdateUserName(user), nil
	}
	return nil, nil
}

func (r *Resolver) DeleteUser(p graphql.ResolveParams) (interface{}, error) {
	userid, ok := p.Args["userid"].(string)
	if ok {
		return r.db.DeleteUser(userid), nil
	}
	return nil, nil
}
