package store_test

import (
	"github.com/SantoDE/datahamster/internal/store"
	"github.com/SantoDE/datahamster/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSaveJobOk(t *testing.T) {
	store := store.NewTestDataStore()
	defer store.Close()

	now := time.Now()
	job := types.Job{
		State:  "test",
		Target: "12345",
		RunAt:  now,
	}

	savedJob, err := store.SaveJob(job)
	assert.Nil(t, err)
	assert.NotNil(t, savedJob.ID)
	assert.Equal(t, savedJob.State, "test")
	assert.Equal(t, savedJob.Target, "12345")
	assert.Equal(t, savedJob.RunAt, now)
}

func TestAllByTargetName(t *testing.T) {
	store := store.NewTestDataStore()
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
}

func TestGetLatestbyTargetName(t *testing.T) {
	store, err := store.NewStore("/tmp/test.db")
	defer store.Close()

	assert.Nil(t, err)

	err = store.Open()
	assert.Nil(t, err)

	job1 := types.Job{
		Target: "12345",
		RunAt:  time.Now(),
		State:  "IShouldBeRetunred",
	}

	job2 := types.Job{
		Target: "12345",
		RunAt:  time.Now().AddDate(0, 0, -1),
		State:  "IShouldNotBeRetunred",
	}

	store.SaveJob(job1)
	store.SaveJob(job2)

	job, err := store.GetLatestJobByTargetName("12345")
	assert.Nil(t, err)
	assert.NotNil(t, job)
	assert.Equal(t, job.State, "IShouldBeRetunred")
	assert.Equal(t, job.Target, "12345")
}
