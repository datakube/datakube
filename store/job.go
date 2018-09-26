package store

import (
	"github.com/asdine/storm/q"
	"github.com/datakube/datakube/log"
	"github.com/datakube/datakube/types"
)

//Save function to save the given Dumper
func (d *DataStore) SaveJob(job types.Job) (types.Job, error) {
	err := d.db.Save(&job)

	if err != nil {
		return *new(types.Job), err
	}

	return job, nil
}

func (d *DataStore) DeleteJob(job types.Job) error {
	err := d.db.DeleteStruct(&job)

	if err != nil {
		return err
	}

	return nil
}

func (d *DataStore) AllJobsByTargetName(name string) ([]types.Job, error) {
	var jobs []types.Job

	err := d.db.Find("Target", name, &jobs)

	if err != nil {
		log.Error("Can't find all Jobs by Target Name ", err)
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

func (d *DataStore) ListJobsByStatus(status string) ([]types.Job, error) {

	var jobs []types.Job

	err := d.db.Find("Status", status, &jobs)

	if err != nil {
		log.Error("Can't list all Jobs by Status ", status, err)
		return jobs, err
	}

	return jobs, nil
}

func (d *DataStore) ListAllJobs() ([]types.Job, error) {
	var jobs []types.Job

	err := d.db.All(&jobs)

	if err != nil {
		log.Error("Can't list all Jobs ", err)
		return nil, err
	}

	return jobs, nil
}

func (d *DataStore) GetJobById(id int) (types.Job, error) {
	var job types.Job
	err := d.db.One("ID", id, &job)

	if err != nil {
		log.Error("Can't get Job By Id ", err)
	}

	return job, err
}
