package domain_test

import (
	"encoder/domain"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestIdIsEmpty(t *testing.T) {
	resourceId := "resource-1"
	filePath := "some/path"
	video, err := domain.NewVideo("", resourceId, filePath)

	require.Nil(t, err)
	require.Equal(t, resourceId, video.ResourceId)
	require.Equal(t, filePath, video.FilePath)
}

func TestIdIsNotUUID(t *testing.T) {
	resourceId := "resource-1"
	filePath := "some/path"
	_, err := domain.NewVideo("123", resourceId, filePath)

	require.Error(t, err)
}

func TestExistingId(t *testing.T) {
	resourceId := "resource-1"
	filePath := "some/path"
	id := uuid.NewV4().String()
	video, err := domain.NewVideo(id, resourceId, filePath)

	require.Nil(t, err)
	require.Equal(t, id, video.Id)
}
