package test

import (
	target "github.com/datakube/datakube/pkg/apis/backuptarget/v1"
	"github.com/stretchr/testify/mock"
	v12 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
)

type KuberntesClientMock struct {
	mock.Mock
	Emit bool
}

func (k KuberntesClientMock) WatchAll(namespaces []string, stopchan <- chan struct{}) (chan string, error) {
	return make(chan string), nil
}

func (k KuberntesClientMock) ListTargetsByNs(namespaces string) []*target.BackupTarget {
	args := k.Called(namespaces)
	return []*target.BackupTarget{args.Get(0).(*target.BackupTarget)}
}

func (k KuberntesClientMock) GetTargetByNamespaceAndName(ns string, name string) (*target.BackupTarget, error) {
	return &target.BackupTarget{}, nil
}

func (k KuberntesClientMock) GetPodsByNamesapceAndSelector(ns string, selector string) ([]*v1.Pod, error) {
	return []*v1.Pod{}, nil
}

func (k KuberntesClientMock) CreateDeployment(string, v12.Deployment) (*v12.Deployment, error) {
	return &v12.Deployment{}, nil
}

func (k KuberntesClientMock) GetPodsByTarget(target target.BackupTarget) ([]*v1.Pod, error) {
	args := k.Called(target)
	return []*v1.Pod{args.Get(0).(*v1.Pod)}, nil
}

func (k KuberntesClientMock) GetDeploymentByTarget(target target.BackupTarget) (*v12.Deployment, error) {
	args := k.Called(target)
	return args.Get(0).(*v12.Deployment), nil
}

func (k KuberntesClientMock) DeleteDeployment(string, string) error {
	return nil
}
