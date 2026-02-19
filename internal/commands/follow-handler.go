package commands

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"araj.com/ar/internal/config"
	"araj.com/ar/internal/database"
	"araj.com/ar/internal/helper"
	"github.com/google/uuid"
)

func FollowHandler(s *config.State, command config.Command) error {

	if command.IsArgsEmpty(0) {
		return fmt.Errorf("Please enter the feed url")
	}

	url := command.Arguments[0]

	feed, err := s.Db.GetFeedByUr(context.Background(), sql.NullString{String: url, Valid: true})

	if err != nil {
		return fmt.Errorf("Could not find feed")
	}

	user, err := helper.GetCurrentUserFromDb(s)

	if err != nil {
		return err
	}

	_, err = s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	fmt.Printf("feed: %s\n", feed.Name.String)
	fmt.Printf("user name: %s\n", user.Name)

	return nil
}
