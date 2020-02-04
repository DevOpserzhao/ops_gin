package k8s

import (
	"fmt"
	"github.com/DevOpserzhao/ops_gin/first/pkg/app"
	client_go "github.com/DevOpserzhao/ops_gin/first/pkg/client-go"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"k8s.io/client-go/util/retry"
	"net/http"
)

func Deployments(c *gin.Context) {
	appD := app.Gin{C: c}
	//DeploymentName := com.StrTo(c.Param("id")).String()
	data := make(map[string]interface{})
	arg := c.Query("namespace")
	name := com.StrTo(arg).String()
	Namespace, _ := client_go.GetDeploymentsAll(name)

	data["Deployments"] = Namespace
	appD.Response(http.StatusOK, 5000, data)

}

func Deployment(c *gin.Context) {
	appD := app.Gin{C: c}
	DeploymentName := com.StrTo(c.Param("id")).String()
	data := make(map[string]interface{})
	arg := c.Query("namespace")
	name := com.StrTo(arg).String()
	Namespace, _ := client_go.GetDeployment(name, DeploymentName)

	data["Deployment"] = Namespace
	appD.Response(http.StatusOK, 5000, data)

}

func SetDeployment(c *gin.Context) {
	appD := app.Gin{C: c}
	DeploymentName := com.StrTo(c.Param("id")).String()
	data := make(map[string]interface{})
	arg := c.Query("namespace")
	name := com.StrTo(arg).String()

	replica := com.StrTo(c.Query("replica")).MustInt()
	image := com.StrTo(c.Query("image")).String()
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		updateErr := client_go.SetDeployment(name, DeploymentName, int32(replica), image)
		return updateErr
	})

	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
		appD.Response(http.StatusOK, 5000, "Update failed")
	}

	data["Deployment"] = "Namespace"
	appD.Response(http.StatusOK, 5000, data)

}
