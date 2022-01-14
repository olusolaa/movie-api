package main

import (
	"github.com/olusolaa/movieApi/cache"
	"github.com/olusolaa/movieApi/db"
	_ "github.com/olusolaa/movieApi/docs"
	"github.com/olusolaa/movieApi/server"
)

func main() {

	var moviesCache = cache.NewRedisCache("localhost:6379", 1, "", 100)

	DB := &db.PostgresDB{}
	DB.Init()

	s := &server.Server{
		DB:    DB,
		Cache: moviesCache,
	}
	s.Start()
}
