package graph

import (
	"context"

	"github.com/Genialngash/graphql-go-test/graph/model"
)

type meetupResolver struct{ *Resolver }

func (r *Resolver) Meetup() MeetupResolver { return &meetupResolver{r} }

// User is the resolver for the user field.
func (r *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {

	return r.UsersRepo.GetUserById(obj.UserId)
}
