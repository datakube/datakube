package job

import (
	"fmt"
	"github.com/SantoDE/datahamster/types"
	"time"
)

type jobStore interface {
	GetLatestJobByTargetName(targetName string) (types.Job, error)
}

func ValidateJobNeededByTarget(target types.Target, store jobStore) bool {

	job, err := store.GetLatestJobByTargetName(target.Name)

	if err != nil {
		fmt.Print("Error fetching jobs for provider %s. Error %s", target.Name, err)
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
	}

	if nextDate.After(time.Now()) {
		return false
	}

	return true
}
