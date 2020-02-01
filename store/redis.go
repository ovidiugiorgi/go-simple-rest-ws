package store

import (
	redisv7 "github.com/go-redis/redis/v7"
)

type Redis struct {
	client *redisv7.Client
}

func (r *Redis) Client() *redisv7.Client {
	if r.client == nil {
		r.client = redisv7.NewClient(&redisv7.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		_, err := r.client.Ping().Result()
		if err != nil {
			panic(err.Error())
		}
	}
	return r.client
}
