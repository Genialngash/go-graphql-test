package graph

import (
	"context"

	"github.com/Genialngash/graphql-go-test/graph/model"
)

// Meetups is the resolver for the meetups field.
func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {

	return r.MeetupRepo.GetMeetups()
}

// Meetup returns MeetupResolver implementation.

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.

type queryResolver struct{ *Resolver }
 