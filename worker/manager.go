package main

import (
	"fmt"
	"github.com/santode/datahamster/worker/configuration"
)

// Manager struct which holds the application configuration
type Manager struct {
	globalConfiguration configuration.GlobalConfiguration
}

// NewManager with the application config
func NewManager(globalConfiguration configuration.GlobalConfiguration) *Manager {
	manager := new(Manager)
	manager.globalConfiguration = globalConfiguration

	return manager
}

func (manager *Manager) run(exit chan struct{}) error {

	schedule := make(chan int, 1)

	scheduler := NewScheduler(manager.globalConfiguration.Schedule)
	scheduler.run(schedule)

	hamster := NewHamster(manager.globalConfiguration.Database)

	for {
		select {
		case <-schedule:
			fmt.Printf("Schedule ticked\n")
			hamster.run()

		case <-exit:
			fmt.Printf("Exit - ESCAPE!")
			return nil
		}
	}
	return nil
}
