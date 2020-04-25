基于k8s进行发布的接口

在k8s环境部署需要增加权限


`If you have RBAC enabled on your cluster, use the following snippet to create role binding which will grant the default service account view permissions.
kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:defaul`