/*
*

	@author:ExiaYe
	@data:2024/4/21
	@note: 限流中间件

*
*/
package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

/**
 * @Author ExiaYe
 * @Description //TODO 限流中间件
 * @Date 11:56 2024/4/21
 **/ // 创建指定填充速率和容量大小的令牌桶
func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就中断本次请求返回 rate limit...
		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		// 取到令牌就放行
		c.Next()
	}
}
