package postgress

import (
	"github.com/Genialngash/graphql-go-test/graph/model"
	"github.com/go-pg/pg/v10"
)

type UsersRepo struct {
	DB *pg.DB
}

func (u *UsersRepo) GetUserById(id string) (*model.User, error) {
	var user model.User

	err := u.DB.Model(&user).Where("id =?", id).First()

	if err != nil {
		return nil, err
	}

	return &user, nil
}
