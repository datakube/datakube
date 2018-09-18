package provider

import (
	"github.com/BurntSushi/toml"
	"github.com/datakube/datakube/log"
	"github.com/datakube/datakube/types"
	"github.com/fsnotify/fsnotify"
)

type FileTargets struct {
	TargetFile string
	Dir        string
}

func (f *FileTargets) Provide(targetChan chan<- types.ConfigTargets, stopChan <-chan bool) error {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	var targets types.ConfigTargets
	targets, _ = loadTargets(f.TargetFile)

	targetChan <- targets

	watcher.Add(f.Dir)
	if err != nil {
		log.Fatal(err)
		return err
	}

	go func() {
		defer watcher.Close()
		for {
			select {
			case event, ok := <-watcher.Events:
				log.Debug("Watcher event:", event)
				if !ok {
					return
				}

				fileName := event.Name
				targets, err = loadTargets(fileName)

				if err == nil {
					targetChan <- targets
				}
			case err := <-watcher.Errors:
				log.Errorf("Watcher event error: %s", err)
			case <-stopChan:
				return
			}
		}
	}()
	return nil
}

func loadTargets(file string) (types.ConfigTargets, error) {
	var targets types.ConfigTargets
	_, err := toml.DecodeFile(file, &targets)

	if err != nil {
		log.Errorf("Error loading the provider file %s => %s", file, err.Error())
		return targets, err
	}

	return targets, nil
}
