package hamster

import (
	"fmt"
	"github.com/SantoDE/datahamster/worker/configuration"
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

// Run Main function to create a scheduler and a hamster and wire it together
func (manager *Manager) Run(exit chan struct{}) error {

	schedule := make(chan int, 1)

	scheduler := NewScheduler(manager.globalConfiguration.Schedule)
	scheduler.run(schedule)

	hamster := NewHamster(manager.globalConfiguration.Database, manager.globalConfiguration.Storage)

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
