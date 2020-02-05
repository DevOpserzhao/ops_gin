package routers

import (
	_ "github.com/DevOpserzhao/ops_gin/first/docs"
	"github.com/DevOpserzhao/ops_gin/first/middleware/jwt"
	"github.com/DevOpserzhao/ops_gin/first/pkg/setting"
	"github.com/DevOpserzhao/ops_gin/first/pkg/upload"
	"github.com/DevOpserzhao/ops_gin/first/routers/api"
	"github.com/DevOpserzhao/ops_gin/first/routers/api/v1/k8s"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/auth", api.GetAuth)
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

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
