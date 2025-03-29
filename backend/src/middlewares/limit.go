package middlewares

import (
	"sync"

	"golang.org/x/time/rate"
)

var limiters = sync.Map{}

func getLimiter(ip string) *rate.Limiter {
	limiter, ok := limiters.Load(ip)

	if !ok {
		newLimiter := rate.NewLimiter(1, 5)
		limiters.Store(ip, newLimiter)
		return newLimiter
	}
	return limiter.(*rate.Limiter)
}
