package commands

import (
	"context"
	"fmt"

	"araj.com/ar/internal/config"
)

func FeedsHandler(s *config.State, command config.Command) error {
	feeds, err := s.Db.GetAllFeeds(context.Background())

	if err != nil {
		return err
	}

	for _, f := range feeds {
		fmt.Printf("name:= %s\n", f.Name.String)
		fmt.Printf("url:= %s\n", f.Url.String)
		userThatCreatedFeed, err := s.Db.GetUseById(context.Background(), f.UserID.UUID)

		if err != nil {
			return err
		}

		fmt.Printf("username:= %s\n", userThatCreatedFeed.Name)
	}

	return nil
}
