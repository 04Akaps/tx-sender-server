package redis

import (
	"encoding/json"
	"github.com/04Akaps/tx-sender-server/config"
	"github.com/go-redis/redis/v7"
	"log"
	"time"
)

type RedisClient struct {
	config *config.Config
	client *redis.Client
}

type RedisImpl interface {
	Save(key string, data interface{}, time time.Duration) error
	Get(key string) (interface{}, error)
}

func NewRedis(config *config.Config) (RedisImpl, error) {
	r := &RedisClient{config: config}

	r.client = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Url,
		Password: config.Redis.Password,
		Username: config.Redis.User,
		DB:       config.Redis.DB,
	})

	if _, err := r.client.Ping().Result(); err != nil {
		return nil, err
	} else {
		log.Println("Success To Connect Redis")
		return r, nil
	}
}

func (r *RedisClient) Save(key string, data interface{}, time time.Duration) error {
	_, err := r.client.Set(key, data, time).Result()
	return err
}

func (r *RedisClient) Get(key string) (interface{}, error) {
	if d, err := r.client.Get(key).Bytes(); err != nil {
		return nil, err
	} else {
		var res interface{}

		if err = json.Unmarshal(d, &res); err != nil {
			return nil, err
		} else {
			return res, nil
		}
	}
}
