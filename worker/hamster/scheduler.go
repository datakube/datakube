package hamster

import (
	"fmt"
	"github.com/SantoDE/datahamster/worker/configuration"
	"time"
)

// Scheduler struct to hold duration
type Scheduler struct {
	duration time.Duration
	startNow bool
}

// NewScheduler function to create a new scheduler
func NewScheduler(configuration configuration.ScheduleConfiguration) *Scheduler {

	scheduler := new(Scheduler)

	duration, err := time.ParseDuration(configuration.Interval)

	if err != nil {
		fmt.Printf("Error Parsing Duration")
	}

	scheduler.duration = duration
	scheduler.startNow = configuration.StartNow

	return scheduler
}

func (scheduler *Scheduler) run(tick chan (int)) {

	ticker := time.NewTicker(scheduler.duration)
	go func() {

		var startedAlready = false

		if !startedAlready && scheduler.startNow {
			startedAlready = true
			tick <- 1
		}

		for t := range ticker.C {
			fmt.Println("Tick at \n", t)
			tick <- 1
		}
	}()
}
