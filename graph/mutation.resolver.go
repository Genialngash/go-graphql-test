package graph

import (
	"context"
	"errors"

	"github.com/Genialngash/graphql-go-test/graph/model"
)

type mutationResolver struct{ *Resolver }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// CreateMeetUp is the resolver for the createMeetUp field.
func (r *mutationResolver) CreateMeetUp(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough ")

	}
	if len(input.Description) < 3 {
		return nil, errors.New("description not long enough")
	}
	meetup := &model.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserId:      "5",
	}

	return r.MeetupRepo.CreateMeetUp(meetup)

}
