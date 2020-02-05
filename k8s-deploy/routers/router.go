package routers

import (
	"fmt"
	_ "github.com/DevOpserzhao/ops_gin/k8s-deploy/docs"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/middleware/jwt"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/pkg/logging"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/pkg/setting"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/routers/api"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/routers/api/v1/k8s"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"io"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	gin.DisableConsoleColor()

	gin.DefaultWriter = io.MultiWriter(logging.F_A)
	//r.Use(gin.Logger())
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - %s \"%s %s %s %d %s \"%s\" %s\"\n",
			param.TimeStamp.Format("2006-01-02 03:04:05"),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/auth", api.GetAuth)
	//r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())

	{

		apiv1.GET("/k8s/Namespaces", k8s.GetNamespacesall)
		apiv1.GET("/k8s/Namespaces/:id", k8s.GetNamespaces)
		apiv1.GET("/k8s/Deployment", k8s.Deployments)
		apiv1.GET("/k8s/Deployment/:id", k8s.Deployment)
		apiv1.PUT("/k8s/Deployment/:id", k8s.SetDeployment)

	}

	return r
}
