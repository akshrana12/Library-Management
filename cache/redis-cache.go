package cache

import (
	"encoding/json"
	"time"

	models "github.com/akshrana12/Library-Management/model"
	"github.com/go-redis/redis/v7"
)

type RedisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) *RedisCache {
	return &RedisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *RedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       cache.db,
	})
}

func (cache *RedisCache) SetBook(key string, post *models.Books) {
	client := cache.getClient()
	json, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}
	client.Set(key, json, cache.expires*time.Second)
}

func (cache *RedisCache) GetBook(key string) *models.Books {
	client := cache.getClient()
	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	book := models.Books{}
	err = json.Unmarshal([]byte(val), &book)
	if err != nil {
		panic(err)
	}
	return &book
}

func (cache *RedisCache) SetBookByArray(key string, post []models.Books) {
	client := cache.getClient()
	json, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}
	client.Set(key, json, cache.expires*time.Second)
}

func (cache *RedisCache) GetBookArray(key string) []models.Books {
	client := cache.getClient()
	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}
	book := []models.Books{}
	err = json.Unmarshal([]byte(val), &book)
	if err != nil {
		panic(err)
	}
	return book
}
