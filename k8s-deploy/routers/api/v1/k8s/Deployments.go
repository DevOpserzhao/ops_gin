package k8s

import (
	"fmt"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/pkg/app"
	"github.com/DevOpserzhao/ops_gin/k8s-deploy/service/k8s_service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// @Summary Deployments
// @Produce  json
// @Param username query string true "namespace"
// @Router /Namespaces [get]
func Deployments(c *gin.Context) {
	appD := app.Gin{C: c}
	//DeploymentName := com.StrTo(c.Param("id")).String()
	data := make(map[string]interface{})
	arg := c.Query("namespace")
	name := com.StrTo(arg).String()
	//fmt.Printf(name)
	DeployService := k8s_service.Deployment{NameSpaceName: name}
	Namespace, _ := DeployService.GetDeploymentsAll()
	Count, _ := DeployService.GetDeploymentsAllCount()
	//Namespace, _ := client_go.GetDeploymentsAll(name)
	//Count, _ := client_go.GetDeploymentsAllCount(name)

	if name == "" {
		name = "all"
	}
	data["Count"] = Count
	data["namespace"] = name
	data["Deployments"] = Namespace
	appD.Response(http.StatusOK, 200, data)

}

func Deployment(c *gin.Context) {
	appD := app.Gin{C: c}
	DeploymentName := com.StrTo(c.Param("id")).String()
	data := make(map[string]interface{})
	arg := c.Query("namespace")
	name := com.StrTo(arg).String()
	//Namespace, _ := client_go.GetDeployment(name, DeploymentName)

	DeployService := k8s_service.Deployment{NameSpaceName: name, DeploymentName: DeploymentName}
	ExistByDeployment, _ := DeployService.ExistByDeployment()
	if ExistByDeployment == false {
		data["Deployment"] = nil
		appD.Response(http.StatusAccepted, 5001, data)
	}
	Namespace, _ := DeployService.GetDeployment()

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
	//retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
	//	updateErr := client_go.SetDeployment(name, DeploymentName, int32(replica), image)
	//	return updateErr
	//})
	//retryErr := client_go.Deploy(name, DeploymentName, int32(replica), image)
	//if retryErr != nil {
	//	panic(fmt.Errorf("Update failed: %v", retryErr))
	//	appD.Response(http.StatusOK, 5000, "Update failed")
	//}
	DeployService := k8s_service.Deployment{NameSpaceName: name, DeploymentName: DeploymentName, Replica: int32(replica), Image: image}
	DeployserviceErr := DeployService.DeployM()
	if DeployserviceErr != nil {
		panic(fmt.Errorf("Update failed: %v", DeployserviceErr))
		appD.Response(http.StatusOK, 5000, "Update failed")
	}
	data["Deployment"] = "更新成功"
	appD.Response(http.StatusOK, 5001, data)

}
