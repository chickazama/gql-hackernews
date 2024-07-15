package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/chickazama/gql-hackernews/graph/model"
	"github.com/chickazama/gql-hackernews/internal/database"
	"github.com/chickazama/gql-hackernews/internal/models"
)

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	link := models.NewLink(input.Title, input.Address, "test")
	_, err := link.Save(database.DB)
	if err != nil {
		return nil, err
	}
	return &model.Link{ID: link.ID, Title: link.Title, Address: link.Address, User: &model.User{ID: link.UserID, Name: "Test"}}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// Links is the resolver for the links field.
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	// var ret []*model.Link
	// dummy := &model.Link{
	// 	Title:   "Dummy Link",
	// 	Address: "https://example.com",
	// 	User:    &model.User{Name: "Matty"},
	// }
	// ret = append(ret, dummy)
	// return ret, nil
	links, err := database.GetAllLinks()
	if err != nil {
		return nil, err
	}
	var ret []*model.Link
	for _, link := range links {
		next := &model.Link{ID: link.ID, Title: link.Title, Address: link.Address, User: &model.User{ID: link.UserID, Name: "Test"}}
		ret = append(ret, next)
	}
	return ret, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
