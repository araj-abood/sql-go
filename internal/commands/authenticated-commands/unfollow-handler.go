package authenticatedcommands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"araj.com/ar/internal/config"
	"araj.com/ar/internal/database"
)

func UnfollowHandler(s *config.State, command config.Command, user database.User) error {
	if command.IsArgsEmpty(1) {
		return errors.New("Expecing feed url")
	}
	url := command.Arguments[0]

	feed, err := s.Db.GetFeedByUr(context.Background(), sql.NullString{Valid: true, String: url})

	if err != nil {
		return err
	}

	err = s.Db.DeleteFollow(context.Background(), database.DeleteFollowParams{UserID: user.ID, FeedID: feed.ID})

	if err != nil {
		return err
	}

	fmt.Printf("Deletee follow\n")

	return nil
}
