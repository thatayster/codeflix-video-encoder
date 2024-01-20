package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func generateDomainResources() (*domain.Video, *domain.Job) {
	resourceId := "resource-1"
	filePath := "some/path"
	videoId := uuid.NewV4().String()
	domainVideo, _ := domain.NewVideo(videoId, resourceId, filePath)

	jobId := uuid.NewV4().String()
	jobDomain, _ := domain.NewJob(jobId, domain.STARTING, domainVideo)

	return domainVideo, jobDomain
}

func TestVideoConversionWithoutJobs(t *testing.T) {
	video, _ := generateDomainResources()

	videoOrm, err := repositories.ToVideoORM(video)

	require.Nil(t, err)
	require.Equal(t, video.Id, videoOrm.ID)
	require.Empty(t, videoOrm.Jobs)

	videoDomain, err := videoOrm.ToVideoDomain()

	require.Nil(t, err)
	require.Equal(t, videoOrm.ID, videoDomain.Id)
	require.EqualValues(t, video, videoDomain)
}

func TestVideoConversionWithJobs(t *testing.T) {
	// Create original domain resources
	video, job1 := generateDomainResources()
	job2, _ := domain.NewJob("", domain.COMPLETED, video)

	video.Jobs = append(video.Jobs, job1)
	video.Jobs = append(video.Jobs, job2)

	// Translate them into ORM models
	videoOrm, err := repositories.ToVideoORM(video)

	require.Nil(t, err)
	require.Len(t, videoOrm.Jobs, 2)

	returnedJobOrm1 := videoOrm.Jobs[0]
	returnedJobOrm2 := videoOrm.Jobs[1]

	require.Equal(t, job1.Id, returnedJobOrm1.ID)
	require.Equal(t, job2.Id, returnedJobOrm2.ID)

	// Restore them to domain Entities
	videoDomain, err := videoOrm.ToVideoDomain()

	require.Nil(t, err)
	require.Equal(t, video.Id, videoDomain.Id)

	returnedJobDomain1 := videoDomain.Jobs[0]
	returnedJobDomain2 := videoDomain.Jobs[1]

	require.Equal(t, job1.Id, returnedJobDomain1.Id)
	require.Equal(t, job2.Id, returnedJobDomain2.Id)
}
