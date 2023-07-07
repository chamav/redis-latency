package latency

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
	"time"
)

var clients = make(map[string]*redis.Client)

func TestLatency(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	redisAddress := r.URL.Query().Get("url")
	if redisAddress == "" {
		redisAddress = "localhost:6379"
	}
	_, ok := clients[redisAddress]
	if ok == false {
		clients[redisAddress] = redis.NewClient(&redis.Options{
			Addr:     redisAddress,
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}
	now := time.Now()
	_, err := clients[redisAddress].Ping(ctx).Result()
	elapsed := time.Since(now)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprint(w, "redis_latency ")
	fmt.Fprintf(w, fmt.Sprintf("%.0f", float64(elapsed.Microseconds())))
	fmt.Fprint(w, " ")
	fmt.Fprint(w, strconv.FormatInt(now.Unix(), 10))
}
