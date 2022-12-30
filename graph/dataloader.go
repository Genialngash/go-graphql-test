package graph

import (
	"context"
	"net/http"
	"time"

	"github.com/Genialngash/graphql-go-test/graph/model"
	"github.com/go-pg/pg/v10"
)

const userloaderKey = "userloader"

func DataLoaderMiddleWare(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*model.User, []error) {
				var users []*model.User

				if err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select(); err != nil {
					return nil, []error{err}
				}
				u := make(map[string]*model.User, len(users))

				for _, user := range users {
					u[user.ID] = user
				}
				result := make([]*model.User, len(ids))

				for i, id := range ids {
					result[i] = u[id]
				}

				return users, nil

			},
		}
		ctx := context.WithValue(r.Context(), userloaderKey, &userLoader)
		next.ServeHTTP(w, r.WithContext(ctx))

	})

}

func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userloaderKey).(*UserLoader)
}
