package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/generated"
	"github.com/jivison/gowon-indexer/lib/graph/model"
	"github.com/jivison/gowon-indexer/lib/services/lastfm"
	"github.com/jivison/gowon-indexer/lib/tasks"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var dbUsers []db.User
	var resultUsers []*model.User

	err := db.Db.Model(&dbUsers).Select()

	if err != nil {
		log.Fatal(err)
	}

	for _, user := range dbUsers {
		resultUsers = append(resultUsers, &model.User{
			ID:             int(user.ID),
			LastFMUsername: user.LastFMUsername,
		})
	}

	return resultUsers, nil
}

func (r *queryResolver) GetUser(ctx context.Context, username string) (*model.User, error) {
	dbUser := new(db.User)

	err := db.Db.Model(dbUser).Where("last_fm_username = ?", username).Limit(1).Select()

	if err != nil {
		lastFM := lastfm.CreateService()

		if !lastFM.ValidateUser(username) {
			return nil, gqlerror.Errorf("The user %s doesn't exist in Last.fm!", username)
		}

		dbUser = &db.User{
			LastFMUsername: username,
		}

		db.Db.Model(dbUser).Insert()
	}

	resultUser := &model.User{
		ID:             int(dbUser.ID),
		LastFMUsername: dbUser.LastFMUsername,
	}

	return resultUser, nil
}

func (r *queryResolver) CoolQuery(ctx context.Context, str string) (*string, error) {
	tasks.TaskServer.SendTestTask(str)

	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
