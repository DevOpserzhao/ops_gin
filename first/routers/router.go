package routers

import (
	_ "github.com/DevOpserzhao/ops_gin/first/docs"
	"github.com/DevOpserzhao/ops_gin/first/middleware/jwt"
	"github.com/DevOpserzhao/ops_gin/first/pkg/setting"
	"github.com/DevOpserzhao/ops_gin/first/pkg/upload"
	"github.com/DevOpserzhao/ops_gin/first/routers/api"
	"github.com/DevOpserzhao/ops_gin/first/routers/api/k8s"
	v1 "github.com/DevOpserzhao/ops_gin/first/routers/api/v1"
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
	r.POST("/upload", api.UploadImage)
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())

	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

		apiv1.GET("/k8s/Namespaces", k8s.GetNamespacesall)
		apiv1.GET("/k8s/Namespaces/:id", k8s.GetNamespaces)
		apiv1.GET("/k8s/Deployment", k8s.Deployments)
		apiv1.GET("/k8s/Deployment/:id", k8s.Deployment)
		apiv1.PUT("/k8s/Deployment/:id", k8s.SetDeployment)

	}

	return r
}
