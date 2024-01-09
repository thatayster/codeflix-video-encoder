package domain_test

import (
	"encoder/domain"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNotUUID(t *testing.T) {
	video := domain.NewVideo()
	video.Id = "abc"
	video.ResourceId = "a"
	video.FilePath = "/some/file"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)

}

func TestVideoIsValid(t *testing.T) {
	video := domain.NewVideo()
	video.Id = uuid.NewV4().String()
	video.ResourceId = "a"
	video.FilePath = "/some/file"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)

}