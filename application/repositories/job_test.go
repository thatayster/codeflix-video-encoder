package repositories_test

import (
	"encoder/application/repositories"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJobsConversion(t *testing.T) {
	video, job := generateDomainResources()

	jobOrm, err := repositories.ToJobORM(job)

	require.Nil(t, err)
	require.Equal(t, job.Id, jobOrm.ID)
	require.Equal(t, video.Id, jobOrm.Video.ID)

	jobDomain, err := jobOrm.ToJobDomain()

	require.Nil(t, err)
	require.Equal(t, jobOrm.ID, jobDomain.Id)
	require.EqualValues(t, video, jobDomain.Video)
}

func TestJobsInvalidConversion(t *testing.T) {
	_, job := generateDomainResources()
	job.Id = "invalid-id"

	_, err := repositories.ToJobORM(job)

	require.Error(t, err)

	invalidJobOrm := &repositories.Job{
		ID: "invalid-id",
	}

	_, err = invalidJobOrm.ToJobDomain()

	require.Error(t, err)
}