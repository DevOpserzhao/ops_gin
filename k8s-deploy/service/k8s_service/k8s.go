package k8s_service

import (
	"fmt"
	clientgo "github.com/DevOpserzhao/ops_gin/k8s-deploy/pkg/client-go"
	appv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
)

// 分层代码，一般将前端所传的参数申明成在一个结构体内
type Deployment struct {
	NameSpaceName  string
	DeploymentName string
	Replica        int32
	Image          string
}

func (d *Deployment) DeployM() error {
	return clientgo.Deploy(d.NameSpaceName, d.DeploymentName, d.Replica, d.Image)
}

func (d *Deployment) GetNameSpaces() (data []string, err error) {
	return clientgo.GetNameSpaces()
}

func (d *Deployment) GetNameSpace() (*apiv1.Namespace, error) {
	fmt.Printf("service 传参k8s namespace=%v\n", d.DeploymentName)
	return clientgo.GetNameSpace(d.NameSpaceName)
}

func (d *Deployment) ExistByNameSpace() (bool, error) {
	return clientgo.ExistByNameSpace(d.NameSpaceName)
}
func (d *Deployment) GetDeploymentsAll() (data []string, err error) {
	return clientgo.GetDeploymentsAll(d.NameSpaceName)
}

func (d *Deployment) GetDeploymentsAllCount() (data int, err error) {
	return clientgo.GetDeploymentsAllCount(d.NameSpaceName)
}

func (d *Deployment) GetDeployment() (*appv1.Deployment, error) {

	return clientgo.GetDeployment(d.NameSpaceName, d.DeploymentName)
}

func (d *Deployment) GetNameSpacesCount() (data int32, err error) {
	return clientgo.GetNameSpacesCount()
}

func (d *Deployment) ExistByDeployment() (bool bool, err error) {
	return clientgo.ExistByDeployment(d.NameSpaceName, d.DeploymentName)
}
