package commands

import (
	"context"
	"fmt"

	"araj.com/ar/internal/config"
)

func UsersHandler(s *config.State, command config.Command) error {
	users, err := s.Db.GetAllUsers(context.Background())

	if err != nil {
		return err
	}

	for _, user := range users {
		appendedText := ""
		isSameUser := user.Name == s.Cfg.CurrentUserName

		if isSameUser {
			appendedText = "(current)"
		}

		fmt.Printf("* %s %s", user.Name, appendedText)
	}

	return nil
}
