package postgress

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type DbLogger struct{}

func (d DbLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	val, _ := q.FormattedQuery()
	fmt.Printf(string(val))
	return ctx, nil

}

func (d DbLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {

	val, _ := q.FormattedQuery()
	fmt.Printf(string(val))

	fmt.Sprintln(val)
	return nil

}

func New(opt *pg.Options) *pg.DB {
	// url := fmt.Sprintf("postgres://ngash:login@localhost:5432/meetup_dev?sslmode=disable")

	// optUrl, err := pg.ParseURL(url)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil
	// }

	// if err != nil {
	// 	fmt.Printf(err.Error())
	// }

	return pg.Connect(opt)
}
