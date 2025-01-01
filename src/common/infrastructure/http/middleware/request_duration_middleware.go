package commonMiddlewares

import (
	"github.com/gin-gonic/gin"
	"gorm-ddd-example/src/common/infrastructure/http/metrics"
	"time"
)

func RequestDurationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.FullPath()
		metrics.RequestDuration.WithLabelValues(method, path, string(rune(statusCode))).Observe(duration)
	}
}
