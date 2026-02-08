package main

import (
	"database/sql"
	"log"
	"os"

	"araj.com/ar/internal/commands"
	"araj.com/ar/internal/config"
	"araj.com/ar/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	cfg, err := config.Read()

	if err != nil {
		log.Fatal(err)
	}

	state := &config.State{
		Cfg: &cfg,
	}

	db, err := sql.Open("postgres", state.Cfg.DBUrl)

	if err != nil {
		log.Fatal(err)

	}

	dbQueries := database.New(db)

	state.Db = dbQueries

	cmds := config.Commands{
		Methods: make(map[string]func(state *config.State, command config.Command) error),
	}

	cmds.Register("login", commands.LoginHandler)
	cmds.Register("register", commands.RegisterHandler)
	cmds.Register("reset", commands.ResetHandler)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.Run(state, config.Command{Name: cmdName, Arguments: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
