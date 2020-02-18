package api

import (
	"github.com/DevOpserzhao/ops_gin/first/pkg/app"
	"github.com/DevOpserzhao/ops_gin/first/pkg/e"
	"github.com/DevOpserzhao/ops_gin/first/service/article_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GormStats(c *gin.Context) {
	data := make(map[string]interface{})
	appG := app.Gin{C: c}

	dbjk := article_service.Article{}
	Stats := dbjk.StatsDB()

	data["jk"] = Stats

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
