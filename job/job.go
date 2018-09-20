package job

import (
	"github.com/datakube/datakube/log"
	"github.com/datakube/datakube/types"
	"time"
)

type jobStore interface {
	GetLatestJobByTargetName(targetName string) (types.Job, error)
	SaveJob(job types.Job) (types.Job, error)
}

func ValidateJobNeededByTarget(target types.Target, store jobStore) bool {

	job, err := store.GetLatestJobByTargetName(target.Name)

	if err != nil {
		if err.Error() == "not found" {
			return true
		}
		log.Debug("Error fetching for provider with Error ", target.Name, err)
		return false
	}

	if &job == new(types.Job) {
		return true
	}

	if job.Status == types.STATUS_QUEUED {
		//dont create a new job if one is queued already
		return false
	}

	var nextDate time.Time

	switch target.Schedule.Interval {
	case "monthly":
		nextDate = job.RunAt.AddDate(0, 1, 0)
	case "weekly":
		nextDate = job.RunAt.AddDate(0, 0, 7)
	case "daily":
		nextDate = job.RunAt.AddDate(0, 0, 1)
	case "every_minute":
		nextDate = job.RunAt.Add(1 * time.Minute)
	}

	if nextDate.After(time.Now()) {
		return false
	}

	return true
}

func Queue(target string, store jobStore) (types.Job, error) {
	return store.SaveJob(types.Job{
		RunAt:  time.Now(),
		Status: types.STATUS_QUEUED,
		Target: target,
	})
}
