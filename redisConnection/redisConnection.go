package redisConnection

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v9"
)

type DBconnector struct {
	rdb redis.Client
}

type DBResult struct {
	Id   int32
	Name string
}

func (m *DBconnector) init() {
	m.rdb = *redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (m DBconnector) SetHash(ctx context.Context, key string, name string, id int32) {
	m.rdb.HSet(ctx, key, "id", id, "name", name).Result()
}

func (m DBconnector) GetHash(ctx context.Context, key string) (*DBResult, error) {
	id, err := m.rdb.HGet(ctx, "id", key).Result()
	if err != nil {
		return nil, errors.New("no value")
	}
	name, err := m.rdb.HGet(ctx, "name", key).Result()
	if err != nil {
		return nil, errors.New("no value")
	}
	var parseId int32
	fmt.Sscan(id, &parseId)
	return &DBResult{
		Id:   parseId,
		Name: name,
	}, nil
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
