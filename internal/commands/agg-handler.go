package commands

import (
	"context"
	"fmt"

	"araj.com/ar/internal/config"
	rssfeed "araj.com/ar/internal/rss_feed"
)

func AggHandler(s *config.State, command config.Command) error {

	url := "https://www.wagslane.dev/index.xml"

	rssFeed, err := rssfeed.FetchFeed(context.Background(), url)

	if err != nil {
		return err
	}

	fmt.Println(rssFeed)

	return nil

}
