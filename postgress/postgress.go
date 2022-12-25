package postgress

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type DbLogger struct {}

func (d DbLogger)BeforeQuery(ctx context.Context, q *pg.QueryEvent)(context.Context,error){
	return ctx,nil

}

func (d DbLogger)AfterQuery(ctx context.Context, q *pg.QueryEvent)(error){
	fmt.Println(q.FormattedQuery())
	return nil
	
}

func New(opt *pg.Options) *pg.DB{
	return pg.Connect(opt)
}