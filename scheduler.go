package main

import (
	"fmt"
	"time"
)

type Scheduler struct {
	duration 	time.Duration
}

func NewScheduler(configuration ScheduleConfiguration) *Scheduler {

	scheduler := new(Scheduler)

	duration, err := time.ParseDuration(configuration.Interval)

	if err != nil {
		fmt.Printf("Error Parsing Duration")
	}

	scheduler.duration = duration

	return scheduler
}

func (scheduler *Scheduler) run(tick chan(int)) {

	ticker := time.NewTicker(scheduler.duration)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at \n", t)
			tick <- 1
		}
	}()
}


