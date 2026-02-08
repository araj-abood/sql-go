package commands

import (
	"context"

	"araj.com/ar/internal/config"
)

func ResetHandler(s *config.State, command config.Command) error {
	err := s.Db.DeleteAllUsers(context.Background())

	if err != nil {
		return err
	}

	return nil
}
