package main

import (
	"context"
	"os/signal"
	"redis-latency/internal/latency"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	latency.TestLatency(ctx)
}
