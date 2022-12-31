package postgress

import (
	"fmt"

	"github.com/Genialngash/graphql-go-test/graph/model"
	"github.com/go-pg/pg/v10"
)

type MeetupsRepo struct {
	DB *pg.DB
}

func (m *MeetupsRepo) GetMeetups(filter *model.MeetUpFilter, limit, offset *int) ([]*model.Meetup, error) {
	var meetups []*model.Meetup
	query := m.DB.Model(&meetups).Order("id")
	if filter != nil {
		if filter.Name != "" {
			query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", filter.Name))

		}
	}

	if limit != nil {
		query.Limit(*limit)

	}
	if offset != nil {
		query.Offset(*offset)

	}
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
