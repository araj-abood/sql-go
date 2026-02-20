package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"araj.com/ar/internal/config"
	"araj.com/ar/internal/database"
	"github.com/google/uuid"
)

func AddFeedHandler(s *config.State, command config.Command) error {
	if command.IsArgsEmpty(2) {
		return errors.New("Expecting name and url arguments please supply them")
	}

	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)

	if err != nil {
		return err
	}

	feed, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      sql.NullString{String: command.Arguments[0], Valid: true},
		Url:       sql.NullString{String: command.Arguments[1], Valid: true},
		UserID:    uuid.NullUUID{UUID: user.ID, Valid: true},
	})

	s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return err
	}

	fmt.Printf("feed_name: %s\nc", feed.Name.String)

	return nil
}
