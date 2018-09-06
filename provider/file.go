package provider

import (
	"github.com/BurntSushi/toml"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/types"
	"github.com/fsnotify/fsnotify"
)

type FileTargets struct {
	TargetFile string
	Dir string
}

func (f *FileTargets) Provide(targetChan chan<- types.ConfigTargets){

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	var targets types.ConfigTargets
	targets = f.loadTargets(f.TargetFile)

	targetChan <- targets

	err = watcher.Add(f.Dir)
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			fileName := event.Name

			targets = f.loadTargets(fileName)
			targetChan <- targets
		case err := <-watcher.Errors:
			log.Errorf("Watcher event error: %s", err)
		}
	}
}

func (f *FileTargets) loadTargets(file string) types.ConfigTargets {
	var targets types.ConfigTargets
	_, err := toml.DecodeFile(file, &targets)

	if err != nil {
		log.Errorf("Error loading the provider file %s => %s", file, err.Error())
	}

	return targets
}
