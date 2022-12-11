package redisConnection

import (
	"context"
	// "fmt"
	"errors"

	"github.com/go-redis/redis/v9"
)

type DBconnector struct {
	rdb redis.Client
}

func (m *DBconnector) Setting() {
	m.rdb = *redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (m DBconnector) SetHash(ctx context.Context, key string, value string) {
	m.rdb.HSet(ctx, "email", key, value).Result()
}

func (m DBconnector) GetHash(ctx context.Context, key string) (string, error) {
	value, err := m.rdb.HGet(ctx, "email", key).Result()
	if err != nil {
		return "", errors.New("no value")
	}
	return value, nil
}

// func ExampleClient() {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})

// 	err := rdb.Set(ctx, "key", "value", 0).Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	val, err := rdb.Get(ctx, "key").Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("key", val)

// 	val2, err := rdb.Get(ctx, "key2").Result()
// 	if err == redis.Nil {
// 		fmt.Println("key2 does not exist")
// 	} else if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("key2", val2)
// 	}
// 	// Output: key value
// 	// key2 does not exist
// }
