package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gql_serv/graph/generated"
	"gql_serv/graph/model"
)

func (r *mutationResolver) CreateLink(ctx context.Context, newlink model.NewLink) (*model.Link, error) {
	var link model.Link
	link.Address = newlink.Address
	link.Title = newlink.Title
	link.ID = "1111"
	link.Author = &model.User{
		ID:   "1",
		Name: "test",
	}
	return &link, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, user model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, login model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var dummyLinks []*model.Link
	link := model.Link{
		Title:   "It`t dummy link",
		Address: "http://dummy.link",
		Author:  &model.User{Name: "user"},
	}
	dummyLinks = append(dummyLinks, &link)
	return dummyLinks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
