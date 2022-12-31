package graph

import (
	"context"

	"github.com/Genialngash/graphql-go-test/graph/model"
)

type userResolver struct{ *Resolver }

func (r *Resolver) User() UserResolver { return &userResolver{r} }

// Meetups is the resolver for the meetups field.
func (r *userResolver) Meetups(ctx context.Context, obj *model.User) ([]*model.Meetup, error) {
	var meetups []*model.Meetup

	for _, m := range meetups {
		if m.UserId == obj.ID {
			meetups = append(meetups, m)

		}
	}
	return meetups, nil
}
