package provider

import (
	"github.com/datakube/datakube/log"
	target "github.com/datakube/datakube/pkg/apis/backuptarget/v1"
	datakubeClientSet "github.com/datakube/datakube/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
)


var list_all = metav1.ListOptions{}

type kubernetesClientImpl struct {
	client kubernetes.Interface
	dc datakubeClientSet.Interface
}

func NewKubernetesClient(creds kubeCredentials) kubernetesClientImpl{
	c, d := getKubernetesClient(creds)

	return kubernetesClientImpl{
		client: c,
		dc: d,
	}
}


func getKubernetesClient(creds kubeCredentials) (kubernetes.Interface, datakubeClientSet.Interface) {

	var config *rest.Config

	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" && os.Getenv("KUBERNETES_SERVICE_PORT") != "" {
		config, _ = rest.InClusterConfig()
	} else {
		log.Infof("Creating cluster-external Provider client")
		config = &rest.Config{
			Host:        creds.Addr,
			BearerToken: creds.Token,
		}
	}

	// generate the client based off of the config
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("getClusterConfig: %v", err)
	}

	myresourceClient, err := datakubeClientSet.NewForConfig(config)
	if err != nil {
		log.Fatalf("getClusterConfig: %v", err)
	}

	log.Info("Successfully constructed k8s client")
	return client, myresourceClient
}


func (k kubernetesClientImpl) WatchAll(namespaces []string, resultChan chan <- string){

	if len(namespaces) == 0 {
		namespaces = append(namespaces, metav1.NamespaceAll)
	}

	var chans  []<-chan watch.Event
	for _, ns := range namespaces {
		watcher, err := k.dc.DatakubeV1().BackupTargets(ns).Watch(list_all)

		if err != nil {
			return
		}

		chans = append(chans, watcher.ResultChan())
	}

	for _, eventChan := range chans {
		for event := range eventChan {
			target, ok := event.Object.(*target.BackupTarget)
			if !ok {
				log.Fatal("unexpected type")
			}
			resultChan <- target.Namespace
		}
	}
}

func  (k kubernetesClientImpl) LoadTargets(ns string) []target.BackupTarget{
	targetList, err := k.dc.DatakubeV1().BackupTargets(ns).List(list_all)

	if err != nil {
		log.Error("Error getting targets from k8s api ", err.Error())
	}

	return targetList.Items
}
