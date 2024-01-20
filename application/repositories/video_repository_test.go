package repositories_test

import (
 	"encoder/application/repositories"
 	"encoder/framework/database"
 	"testing"

 	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video, _ := generateDomainResources()

	repo := repositories.VideoRepositoryDB{Db:db}
	_, err := repo.Insert(video)

	require.Nil(t, err)

	foundVideo, err := repo.Find(video.Id)

	require.Nil(t, err)
	require.NotEmpty(t, foundVideo.Id)
	require.Equal(t, foundVideo.Id, video.Id)
}

func TestVideoNotFound(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	repo := repositories.VideoRepositoryDB{ Db:db }
	_, err := repo.Find("non-existent-id")

	require.Error(t, err)
}