package route

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"seckill/common/middleware/ratelimit"
	"seckill/storge/redis"
	"time"
)

func init() {
	router := gin.Default()
	var popGroup = router.Group("/pop")
	popGroup.Use(ratelimit.NewBuilder(redis.GetRedisClient(), time.Second, 10).Build())
	popGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
