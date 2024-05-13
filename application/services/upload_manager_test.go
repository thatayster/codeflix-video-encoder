package services_test

import (
	"encoder/application/services"
	"encoder/domain"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("error loading .env file: ", err)
	}
}

func TestVideoServiceUpload(t *testing.T) {

	encoded_video_path := "../../resources/encoded_files"

	videoUpload := services.NewVideoUpload()
	videoUpload.OutputBucket = "bucket-name" // it must be removed from here to the cloud repository
	videoUpload.VideoPath = encoded_video_path

	doneUpload := make(chan string)
	go videoUpload.ProcessUpload(20, doneUpload)

	result := <- doneUpload
	require.Empty(t, videoUpload.Errors)
	require.Equal(t, result, domain.COMPLETED.ToString())
}
