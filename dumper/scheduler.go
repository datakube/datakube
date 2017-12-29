package dumper

import (
	"github.com/robfig/cron"
	"github.com/SantoDE/datahamster/configuration"
	"strconv"
	"strings"
	"github.com/SantoDE/datahamster/dumper/jobs"
)

type Scheduler struct {
	Cron *cron.Cron
}

func NewScheduler() *Scheduler{
	s := new(Scheduler)
	s.Cron = cron.New()

	return s
}

func (s *Scheduler) Schedule(targetSchedule *configuration.ScheduleConfiguration, j *jobs.DumpJob) {

	schedule := new(cron.SpecSchedule)
	day, _ := strconv.ParseUint(targetSchedule.Day, 10, 64)
	minutes, _ := strconv.ParseUint(strings.Split(targetSchedule.At, ":")[0], 10, 64)
	seconds, _ := strconv.ParseUint(strings.Split(targetSchedule.At, ":")[1], 10, 64)

	switch targetSchedule.Interval {
	case "monthly":
		schedule.Dom = day
		schedule.Minute = minutes
		schedule.Second = seconds
	case "weekly":
		schedule.Dow = day
		schedule.Minute = minutes
		schedule.Second = seconds
	case "daily":
		schedule.Minute = minutes
		schedule.Second = seconds
	}

	s.Cron.Schedule(schedule, j)
}