package k8s

// 下面是使用例子
//import (
//	"k8s.io/client-go/kubernetes"
//
//	"github.com/penguinn/go-sdk/k8s"
//	"github.com/penguinn/goframe/config"
//)
//
//var client *kubernetes.Clientset
//
//func Init() error {
//	cfg, err := k8s.NewConfig(&k8s.K8sConfig{
//		K8sHost: config.Config.Kubernetes.Host,
//		Jwt:     config.Config.Kubernetes.JWT,
//	})
//	if err != nil {
//		return err
//	}
//
//	client, err = kubernetes.NewForConfig(cfg)
//	if err != nil {
//		return err
//	}
//	return nil
//}
