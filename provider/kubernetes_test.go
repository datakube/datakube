package provider

import (
	"github.com/datakube/datakube/internal/test"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestKubernetesProvider_loadTargets(t *testing.T) {

	clientMock := test.KuberntesClientMock{}

	kp := KubernetesProvider{
		"127.0.0.1:8081",
		"",
		"",
	}

	targets := kp.loadTargets("test", clientMock)

	assert.Equal(t, len(targets.Targets), 1)

	target := targets.Targets[0]

	assert.Equal(t, target.Name, "test-target")
	assert.Equal(t, target.Schedule.Interval, "test")
	assert.Equal(t, target.DBConfig.DatabaseType, "mysql")
	assert.Equal(t, target.DBConfig.DatabaseUserName, "user")
	assert.Equal(t, target.DBConfig.DatabaseHost, "localhost")
	assert.Equal(t, target.DBConfig.DatabasePassword, "password")
	assert.Equal(t, target.DBConfig.DatabasePort, "port")
	assert.Equal(t, target.DBConfig.DatabaseName, "dbname")
}

func TestKubernetesProvider_watch(t *testing.T) {

	clientMock := test.KuberntesClientMock{}

	kp := KubernetesProvider{
		"127.0.0.1:8081",
		"",
		"",
	}

	watchChan := make(chan string)

	clientMock.Emit = true
	kp.watch(clientMock, watchChan)

	select {
	case ns := <-watchChan:
		assert.Equal(t, ns, "test")
	}
}