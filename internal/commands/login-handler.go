package commands

import (
	"context"
	"errors"
	"fmt"

	"araj.com/ar/internal/config"
)

func LoginHandler(s *config.State, command config.Command) error {
	if len(command.Arguments) <= 0 {
		return errors.New("the login handler expects a single argument, the username.")
	}

	username := command.Arguments[0]

	_, err := s.Db.GetUser(context.Background(), username)

	if err != nil {
		return err
	}

	err = s.Cfg.SetUser(username)

	if err != nil {
		return err
	}
	fmt.Println("The user has been set")
	return nil

}
