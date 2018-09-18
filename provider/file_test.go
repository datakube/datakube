package provider_test

import (
	"io/ioutil"
	"os"
	"testing"
	"text/template"

	"github.com/datakube/datakube/provider"
	"github.com/datakube/datakube/types"
	"github.com/stretchr/testify/assert"
)

type testTargets struct {
	Targets []testData
}

type testData struct {
	Host     string
	User     string
	PW       string
	Port     string
	Interval string
	Name     string
}

func TestProvide(t *testing.T) {

	tempDir, _ := ioutil.TempDir("", "")
	tempFile, _ := ioutil.TempFile(tempDir, "")

	defer os.RemoveAll(tempDir)

	var data []testData

	data = append(data, testData{
		Host:     "localhost",
		User:     "test",
		PW:       "12345",
		Interval: "weekly",
		Name:     "testtarget",
		Port:     "3306",
	})

	generateTestFile(tempFile, data)

	ft := provider.FileTargets{
		tempFile.Name(),
		tempDir,
	}
	testChan := make(chan types.ConfigTargets)
	stopChan := make(chan bool)

	go func() {
		err := ft.Provide(testChan, stopChan)
		assert.NoError(t, err)
	}()

	select {
	case config := <-testChan:
		assert.Equal(t, len(config.Targets), 1)
		stopChan <- true
	}

	os.RemoveAll(tempDir)
}

func TestProvideWatch(t *testing.T) {

	tempDir, err := ioutil.TempDir("", "datakube")

	if err != nil {
		t.Fatal(t, err.Error())
	}

	tempFile, err := ioutil.TempFile(tempDir, "targets")

	if err != nil {
		t.Fatal(t, err.Error())
	}

	//defer os.RemoveAll(tempDir)

	ft := provider.FileTargets{
		tempFile.Name(),
		tempDir,
	}

	testChan := make(chan types.ConfigTargets)
	stopChan := make(chan bool)
	inChan := make(chan struct{})

	go func() {
		err := ft.Provide(testChan, stopChan)
		assert.NoError(t, err)
	}()

	firstTargets := <-testChan

	assert.Equal(t, len(firstTargets.Targets), 0)

	var data []testData

	data = append(data, testData{
		Host:     "localhost",
		User:     "test",
		PW:       "12345",
		Interval: "weekly",
		Name:     "testtarget",
		Port:     "3306",
	})

	generateTestFile(tempFile, data)

	go func() {
		targets := <-testChan
		assert.Equal(t, len(targets.Targets), 1)
		stopChan <- true
		inChan <- struct{}{}
	}()

	<-inChan
}

func generateTestFile(file *os.File, data []testData) {

	var tmpl = `
[[target]]
    name = "test_target"
    [target.db]
        type = "mysql"
        host = "localhost"
        name = "test"
        user = "root"
        password = "root"
        port = "3306"
        [target.db.sql]
        tempdir = "/tmp"
    [target.schedule]
        interval = "weekly"
        day = "2"
        at = "03:00"
`
	configTemplate, _ := template.New(file.Name()).Parse(tmpl)

	testTargets := testTargets{
		Targets: data,
	}

	err := configTemplate.Execute(file, testTargets)
	defer file.Close()
	if err != nil {
		panic(err)
	}
}
