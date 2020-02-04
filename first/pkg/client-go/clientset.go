package client_go

import (
	"flag"
	"fmt"
	appv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

var ClientSetConn *kubernetes.Clientset

func Setup() {
	var kubeconfig *string
	//kubeconfig = flag.String("kubeconfig", "/etc/kubernetes/admin.conf", "absolute path to the kubeconfig file")

	if home, _ := os.Getwd(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "/etc/kubernetes/admin.conf", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	//在 kubeconfig 中使用当前上下文环境，config 获取支持 url 和 path 方式
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// 根据指定的 config 创建一个新的 clientset
	ClientSetConn, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

}

func GetNameSpaces() ([]v1.Namespace, error) {
	ns, err := ClientSetConn.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	nss := ns.Items
	return nss, err

}

func GetNameSpace(NameSpace string) (*v1.Namespace, error) {
	ns, err := ClientSetConn.CoreV1().Namespaces().Get(NameSpace, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	return ns, err
}

func ExistByNameSpace(NameSpace string) (bool, error) {
	ns, err := ClientSetConn.CoreV1().Namespaces().Get(NameSpace, metav1.GetOptions{})

	//if err != nil {
	//	panic(err)
	//}
	if ns != nil {
		return true, err
	}

	return false, nil
}

func GetDeploymentsAll(NameSpace string) (*appv1.DeploymentList, error) {
	deployments, err := ClientSetConn.AppsV1().Deployments(NameSpace).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	return deployments, err
}

func GetDeployment(NameSpace string, DeploymentName string) (*appv1.Deployment, error) {
	deployments, err := ClientSetConn.AppsV1().Deployments(NameSpace).Get(DeploymentName, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	return deployments, err
}

func int32Ptr(i int32) *int32 { return &i }

func SetDeployment(NameSpace string, DeploymentName string, Replicas int32, Image string) error {
	deploymentsClient := ClientSetConn.AppsV1().Deployments(NameSpace)
	result, getErr := deploymentsClient.Get(DeploymentName, metav1.GetOptions{})
	if getErr != nil {
		panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
	}

	result.Spec.Replicas = int32Ptr(Replicas) // reduce replica count
	if Image != "" {
		result.Spec.Template.Spec.Containers[0].Image = Image // change nginx version
	}

	_, updateErr := deploymentsClient.Update(result)

	return updateErr
}
