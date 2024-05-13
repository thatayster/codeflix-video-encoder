package domain_test

import (
	"encoder/domain"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestJobIdIsEmpty(t *testing.T) {
	video, err := domain.NewVideo("", "123", "some/path")
	require.Nil(t, err)

	job, err := domain.NewJob("", domain.STARTING, video)
	require.Nil(t, err)
	require.NotEmpty(t, job.Id)
	require.Equal(t, domain.STARTING, job.Status)
}

func TestVideoIdIsNotUUID(t *testing.T) {
	video, err := domain.NewVideo("", "123", "some/path")
	require.Nil(t, err)

	_, err = domain.NewJob("123", domain.STARTING, video)
	require.Error(t, err)
}

func TestExistingVideoId(t *testing.T) {
	video, err := domain.NewVideo("", "123", "some/path")
	require.Nil(t, err)

	id := uuid.NewV4().String()
	job, err := domain.NewJob(id, domain.STARTING, video)
	require.Nil(t, err)
	require.Equal(t, id, job.Id)
}