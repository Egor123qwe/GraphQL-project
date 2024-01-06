package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"

	"github.com/Egor123qwe/GraphQL-project/graph/model"
)

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id *string) (*model.User, error) {
	var Id = ""
	if id != nil {
		Id = *id
	}

	usr, err := r.UserList.GetUser(ctx, Id)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:         usr.Id,
		Email:      usr.Email,
		Name:       &usr.Name,
		Age:        int(usr.Age),
		University: usr.University,
		Hobbies:    usr.Hobbies,
		Course:     int(usr.Course),
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
