package commands

import (
	"context"
	"fmt"

	"araj.com/ar/internal/config"
	"araj.com/ar/internal/helper"
)

func FollowingHandler(s *config.State, command config.Command) error {
	user, err := helper.GetCurrentUserFromDb(s)
	if err != nil {
		return fmt.Errorf("Could not find current user")
	}

	feedsFollowed, err := s.Db.GetFeedsThatUserFollows(context.Background(), user.ID)

	for _, feed := range feedsFollowed {
		fmt.Printf("feed name %s", feed.FeedName.String)
	}
	return nil
}
