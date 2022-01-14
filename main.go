package main

import (
	"github.com/olusolaa/movieApi/db"
	_ "github.com/olusolaa/movieApi/docs"
	"github.com/olusolaa/movieApi/server"
)

func main() {

	DB := &db.PostgresDB{}
	DB.Init()

	s := &server.Server{
		DB: DB,
	}
	s.Start()
}
