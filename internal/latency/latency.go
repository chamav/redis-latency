package latency

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

func TestLatency(ctx context.Context) {
	redisAdd := ""
	if len(os.Args) > 1 {
		redisAdd = os.Args[1]
	} else {
		redisAdd = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAdd,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	now := time.Now()
	_, err := rdb.Ping(ctx).Result()
	elapsed := time.Since(now)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	//fmt.Println("Ping latency: ", float64(elapsed.Microseconds()))
	fmt.Println(float64(elapsed.Microseconds()))

}
