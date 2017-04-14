package main

import "fmt"

type Manager struct {
	globalConfiguration GlobalConfiguration
}

func NewManager(globalConfiguration GlobalConfiguration) *Manager {
	manager := new(Manager)
	manager.globalConfiguration = globalConfiguration

	return manager
}

func (manager *Manager) run(exit chan struct{}) error{

	schedule := make(chan int, 1)

	scheduler := NewScheduler(manager.globalConfiguration.Schedule);
	scheduler.run(schedule);

	for {
		select {
		case <- schedule:
			fmt.Printf("Schedule ticked\n")

		case <- exit:
			fmt.Printf("Exit - ESCAPE!")
			return nil
		}
	}
	return nil
}