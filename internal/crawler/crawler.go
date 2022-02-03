package crawler

import (
	"context"
	"time"
)

func Run(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	default:
		time.Sleep(time.Minute)
	}
}
