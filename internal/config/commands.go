package config

import (
	"errors"
)

type Commands struct {
	Methods map[string]func(state *State, command Command) error
}

func (c *Commands) Run(state *State, command Command) error {
	cmd, ok := c.Methods[command.Name]

	if !ok {
		return errors.New("Command does not exist")
	}

	return cmd(state, command)
}

func (c *Commands) Register(name string, f func(s *State, command Command) error) error {
	c.Methods[name] = f
	return nil
}
