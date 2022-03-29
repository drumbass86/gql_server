package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gql_serv/db"
	"gql_serv/graph/generated"
	"gql_serv/graph/model"
	"strconv"
)

func (r *mutationResolver) CreateLink(ctx context.Context, newlink model.NewLink) (*model.Link, error) {
	var link db.Link
	link.Address = newlink.Address
	link.Title = newlink.Title
	//!TODO userid on current user
	link.UserID = 1
	_, err := db.CreateLink(&link)
	if err == nil {
		return &model.Link{
			ID:      strconv.FormatUint(uint64(link.ID), 10),
			Title:   link.Title,
			Address: link.Address,
			Author: &model.User{
				ID:   strconv.FormatUint(uint64(link.UserID), 10),
				Name: "default",
			},
		}, nil
	} else {
		return &model.Link{
			ID:      "-1",
			Title:   link.Title,
			Address: link.Address,
			Author: &model.User{
				ID:   strconv.FormatUint(uint64(link.UserID), 10),
				Name: "default",
			},
		}, err
	}

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
	dbLinks, err := db.GetAllFullLinks()
	if err == nil {
		var links []*model.Link
		for _, l := range dbLinks {
			link := model.Link{
				Title:   l.Title,
				Address: l.Address,
				Author: &model.User{
					ID:   strconv.FormatUint(uint64(l.User_.ID), 10),
					Name: l.User_.Username,
				},
			}
			links = append(links, &link)
		}
		return links, nil
	} else {
		return nil, err
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
