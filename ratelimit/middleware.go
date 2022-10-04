package ratelimit

import (
	"time"

	"github.com/gin-gonic/gin"
	limiter "github.com/ulule/limiter/v3"
	_gin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimit(period time.Duration, limit int64) gin.HandlerFunc {
	rate := limiter.Rate{Period: period, Limit: limit}
	limiterStore := memory.NewStore()
	return _gin.NewMiddleware(limiter.New(limiterStore, rate))
}
