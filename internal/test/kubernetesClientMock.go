package test

import (
	target "github.com/datakube/datakube/pkg/apis/backuptarget/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KuberntesClientMock struct {
	Emit bool
}

func (k KuberntesClientMock)  LoadTargets(ns string) []target.BackupTarget {
	var targets []target.BackupTarget

	meta := v1.ObjectMeta{
		Name: "test-target",
		Namespace: "test",
	}

	targets = append(targets, target.BackupTarget{
		ObjectMeta: meta,
		Spec: target.BackupTargetSpec{
			Host: "localhost",
			User: "user",
			Password: "password",
			Port: "port",
			Name: "dbname",
			Type: "mysql",
			Interval: "test",
		},
	})

	return targets
}

func (k KuberntesClientMock) WatchAll(namespaces []string, resultChan chan <- string) {
	if k.Emit {
		resultChan <- "test"
	}
}
