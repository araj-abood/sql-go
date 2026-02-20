package commands

import (
	"context"
	"fmt"

	"araj.com/ar/internal/config"
	"araj.com/ar/internal/database"
)

func FollowingHandler(s *config.State, command config.Command, user database.User) error {

	feedsFollowed, err := s.Db.GetFeedsThatUserFollows(context.Background(), user.ID)

	if err != nil {
		return err
	}

	for _, feed := range feedsFollowed {
		fmt.Printf("feed name %s", feed.FeedName.String)
	}
	return nil
}
