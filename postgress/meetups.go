package postgress

import (
	"fmt"

	"github.com/Genialngash/graphql-go-test/graph/model"
	"github.com/go-pg/pg/v10"
)

type MeetupsRepo struct {
	DB *pg.DB
}

func (m *MeetupsRepo) GetMeetups() ([]*model.Meetup, error) {
	var meetups []*model.Meetup
	err := m.DB.Model(&meetups).Select()
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	return meetups, nil
}

func (m *MeetupsRepo) CreateMeetUp(meetup *model.Meetup) (*model.Meetup, error) {
	_, err := m.DB.Model(meetup).Returning("*").Insert()

	return meetup, err
}
