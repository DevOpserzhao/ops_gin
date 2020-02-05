package k8s

import (
	"fmt"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/pkg/app"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/pkg/client-go"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetNamespacesall(c *gin.Context) {

	Namespacesall, _ := client_go.GetNameSpaces()
	data := make(map[string]interface{})
	//fmt.Printf("na%v", Namespacesall)

	data["namespaces"] = Namespacesall

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  200,
		"data": data,
	})

}

func GetNamespaces(c *gin.Context) {
	appN := app.Gin{C: c}
	name := com.StrTo(c.Param("id")).String()

	fmt.Printf("k8s namespace%v", name)
	data := make(map[string]interface{})

	count, _ := client_go.GetNameSpacesCount()
	fmt.Printf("k8s namespace num %v", count)
	data["count"] = count
	exists, err := client_go.ExistByNameSpace(name)
	fmt.Print(exists)

	if err != nil {
		data["namespaces"] = nil
		appN.Response(http.StatusInternalServerError, 5000, data)
		return
	}

	if !exists {
		appN.Response(http.StatusOK, 222, nil)

		return
	}
	Namespace, _ := client_go.GetNameSpace(name)

	fmt.Printf("na%v", Namespace)

	data["namespaces"] = Namespace

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  200,
		"data": data,
	})

}
