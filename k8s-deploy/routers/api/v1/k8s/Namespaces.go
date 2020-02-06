package k8s

import (
	"fmt"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/pkg/app"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/pkg/client-go"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/service/k8s_service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetNamespacesall(c *gin.Context) {

	Namespacesall, _ := client_go.GetNameSpaces()
	count, _ := client_go.GetNameSpacesCount()
	data := make(map[string]interface{})
	//fmt.Printf("na%v", Namespacesall)

	data["namespaces"] = Namespacesall

	fmt.Printf("k8s namespace num %v", count)
	data["count"] = count
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  200,
		"data": data,
	})

}

func GetNamespaces(c *gin.Context) {
	appN := app.Gin{C: c}
	name := com.StrTo(c.Param("id")).String()

	fmt.Printf("url 传参k8s namespace=%v\n", name)
	data := make(map[string]interface{})
	NamespaceService := k8s_service.Deployment{NameSpaceName: name}

	exists, err := NamespaceService.ExistByNameSpace()

	fmt.Print(exists)

	if err != nil {
		data["namespaces"] = nil
		appN.Response(http.StatusInternalServerError, 5000, data)
		return
	}

	//if exists == {
	//	data["namespaces"] = nil
	//	appN.Response(http.StatusInternalServerError, 5000, data)
	//
	//
	//	return
	//}
	Namespace, _ := NamespaceService.GetNameSpace()

	//fmt.Printf("na%v", Namespace)

	data["namespaces"] = Namespace

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  200,
		"data": data,
	})

}
