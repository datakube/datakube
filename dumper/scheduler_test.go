package dumper_test

import (
	"testing"
	"github.com/SantoDE/datahamster/dumper"
	"github.com/stretchr/testify/assert"
	"github.com/SantoDE/datahamster/dumper/jobs"
	"github.com/SantoDE/datahamster/configuration"
)

func TestNewScheduler(t *testing.T) {
	scheduler := dumper.NewScheduler()
	assert.NotNil(t, scheduler)
}

func TestSchedule(t *testing.T) {
	scheduler := dumper.NewScheduler()

	config := new(configuration.ScheduleConfiguration)
	config.At = "15:00"
	config.Interval = "weekly"
	config.Day = "3"

	j := new(jobs.DumpJob)
	scheduler.Schedule(config, j)

	assert.NotNil(t, scheduler)

	entries := scheduler.Cron.Entries()

	assert.Equal(t,1,  len(entries))

	for _, entry := range entries {
		assert.NotNil(t, entry.Schedule.Next)
		assert.NotNil(t, entry.Job)
	}
}
