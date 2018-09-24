package provider

import (
	"fmt"
	"github.com/datakube/datakube/log"
	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"time"
)

type kubernetesControllerImpl struct {
	advertise 		string
	client    		kubernetesClient
	queue			TargetEventQueue
}


type TargetEventQueue interface {
	ShutDown()
	Done(interface{})
	Get()  (interface{}, bool)
}

func NewKubernetesController(client kubernetesClient, advertise string, queue TargetEventQueue) kubernetesControllerImpl {
	return kubernetesControllerImpl{
		advertise: advertise,
		client: client,
		queue: queue,
	}
}

func (c *kubernetesControllerImpl) Setup(stopCh <-chan struct{}, advertise string) {
	c.advertise = advertise
}

func (c *kubernetesControllerImpl)  Run(stopCh <-chan struct{}) {
	// handle a panic with logging and exiting
	defer utilruntime.HandleCrash()
	// ignore new items in the queue but when all goroutines
	// have completed existing items then shutdown
	defer c.queue.ShutDown()

	log.Info("Controller.Run: cache sync complete")

	// run the runWorker method every second with a stop channel
	wait.Until(c.runWorker, time.Second, stopCh)
}

// runWorker executes the loop to process new items added to the queue
func (c *kubernetesControllerImpl) runWorker() {
	log.Info("Controller.runWorker: starting")

	for c.processNextItem() {
		log.Info("Controller.runWorker: processing next item")
	}

	log.Info("Controller.runWorker: completed")
}

func (c *kubernetesControllerImpl) processNextItem() bool {
	log.Info("Controller.processNextItem: start")

	key, quit := c.queue.Get()

	if quit {
		return false
	}

	defer c.queue.Done(key)

	keyRaw := key.(string)

	namespace, name, err := cache.SplitMetaNamespaceKey(keyRaw)

	if err != nil {
		fmt.Printf("Error! %s", err.Error())
	}

	target, err := c.client.GetTargetByNamespaceAndName(namespace, name)

	if err != nil || target == nil{
		c.client.DeleteDeployment(namespace, name)
		return true
	}

	sel := labels.NewSelector()

	var targetPod corev1.Pod

	for k, expr := range target.Spec.Selector {

		req, err := labels.NewRequirement(k, selection.Equals, []string{expr})

		if err != nil {
			fmt.Printf("Parse Error for Req %s", err.Error())
		}

		pods, err := c.client.GetPodsByNamesapceAndSelector(namespace, req.String())

		if err != nil {
			fmt.Printf("Get Pods By NS and Selector error %s", err.Error())
		}

		targetPod = *pods[0]
		log.Debugf("Found targetPod with UiD %s for target %s and selector %s", targetPod.UID, keyRaw, sel.String())

		deployment := createDeployment(target.Name, "datakube/agent:dev", c.advertise)

		result, err := c.client.CreateDeployment(namespace, deployment)

		if err != nil {
			fmt.Printf("Create Deployment error %s", err.Error())
			continue
		}

		log.Debugf("Created deployment with name %s and id %s", result.Name, result.UID)
	}

	// keep the worker loop running by returning true
	return true
}

func createDeployment(deploymentName string, image string, address string) appsv1.Deployment{

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": deploymentName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": deploymentName,
					},
				},

				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  deploymentName,
							Image: image,
							Args: []string{"--logLevel=debug", "--server=" + address, "--interval=10"},
							},
						},
					},
				},
			},
		}

	fmt.Printf("%s", deployment.UID)
	return *deployment
}

func int32Ptr(i int32) *int32 { return &i }