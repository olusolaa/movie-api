package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/olusolaa/movieApi/models"
	"log"
	"os"
	"time"
)

var ctx = context.Background()

type redisCache struct {
	host    string
	db      int
	pword   string
	expires time.Duration
}

func NewRedisCache(host string, db int, pword string, exp time.Duration) Cache {
	return &redisCache{
		host:    host,
		db:      db,
		pword:   pword,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	var opts *redis.Options
	if os.Getenv("LOCAL") == "true" {
		redisAddress := fmt.Sprintf("%s:6379", os.Getenv("REDIS_URL"))
		opts = &redis.Options{
			Addr:     redisAddress,
			Password: cache.pword,
			DB:       cache.db,
		}
	} else {
		var err error
		opts, err = redis.ParseURL(os.Getenv("REDIS_URL"))
		if err != nil {
			panic(err)
		}
	}
	return redis.NewClient(opts)
}

func (cache *redisCache) Set(key string, value *[]models.Movie) {
	client := cache.getClient()
	marshal, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	err = client.Set(ctx, key, string(marshal), cache.expires*time.Second).Err()
	if err != nil {
		panic(err)
	}
	return
}

func (cache *redisCache) Get(key string) *[]models.Movie {
	client := cache.getClient()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}
	var movies []models.Movie
	err = json.Unmarshal([]byte(val), &movies)
	if err != nil {
		panic(err)
	}
	log.Println("Movies retrieved from cache")
	return &movies
}

func (cache *redisCache) SetCharacters(key string, value []models.Character) {
	client := cache.getClient()
	marshal, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	err = client.Set(ctx, key, string(marshal), cache.expires*time.Second).Err()
	if err != nil {
		panic(err)
	}
	log.Println("Added characters to cache")
	return
}

func (cache *redisCache) GetCharacters(key string) []models.Character {
	client := cache.getClient()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}
	var characters []models.Character
	err = json.Unmarshal([]byte(val), &characters)
	if err != nil {
		panic(err)
	}
	log.Println("Characters retrieved from cache")
	return characters
}
