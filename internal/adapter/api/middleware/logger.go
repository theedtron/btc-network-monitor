package middleware

import (
	"btc-network-monitor/internal/helper"
	"btc-network-monitor/internal/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		log := logger.NewLogrusLogger()
		start := time.Now()
		fields := logger.Fields{
			"name":       "user-service",
			"status":     c.Writer.Status(),
			"path":       c.Request.URL.Path,
			"method":     c.Request.Method,
			"ip":         c.ClientIP(),
			"latency":    time.Since(start).Milliseconds(),
			"user-agent": c.Request.UserAgent(),
		}
		header := c.Request.Header
		if val, ok := header["Authorization"]; ok {
			token := strings.Split(val[0], " ")[1]
			claims, err := helper.ValidateJWT(token)
			if err != nil {
				return
			}
			if user, ok := claims["data"].(map[string]interface{}); ok {
				fields["user"] = user["id"]
				fields["user-type"] = user["user_type"]
			} else {
				logger.Error("Type assertion failed.")
			}
		}

		c.Next()
		log.WithFields(fields).Infof("request handled")
	}
}
