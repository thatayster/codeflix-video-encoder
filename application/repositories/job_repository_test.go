package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJobRepositoryInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video, job := generateDomainResources()
	repoVideo := repositories.VideoRepositoryDB{Db:db}
	repoJob := repositories.JobRepositoryDB{Db:db}

	// 'Insert' for Video Repository is tested in another file
	repoVideo.Insert(video)
	_, err := repoJob.Insert(job)

	require.Nil(t, err)

	foundJob, err := repoJob.Find(job.Id)

	require.Nil(t, err)
	require.Equal(t, job.Id, foundJob.Id)
	require.Equal(t, job.Status, foundJob.Status)
	require.EqualValues(t, job.Video, foundJob.Video)
}

func TestJobRepositoryUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video, job := generateDomainResources()

	repo := repositories.VideoRepositoryDB{Db:db}
	repoJob := repositories.JobRepositoryDB{Db:db}

	repo.Insert(video)	
	_, err := repoJob.Insert(job)

	require.Nil(t, err)

	job.Status = domain.COMPLETED

	_, err = repoJob.Update(job)

	require.Nil(t, err)

	foundJob, err := repoJob.Find(job.Id)

	require.Nil(t, err)
	require.Equal(t, job.Status, foundJob.Status)
}

func TestJobNotFound(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	repo := repositories.JobRepositoryDB{Db:db}

	_, err := repo.Find("non-existent-id")

	require.Error(t, err)
}