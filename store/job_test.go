package store_test

import (
	"github.com/SantoDE/datahamster/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSaveJobOk(t *testing.T) {
	store := NewTestDataStore()
	defer store.Close()

	now := time.Now()
	job := types.Job{
		Status: "test",
		Target: "12345",
		RunAt:  now,
	}

	savedJob, err := store.SaveJob(job)
	assert.Nil(t, err)
	assert.NotNil(t, savedJob.ID)
	assert.Equal(t, savedJob.Status, "test")
	assert.Equal(t, savedJob.Target, "12345")
	assert.Equal(t, savedJob.RunAt, now)
}

func TestAllByTargetName(t *testing.T) {
	store := NewTestDataStore()
	defer store.Close()

	job1 := types.Job{
		Target: "12345",
	}

	job2 := types.Job{
		Target: "12345",
	}

	job3 := types.Job{
		Target: "aaaaaa",
	}

	store.SaveJob(job1)
	store.SaveJob(job2)
	store.SaveJob(job3)

	jobs, err := store.AllJobsByTargetName("12345")

	assert.Nil(t, err)
	assert.NotNil(t, jobs)
	assert.Equal(t, len(jobs), 2)

	jobs, err = store.AllJobsByTargetName("bbbbbbbb")

	assert.NotNil(t, err)
	assert.Nil(t, jobs)
	assert.Equal(t, len(jobs), 0)
}

func TestGetLatestbyTargetName(t *testing.T) {
	store := NewTestDataStore()
	defer store.Close()

	job1 := types.Job{
		Target: "12345",
		RunAt:  time.Now(),
		Status: "IShouldBeRetunred",
	}

	job2 := types.Job{
		Target: "12345",
		RunAt:  time.Now().AddDate(0, 0, -1),
		Status: "IShouldNotBeRetunred",
	}

	store.SaveJob(job1)
	store.SaveJob(job2)

	job, err := store.GetLatestJobByTargetName("12345")
	assert.Nil(t, err)
	assert.NotNil(t, job)
	assert.Equal(t, job.Status, "IShouldBeRetunred")
	assert.Equal(t, job.Target, "12345")
}

func TestSaveJob(t *testing.T) {
	store := NewTestDataStore()
	defer store.Close()

	job1 := types.Job{
		Target: "12345",
		RunAt:  time.Now(),
	}

	err := store.DeleteJob(job1)
	assert.NotNil(t, err)

	job1, _ = store.SaveJob(job1)

	err = store.DeleteJob(job1)
	assert.Nil(t, err)
}

func TestListJobsByStatus(t *testing.T) {
	store := NewTestDataStore()
	defer store.Close()

	job1 := types.Job{
		Target: "12345",
		Status: "teststatus",
	}
	job2 := types.Job{
		Target: "12345",
		Status: "teststatus",
	}
	job3 := types.Job{
		Target: "12345",
		Status: "donotfindme",
	}

	store.SaveJob(job1)
	store.SaveJob(job2)
	store.SaveJob(job3)

	jobs, err := store.ListJobsByStatus("teststatus")

	assert.Nil(t, err)
	assert.NotNil(t, jobs)
	assert.Equal(t, len(jobs), 2)

	jobs, err = store.ListJobsByStatus("imnotexisting")

	assert.NotNil(t, err)
	assert.Nil(t, jobs)
	assert.Equal(t, len(jobs), 0)
}

func TestListAllJobs(t *testing.T) {
	store := NewTestDataStore()
	defer store.Close()

	jobs, err := store.ListAllJobs()

	assert.NotNil(t, jobs)
	assert.Nil(t, err)
	assert.Equal(t, len(jobs), 0)

	job1 := types.Job{
		Target: "12345",
		Status: "teststatus",
	}
	job2 := types.Job{
		Target: "12345",
		Status: "teststatus",
	}
	job3 := types.Job{
		Target: "12345",
		Status: "donotfindme",
	}

	store.SaveJob(job1)
	store.SaveJob(job2)
	store.SaveJob(job3)

	jobs, err = store.ListAllJobs()

	assert.Nil(t, err)
	assert.NotNil(t, jobs)
	assert.Equal(t, len(jobs), 3)
}

func TestDataStore_GetJobById(t *testing.T) {
	store := NewTestDataStore()
	defer store.Close()

	job1 := types.Job{
		Target: "12345",
	}

	job2 := types.Job{
		Target: "123456",
	}

	job1, _ = store.SaveJob(job1)
	job2, _ = store.SaveJob(job2)

	foundJob, err := store.GetJobById(job1.ID)
	assert.Nil(t, err)
	assert.Equal(t, job1.ID, foundJob.ID)

	foundJob, err = store.GetJobById(job2.ID)
	assert.Nil(t, err)
	assert.Equal(t, job2.ID, foundJob.ID)
}
