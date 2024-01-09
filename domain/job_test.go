package domain_test

import (
	"encoder/domain"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	video := domain.NewVideo()
	video.Id = uuid.NewV4().String()
	video.FilePath = "some/file"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("some/path", "Converted", video)

	require.NotNil(t, job)
	require.Nil(t, err)
}