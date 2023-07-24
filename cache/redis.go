package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var (
	mycache = cache.New(&cache.Options{
		Redis: redis.NewRing(&redis.RingOptions{
			Addrs: map[string]string{
				"server": ":6379",
			},
		}),
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
)

type Object struct {
	Str string
	Num int
}

func main() {
	basicUsage()
	advancedUsage()
}

func basicUsage() {
	ctx := context.TODO()
	key := "mykey"
	obj := &Object{
		Str: "mystring",
		Num: 42,
	}
	if err := mycache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: obj,
		TTL:   time.Hour,
	}); err != nil {
		panic(err)
	}
	var wanted Object
	if err := mycache.Get(ctx, key, &wanted); err == nil {
		fmt.Println(wanted) // {mystring 42}
	}
}

func advancedUsage() {
	obj := new(Object)
	err := mycache.Once(&cache.Item{
		Key:   "mykey",
		Value: obj, // destination
		Do: func(*cache.Item) (interface{}, error) {
			return &Object{
				Str: "mystring",
				Num: 42,
			}, nil
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(obj) // &{mystring 42}
}
