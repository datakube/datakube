package provider

import (
	"errors"
	"fmt"
	"github.com/datakube/datakube/log"
	"github.com/datakube/datakube/pkg/apis/backuptarget/v1"
	datakubeClientSet "github.com/datakube/datakube/pkg/client/clientset/versioned"
	datakubeInformers "github.com/datakube/datakube/pkg/client/informers/externalversions"
	v12 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
	"strings"
	"time"
)

var list_all = metav1.ListOptions{}

const resyncPeriod = 10 * time.Minute

type kubernetesClientImpl struct {
	client                    kubernetes.Interface
	dc                        datakubeClientSet.Interface
	backupEventHandler        backupEventHandler
	kubeInformerFactories     map[string]kubeinformers.SharedInformerFactory
	datakubeInformerFactories map[string]datakubeInformers.SharedInformerFactory
	isNamespaceAll            bool
}

func NewKubernetesClient(addr string, token string, caFile string) *kubernetesClientImpl {
	c, d := getKubernetesClient(addr, token, caFile)

	return &kubernetesClientImpl{
		client:                    c,
		dc:                        d,
		backupEventHandler:        NewBackupEventHandler(),
		kubeInformerFactories:     make(map[string]kubeinformers.SharedInformerFactory),
		datakubeInformerFactories: make(map[string]datakubeInformers.SharedInformerFactory),
	}
}

func getKubernetesClient(addr string, token string, caFile string) (kubernetes.Interface, datakubeClientSet.Interface) {

	var config *rest.Config

	if os.Getenv("KUBERNETES_SERVICE_HOST") != "" && os.Getenv("KUBERNETES_SERVICE_PORT") != "" {
		config, _ = rest.InClusterConfig()
	} else {
		log.Infof("Creating cluster-external Provider client")
		config = &rest.Config{
			Host:        addr,
			BearerToken: token,
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

func (k *kubernetesClientImpl) WatchAll(namespaces []string, stopChan <-chan struct{}) (chan string, error) {

	eventCh := make(chan string, 1)

	k.backupEventHandler.Notify = eventCh

	if len(namespaces) == 0 {
		namespaces = append(namespaces, metav1.NamespaceAll)
		k.isNamespaceAll = true
	}

	for _, ns := range namespaces {
		kubeFactory := kubeinformers.NewSharedInformerFactory(k.client, resyncPeriod)
		dcFactory := datakubeInformers.NewSharedInformerFactory(k.dc, resyncPeriod)

		dcFactory.Datakube().V1().BackupTargets().Informer().AddEventHandler(k.backupEventHandler)
		kubeFactory.Core().V1().Pods().Lister()
		kubeFactory.Apps().V1().Deployments().Lister()

		k.kubeInformerFactories[ns] = kubeFactory
		k.datakubeInformerFactories[ns] = dcFactory
	}

	for _, ns := range namespaces {
		k.kubeInformerFactories[ns].Start(stopChan)
		k.datakubeInformerFactories[ns].Start(stopChan)
	}

	for _, ns := range namespaces {
		for t, ok := range k.datakubeInformerFactories[ns].WaitForCacheSync(stopChan) {
			if !ok {
				log.Errorf("timed out waiting for controller caches to sync %s in namespace %s", t.String(), ns)
			}
		}
		for b, ok := range k.kubeInformerFactories[ns].WaitForCacheSync(stopChan) {
			if !ok {
				log.Errorf("timed out waiting for controller caches to sync %s in namespace %s", b.String(), ns)
			}
		}
	}

	return eventCh, nil
}

func (k *kubernetesClientImpl) ListTargetsByNs(ns string) []*v1.BackupTarget {
	all, _ := labels.Parse("")
	backups, err := k.datakubeInformerFactories[k.lookupNamespace(ns)].Datakube().V1().BackupTargets().Lister().BackupTargets(ns).List(all)

	if err != nil {
		log.Errorf(err.Error())
	}

	return backups
}

func (k *kubernetesClientImpl) GetTargetByNamespaceAndName(ns string, name string) (*v1.BackupTarget, error) {
	backup, err := k.datakubeInformerFactories[k.lookupNamespace(ns)].Datakube().V1().BackupTargets().Lister().BackupTargets(ns).Get(name)

	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return backup, nil
}

func (k *kubernetesClientImpl) GetPodsByNamesapceAndSelector(namespace string, selector string) ([]*corev1.Pod, error) {

	sel, err := labels.Parse(selector)

	if err != nil {
		fmt.Printf("error %s", err.Error())
		return nil, err
	}

	pods, err := k.kubeInformerFactories[k.lookupNamespace(namespace)].Core().V1().Pods().Lister().List(sel)

	if err != nil {
		fmt.Printf("error %s", err.Error())
		return nil, err
	}

	return pods, nil
}

func (k *kubernetesClientImpl) CreateDeployment(ns string, deployment v12.Deployment) (*v12.Deployment, error) {
	deploymentClient := k.client.AppsV1().Deployments(ns)

	result, err := deploymentClient.Create(&deployment)

	if err != nil {
		log.Error("Error creating deployment ", err.Error())
		return &deployment, err
	}

	return result, nil
}

func (k *kubernetesClientImpl) DeleteDeployment(ns string, deployment string) error {
	deploymentClient := k.client.AppsV1().Deployments(ns)

	err := deploymentClient.Delete(deployment, nil)

	return err
}

func (k *kubernetesClientImpl) GetDeploymentByTarget(target v1.BackupTarget) (*v12.Deployment, error) {

	sel, err := labels.Parse(labels.FormatLabels(target.Spec.Selector))

	if err != nil {
		log.Error("Parsing error for label ", err.Error())
		return nil, err
	}

	selectorArray := strings.Split(sel.String(), "=")
	selectorName := selectorArray[0]
	value := selectorArray[1]

	if selectorName != "app" {
		return nil, errors.New("No proper selector on the target")
	}

	deployment, err := k.kubeInformerFactories[k.lookupNamespace(target.Namespace)].Apps().V1().Deployments().Lister().Deployments(target.Namespace).Get(value)

	if err != nil {

	}

	return deployment, err
}

func (k *kubernetesClientImpl) GetPodsByTarget(target v1.BackupTarget) ([]*corev1.Pod, error) {

	sel, err := labels.Parse(labels.FormatLabels(target.Spec.Selector))

	if err != nil {
		log.Error("Parsing error for label ", err.Error())
		return nil, err
	}

	pods, err := k.GetPodsByNamesapceAndSelector(target.Namespace, sel.String())
	return pods, err
}

func (k *kubernetesClientImpl) lookupNamespace(ns string) string {
	if k.isNamespaceAll {
		return metav1.NamespaceAll
	}
	return ns
}
