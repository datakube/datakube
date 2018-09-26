package provider

import (
	"github.com/datakube/datakube/log"
	target "github.com/datakube/datakube/pkg/apis/backuptarget/v1"
	"github.com/datakube/datakube/types"
	v12 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	"strings"
)

type KubernetesProvider struct {
	Addr      string
	Token     string
	CaFile    string
	Advertise string
}

type kubernetesClient interface {
	WatchAll(namespaces []string, stopchan <-chan struct{}) (chan string, error)
	ListTargetsByNs(ns string) []*target.BackupTarget
	GetTargetByNamespaceAndName(ns string, name string) (*target.BackupTarget, error)
	GetPodsByNamesapceAndSelector(ns string, selector string) ([]*v1.Pod, error)
	CreateDeployment(string, v12.Deployment) (*v12.Deployment, error)
	GetPodsByTarget(target target.BackupTarget) ([]*v1.Pod, error)
	GetDeploymentByTarget(target target.BackupTarget) (*v12.Deployment, error)
	DeleteDeployment(string, string) error
}

type kubernetesController interface {
	Run(stopCh <-chan struct{})
}

func (k *KubernetesProvider) Provide(targetChan chan<- types.ConfigTargets, stopChan <-chan struct{}) error {

	client := NewKubernetesClient(k.Addr, k.Token, k.CaFile)

	watchChan, err := client.WatchAll([]string{}, stopChan)
	if err != nil {
		log.Errorf(err.Error())
	}

	controller := NewKubernetesController(client, k.Advertise, client.backupEventHandler)
	// run the controller loop to process items
	go controller.Run(stopChan)

	for {
		select {
		case ns := <-watchChan:
			cfgTargets := k.loadTargets(ns, client)
			targetChan <- cfgTargets
		}
	}

	return nil
}

func (k *KubernetesProvider) loadTargets(ns string, client kubernetesClient) types.ConfigTargets {
	var cfgTargets types.ConfigTargets
	cfgTargets.Provider = "kubernetes"
	targets := client.ListTargetsByNs(ns)

	for _, target := range targets {
		//if not all is set on the target, attempt to lookup values
		if !validateSpec(*target) {
			targetPods, err := client.GetPodsByTarget(*target)

			if err != nil || len(targetPods) == 0 {
				log.Debugf("Cant find any pods for target %s - skipping", target.Name)
				continue
			}

			targetPod := targetPods[0]

			targetDeployment, err := client.GetDeploymentByTarget(*target)

			if err != nil {
				log.Debugf("Cant find any deployments for target %s - skipping", target.Name)
				continue
			}

			settings := prepareSettings(*targetPod, *targetDeployment)

			if settings["password"] != "" {
				target.Spec.Password = settings["password"]
			}

			if settings["user"] != "" {
				target.Spec.User = settings["user"]
			}

			if settings["db"] != "" {
				target.Spec.Name = settings["db"]
			}

			if settings["host"] != "" {
				target.Spec.Host = settings["host"]
			}

			//validate again to check if now all is valid!
			if !validateSpec(*target) {
				log.Debugf("Not all required values are set and are not lookupable for target %s - skipping", target.Name)
				continue
			}
		}

		cfgTarget := types.Target{
			Name: target.Name,
			DBConfig: types.Database{
				DatabasePort:     target.Spec.Port,
				DatabasePassword: target.Spec.Password,
				DatabaseName:     target.Spec.Name,
				DatabaseHost:     target.Spec.Host,
				DatabaseUserName: target.Spec.User,
				DatabaseType:     target.Spec.Type,
			},
			Schedule: types.Schedule{
				Interval: target.Spec.Interval,
			},
		}

		cfgTargets.Targets = append(cfgTargets.Targets, cfgTarget)
	}

	return cfgTargets
}

func prepareSettings(pod v1.Pod, deployment v12.Deployment) map[string]string {

	settings := filterEnvVars(pod.Spec.Containers[0].Env)
	settings["host"] = deployment.Name
	return settings
}

func validateSpec(target target.BackupTarget) bool {

	if target.Name == "" {
		return false
	}

	if target.Spec.Host == "" {
		return false
	}

	if target.Spec.Password == "" {
		return false
	}

	if target.Spec.Port == "" {
		return false
	}

	if target.Spec.Type == "" {
		return false
	}

	if target.Spec.Interval == "" {
		return false
	}

	if target.Spec.User == "" {
		return false
	}

	if target.Spec.Name == "" {
		return false
	}

	return true
}

func filterEnvVars(vars []v1.EnvVar) map[string]string {
	var settings = make(map[string]string)

	for _, envVar := range vars {
		if strings.Contains(strings.ToLower(envVar.Name), "password") {
			settings["password"] = envVar.Value
		}
		if strings.Contains(strings.ToLower(envVar.Name), "user") {
			settings["user"] = envVar.Value
		}
		if strings.Contains(strings.ToLower(envVar.Name), "database") {
			settings["db"] = envVar.Value
		}
		if strings.Contains(strings.ToLower(envVar.Name), "db") {
			settings["db"] = envVar.Value
		}
	}

	return settings
}
