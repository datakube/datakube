package provider

import (
	"github.com/datakube/datakube/internal/test"
	"github.com/magiconair/properties/assert"
	v13 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
	target "github.com/datakube/datakube/pkg/apis/backuptarget/v1"
)

func TestKubernetesProvider_loadTargetsAllGiven(t *testing.T) {

	clientMock := test.KuberntesClientMock{}

	kp := KubernetesProvider{
		"127.0.0.1:8081",
		"",
		"",
		"no-real-advert",
	}

	targetMock := target.BackupTarget{
		ObjectMeta: v1.ObjectMeta{
			Name: "test-target",
			Namespace: "test",
		},
		Spec: target.BackupTargetSpec{
			Interval: "test",
			Type: "mysql",
			User: "user",
			Host: "localhost",
			Port: "3306",
			Password: "password",
			Name: "dbname",
		},
	}

	clientMock.On("ListTargetsByNs", "test").Return(&targetMock)

	targets := kp.loadTargets("test", clientMock)

	assert.Equal(t, len(targets.Targets), 1)

	target := targets.Targets[0]

	assert.Equal(t, target.Name, "test-target")
	assert.Equal(t, target.Schedule.Interval, "test")
	assert.Equal(t, target.DBConfig.DatabaseType, "mysql")
	assert.Equal(t, target.DBConfig.DatabaseUserName, "user")
	assert.Equal(t, target.DBConfig.DatabaseHost, "localhost")
	assert.Equal(t, target.DBConfig.DatabasePassword, "password")
	assert.Equal(t, target.DBConfig.DatabasePort, "3306")
	assert.Equal(t, target.DBConfig.DatabaseName, "dbname")
}

func TestKubernetesProvider_loadTargetsLookUp(t *testing.T) {

	clientMock := test.KuberntesClientMock{}

	kp := KubernetesProvider{
		"127.0.0.1:8081",
		"",
		"",
		"no-real-advert",
	}

	targetMock := target.BackupTarget{
		ObjectMeta: v1.ObjectMeta{
			Name: "test-target",
			Namespace: "test",
		},
		Spec: target.BackupTargetSpec{
			Interval: "test",
			Type: "mysql",
			Port: "3306",
			Selector: map[string]string{"app":"mysql"},
		},
	}

	clientMock.On("ListTargetsByNs", "test").Return(&targetMock)

	podMock := v12.Pod{
		Spec: v12.PodSpec{
			Containers: []v12.Container{
				{
					Env: []v12.EnvVar{
						{
							Name: "MYSQL_USER",
							Value: "user",
						},
						{
							Name: "MYSQL_PASSWORD",
							Value: "password",
						},
						{
							Name: "MYSQL_DATABASE",
							Value: "testdb",
						},
					},
				},
			},
		},
	}

	clientMock.On("GetPodsByTarget", targetMock).Return(&podMock)

	deploymentMock := v13.Deployment{
		ObjectMeta: v1.ObjectMeta{
			Name: "mysql",
		},
	}

	clientMock.On("GetDeploymentByTarget", targetMock).Return(&deploymentMock)

	targets := kp.loadTargets("test", clientMock)

	assert.Equal(t, len(targets.Targets), 1)

	target := targets.Targets[0]

	assert.Equal(t, target.Name, "test-target")
	assert.Equal(t, target.Schedule.Interval, "test")
	assert.Equal(t, target.DBConfig.DatabaseType, "mysql")
	assert.Equal(t, target.DBConfig.DatabaseUserName, "user")
	assert.Equal(t, target.DBConfig.DatabaseHost, "mysql")
	assert.Equal(t, target.DBConfig.DatabasePassword, "password")
	assert.Equal(t, target.DBConfig.DatabasePort, "3306")
	assert.Equal(t, target.DBConfig.DatabaseName, "testdb")
}