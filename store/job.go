package store

import (
	"github.com/SantoDE/datahamster/types"
	"github.com/SantoDE/datahamster/log"
	"github.com/asdine/storm/q"
)

//Save function to save the given Dumper
func (d *DataStore) SaveJob(job types.Job) (types.Job, error) {
	err := d.db.Save(&job)

	if err != nil {
		return *new(types.Job), err
	}

	return job, nil
}

func (d *DataStore) AllJobsByTargetName(name string) ([]types.Job, error) {
	var jobs []types.Job

	err := d.db.Find("Target", name, &jobs)

	if err != nil {
		log.Error("Can't find all Jobs by Target Name", err)
		return jobs, err
	}

	return jobs, nil
}

func (d *DataStore) GetLatestJobByTargetName(name string) (types.Job, error) {

	var job types.Job
	query := d.db.Select(q.Eq("Target", name)).OrderBy("RunAt").Reverse()
	err := query.First(&job)

	return job, err
}

func (d *DataStore) AllQueued() ([]types.Job, error) {

	var jobs []types.Job

	err := d.db.Find("Status", types.STATUS_QUEUED, &jobs)

	if err != nil {
		log.Error("Can't find all Jobs by Target Name", err)
		return jobs, err
	}

	return jobs, nil
}

func (d *DataStore) ListAllJobs() ([]types.Job, error) {
	var jobs []types.Job

	err := d.db.All(&jobs)

	if err != nil {
		log.Error("Can't list all Jobs", err)
		return jobs, err
	}

	return jobs, nil
}