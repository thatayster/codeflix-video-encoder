package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestJobRepositoryInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.Id = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDB{Db:db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDB{Db:db}
	repoJob.Insert(job)

	j, err := repoJob.Find(job.Id)
	require.NotEmpty(t, j.Id)
	require.Nil(t, err)
	require.Equal(t, job.Id, j.Id)
	require.Equal(t, j.VideoId, video.Id)

	foundVideo, err := repo.Find(j.VideoId)

	require.Nil(t, err)
	require.Equal(t, video.Id, foundVideo.Id)
	require.NotEmpty(t, foundVideo.Jobs)
}

func TestJobRepositoryUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.Id = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDB{Db:db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDB{Db:db}
	repoJob.Insert(job)

	job.Status = "Completed"

	repoJob.Update(job)

	j, err := repoJob.Find(job.Id)
	require.NotEmpty(t, j.Id)
	require.Nil(t, err)
	require.Equal(t, j.Status, job.Status)
}