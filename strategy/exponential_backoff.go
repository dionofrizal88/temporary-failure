package strategy

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RdbConf struct {
	Host     string
	Port     string
	Username string
	Password string
}

func (r RdbConf) ConnectToRedisDB() (*redis.Client, error) {
	ctx := context.Background()
	var rdb *redis.Client

	conn := func() error {
		log.Printf("Retry connect rdb")
		rdb = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", r.Host, r.Port),
			Username: r.Username,
			Password: r.Password,
		})

		_, err := rdb.Ping(ctx).Result()
		if err != nil {
			return fmt.Errorf("Could not connect to redis %v", err)
		}

		log.Println("Success connect to Redis")

		return nil
	}

	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.MaxElapsedTime = 10 * time.Second

	err := backoff.Retry(conn, expBackoff)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after retrying: %v", err)
	}

	return rdb, nil
}
