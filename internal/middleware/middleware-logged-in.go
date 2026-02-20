package middleware

import (
	"context"

	"araj.com/ar/internal/config"
	"araj.com/ar/internal/database"
)

func MiddlewareLoggedIn(handler func(s *config.State, cmd config.Command, user database.User) error) func(*config.State, config.Command) error {

	return func(s *config.State, c config.Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)

		if err != nil {
			return err
		}

		handler(s, c, user)
		return nil
	}
}
