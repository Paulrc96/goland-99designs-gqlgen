package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/paul/go-server/dataloader"
	database "github.com/paul/go-server/graph/db"
	"github.com/paul/go-server/graph/generated"
	"github.com/paul/go-server/graph/model"
)

func (r *mutationResolver) CreateClient(ctx context.Context, client model.ClientInput) (*model.Client, error) {
	sqlStatement := `
		INSERT INTO clients (name, last_name, email, address, birthday, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	id := 0
	var err = database.Db.QueryRow(sqlStatement, &client.Name, &client.LastName, &client.Email, &client.Address, &client.Birthday, &client.CreatedAt).Scan(&id)
	if err != nil {
		// panic(err)
		log.Fatalln("Error creating client", err)
	}

	var response = model.Client{
		ID:        id,
		Name:      client.Name,
		LastName:  client.LastName,
		Email:     client.Email,
		Address:   client.Address,
		Birthday:  client.Birthday,
		CreatedAt: client.CreatedAt,
	}

	return &response, err
}

func (r *postResolver) Comments(ctx context.Context, obj *model.Post) ([]*model.Comment, error) {
	return dataloader.For(ctx).CommentsByPostIds.Load(obj.PostID)
}

func (r *queryResolver) Users(ctx context.Context, first *int) ([]*model.User, error) {
	var limit *int
	limit = first
	if *first == 0 {
		*limit = 10
	}

	allowedColumns := []string{"id", "name", "last_name", "email", "address", "email_verified_at", "password", "remember_token", "created_at", "updated_at"}
	sort.Strings(allowedColumns)

	preloads := GetPreloads(ctx)
	sort.Strings(preloads)
	mytest := func(s string) bool {
		return contains(allowedColumns, s)
	}
	fields := filter(preloads, mytest)
	dbUsers := []*model.User{}
	err := database.Db.Select(&dbUsers, fmt.Sprintf("SELECT %s FROM USERS ORDER BY id asc LIMIT $1", strings.Join(fields, ", ")), *limit)

	if err != nil {
		log.Fatalln(err)
	}

	return dbUsers, err
}

func (r *userResolver) Posts(ctx context.Context, obj *model.User) ([]*model.Post, error) {
	return dataloader.For(ctx).PostsByUserId.Load(obj.ID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
