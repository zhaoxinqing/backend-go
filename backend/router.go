package backend

import (
	"demo/public"
	"demo/public/logger"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(addr string) {

	gin.SetMode(gin.ReleaseMode) // DebugMode、ReleaseMode、TestMode

	engine := gin.Default()

	engine.SetTrustedProxies(nil) // 代理

	routers(engine) // 路由注册

	logger := logger.NewMyLogger()

	logger.Info("ALL MISSION SUCCESS !!!")

	logger.Info(fmt.Sprintf("%s, 服务已启动.", addr))

	engine.Run(addr)
}

func routers(app *gin.Engine) {

	api := app.Group("/api")

	api.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 10000, "data": "ok", "message": time.Now().Format(public.TIME_FORMAT)})
	})
}
