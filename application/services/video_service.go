package services

import (
	"encoder/application/repositories"
	"encoder/domain"
	"io"
	"log"
	"os/exec"

	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/credentials"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/s3"
	// "github.com/aws/aws-sdk-go/service/s3/s3manager"

	"os"
)

type VideoService struct {
	Video *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

func (v *VideoService) Download(bucketName string) error {
    // awsRegion := "us-east-1"
	// bucketName := "dev-ifood-ml-sagemaker"
	// objectKey := v.Video.FilePath 
	// localFilePath := os.Getenv("LOCAL_STORAGE_PATH") + "/" + v.Video.Id + ".mp4"

    // // Load AWS credentials from the /root/.aws/credentials file
	// creds := credentials.NewSharedCredentials("/root/.aws/credentials", "default")

    // // Create an AWS session with the loaded credentials
	// sess, err := session.NewSession(&aws.Config{
	// 	Region:      aws.String(awsRegion),
	// 	Credentials: creds,
	// })
	// if err != nil {
	// 	log.Fatal("Error creating AWS session:", err)
	// 	return err
	// }

    // // Create an S3 client using the session
	// client := s3.New(sess)

	// // Create a file to write the S3 object data to
	// file, err := os.Create(localFilePath)
	// if err != nil {
	// 	log.Fatal("Error creating file:", err)
	// 	return err
	// }
	// defer file.Close()

    // // Create a downloader using the S3 client
	// downloader := s3manager.NewDownloaderWithClient(client)

	// // Download the file from S3
	// _, err = downloader.Download(file, &s3.GetObjectInput{
	// 	Bucket: aws.String(bucketName),
	// 	Key:    aws.String(objectKey),
	// })
	// if err != nil {
	// 	log.Fatal("Error downloading file from S3:", err)
	// 	return err
	// }
	// log.Println("File downloaded successfully!")	
    // return nil

	sourceFile, err := os.Open(bucketName + "/" + v.Video.FilePath)
	if err != nil {
		return err
	}
	destinationFile, err := os.Create(os.Getenv("LOCAL_STORAGE_PATH") + "/" + v.Video.Id + ".mp4")
	if err != nil {
		return err
	}
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}
	return nil
}

func (v *VideoService) Fragment() error {
	err := os.Mkdir(os.Getenv("LOCAL_STORAGE_PATH") + "/" + v.Video.Id, os.ModePerm)
	if err != nil {
		log.Fatal("Could not create temporary storage directory", err)
		return err
	}
	source := os.Getenv("LOCAL_STORAGE_PATH") + "/" + v.Video.Id + ".mp4"
	target := os.Getenv("LOCAL_STORAGE_PATH") + "/" + v.Video.Id + ".frag"

	cmd := exec.Command("mp4fragment", source, target)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	printOutput(output)
	return nil
}

func (v *VideoService) Encode() error {
	cmdArgs := []string{}
	cmdArgs = append(cmdArgs, os.Getenv("LOCAL_STORAGE_PATH") + "/" + v.Video.Id + ".frag")
	cmdArgs = append(cmdArgs, "--use-segment-timeline")
	cmdArgs = append(cmdArgs, "-o")
	cmdArgs = append(cmdArgs, os.Getenv("LOCAL_STORAGE_PATH") + "/" + v.Video.Id)
	cmdArgs = append(cmdArgs, "-f")	

	cmd := exec.Command("mp4dash", cmdArgs...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Println("\n\nError:", err)
    	log.Println("\nStderr:", cmd.Stderr)
		return err
	}
	log.Println("\n\n\nOutput:", cmd.Stdout)
	printOutput(output)
	return nil
}

func (v *VideoService) Finish() error {
	local_path := os.Getenv("LOCAL_STORAGE_PATH")

	err := os.Remove(local_path + "/" + v.Video.Id + ".mp4")
	if err != nil {
		return err
	}
	err = os.Remove(local_path + "/" + v.Video.Id + ".frag")
	if err != nil {
		return err
	}
	err = os.RemoveAll(local_path + "/" + v.Video.Id)
	if err != nil {
		return err
	}

	log.Println("files have been removed: ", v.Video.Id)
	return nil
}

func (v *VideoService) InsertVideo() error {
	_, err := v.VideoRepository.Insert(v.Video)
	if err != nil {
		return err
	}
	return nil
}

func printOutput(out []byte) {
	if len(out) > 0 {
		log.Printf("=======> Output: %s\n", string(out))
	}
}
