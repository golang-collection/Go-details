package routers

import (
	"Go-details/global"
	"Go-details/internal/routers/api/sd"
	"Go-details/pkg/limiter"
	"github.com/gin-gonic/gin"
	"time"

	//_ "Go-details/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

/**
* @Author: super
* @Date: 2020-10-07 14:43
* @Description:
**/

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		//r.Use(middleware.AccessLog())
		//r.Use(middleware.Recovery())
	}

	//r.Use(middleware.Tracing())
	//r.Use(middleware.RateLimiter(methodLimiters))
	//r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	//r.Use(middleware.Translations())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// The health check handlers
	svcd := r.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return r
}