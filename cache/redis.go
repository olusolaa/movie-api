package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/olusolaa/movieApi/models"
	"log"
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
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		DB:       cache.db,
		Password: cache.pword,
	})
}

// add stack trace dependencies injection

func (cache *redisCache) Set(key string, value *[]models.Movie) {
	client := cache.getClient()
	marshal, err := json.Marshal(value)
	if err != nil {
		panic(any(err))
	}
	err = client.Set(ctx, key, string(marshal), cache.expires*time.Second).Err()
	if err != nil {
		panic(any(err))
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
		panic(any(err))
	}
	log.Println("Movies retrieved from cache")
	return &movies
}
