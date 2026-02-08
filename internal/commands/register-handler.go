package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"araj.com/ar/internal/config"
	"araj.com/ar/internal/database"
	"github.com/google/uuid"
)

func RegisterHandler(s *config.State, command config.Command) error {
	if len(command.Arguments) <= 0 {
		return errors.New("Expecting username argument")
	}

	user, err := s.Db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: command.Arguments[0]})

	if err != nil {
		return err
	}

	s.Cfg.SetUser(user.Name)

	fmt.Println("User created successfully")
	return nil
}
