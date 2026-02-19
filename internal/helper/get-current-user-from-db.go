package helper

import (
	"context"
	"fmt"

	"araj.com/ar/internal/config"
	"araj.com/ar/internal/database"
)

func GetCurrentUserFromDb(s *config.State) (database.User, error) {
	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)

	if err != nil {
		return database.User{}, fmt.Errorf("Could not find user")
	}

	return user, nil
}
