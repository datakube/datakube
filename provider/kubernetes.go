package provider

import (
	"github.com/datakube/datakube/types"
	target "github.com/datakube/datakube/pkg/apis/backuptarget/v1"
)

type KubernetesProvider struct {
	Addr string
	Token string
	CaFile string
}

type kubeCredentials struct {
	Addr string
	Token string
	CaFile string
}

type kubernetesClient interface {
	WatchAll(namespaces []string, resultChan chan <- string)
	LoadTargets(ns string) []target.BackupTarget
}

func (k *KubernetesProvider) Provide(targetChan chan<- types.ConfigTargets, stopChan <-chan bool) error {

	watchChan := make(chan string)

	client := NewKubernetesClient(kubeCredentials{k.Addr, k.Token, k.CaFile})

	k.watch(client, watchChan)

	for {
		select {
		case ns := <-watchChan:
			cfgTargets := k.loadTargets(ns, client)
			targetChan <- cfgTargets
		}
	}

	return nil
}

func (k *KubernetesProvider) watch(client kubernetesClient, watchChan chan string) {

	go func() {
		client.WatchAll([]string{""}, watchChan)
	}()
}

func (k *KubernetesProvider) loadTargets(ns string, client kubernetesClient) types.ConfigTargets{
	var cfgTargets types.ConfigTargets
	cfgTargets.Provider = "kubernetes"
	targets := client.LoadTargets(ns)

	for _, target := range targets {
		cfgTarget := types.Target{
			Name: target.Name,
			DBConfig: types.Database{
				DatabasePort: target.Spec.Port,
				DatabasePassword: target.Spec.Password,
				DatabaseName: target.Spec.Name,
				DatabaseHost: target.Spec.Host,
				DatabaseUserName: target.Spec.User,
				DatabaseType: target.Spec.Type,
			},
			Schedule: types.Schedule{
				Interval: target.Spec.Interval,
			},
		}

		cfgTargets.Targets = append(cfgTargets.Targets, cfgTarget)
	}

	return cfgTargets
}

