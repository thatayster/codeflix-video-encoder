package services_test

import (
	"encoder/application/repositories"
	"encoder/application/services"
	"encoder/domain"
	"encoder/framework/database"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("error loading .env file: ", err)
	}
}	

func prepare() (*domain.Video, repositories.VideoRepositoryDB) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.Id = uuid.NewV4().String()
	//video.FilePath = "tbs_tests/input/data/test/test.mp4"
	video.FilePath = "test.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDB{Db:db}
	return video, repo
}

func TestVideoServiceDownload(t *testing.T) {

	video, repo := prepare()
	videoService := services.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download("../../" + os.Getenv("INPUT_BUCKET_NAME"))
	require.Nil(t, err)
	require.FileExists(t, os.Getenv("LOCAL_STORAGE_PATH") + "/" + video.Id + ".mp4")

	err = videoService.Fragment()
	require.Nil(t, err)	
	require.FileExists(t, os.Getenv("LOCAL_STORAGE_PATH") + "/" + video.Id + ".frag")

	err = videoService.Encode()
	require.Nil(t, err)
	require.DirExists(t, os.Getenv("LOCAL_STORAGE_PATH") + "/" + video.Id)

	err = videoService.Finish()
	require.Nil(t, err)

	require.NoDirExists(t, os.Getenv("LOCAL_STORAGE_PATH") + "/" + video.Id)
	require.NoFileExists(t, os.Getenv("LOCAL_STORAGE_PATH") + "/" + video.Id + ".mp4")
	require.NoFileExists(t, os.Getenv("LOCAL_STORAGE_PATH") + "/" + video.Id + ".frag")
}