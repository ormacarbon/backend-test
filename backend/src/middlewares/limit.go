package middlewares

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
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

func RateLimiter(ctx *gin.Context) {
	ip := ctx.ClientIP()
	limiter := getLimiter(ip)

	if !limiter.Allow() {
		ctx.JSON(
			http.StatusTooManyRequests,
			gin.H{"error": "too many requests"},
		)
		ctx.Abort()
		return
	}
	ctx.Next()
}
