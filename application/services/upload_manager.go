package services

import (
	"encoder/domain"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type VideoUpload struct {
	Paths        []string
	VideoPath    string
	OutputBucket string      // it should be an abstraction (repository), not a bucket name
	Errors       []string
}

func NewVideoUpload() *VideoUpload {
	return &VideoUpload{}
}

func (vu *VideoUpload) UploadObject(objectPath string) error {
	// objectPath has the format local_directory/video_id/file_name
	// but we want to retrieve only the video_id/file name
	//path := strings.Split(objectPath, os.Getenv("LOCAL_STORAGE_PATH") + "/")
	//log.Println("Uploading file: ", objectPath)
	// find a way to test it without having to upload the files to the cloud provider
	return nil
}

func (vu *VideoUpload) LoadPaths() error {
	// TODO: update the Go verstion to +1.16 to use a more efficient function WalkDir instead
	err := filepath.Walk(vu.VideoPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			vu.Paths = append(vu.Paths, path)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func getClientUpload() error {
	// the original function returns an storage_client, a context and an error
	// it is used to avoid creating a new storage client every time a upload must be performed
	// a possible solution to my future repository is implementing a singleton pattern
	return nil
}

func (vu *VideoUpload) ProcessUpload(concurrency int, doneUpload chan string) error {
	// responsible for managing the upload process itself
	in := make(chan int, runtime.NumCPU())
	returnChannel := make(chan string)

	err := vu.LoadPaths()
	if err != nil {
		return err
	}

	err = getClientUpload()
	if err != nil {
		return err
	}

	// create the parallel go routines
	for process := 0; process < concurrency; process++ {
		go vu.UploadWorker(in, returnChannel)
	}

	go func() {
		for x := 0; x < len(vu.Paths); x++ {
			in <- x
		}
		close(in)
	}()

	for r := range returnChannel {
		if r != domain.SUCCEEDED.ToString() {
			doneUpload <- r
			break
		}
	}
	return nil
}

// originally also receives the storage client and the context
func (vu *VideoUpload) UploadWorker(in chan int, returnChan chan string) {
	// go routine that is responsible for calling the UploadObject function
	// its name must be changed
	for x := range in {
		err := vu.UploadObject(vu.Paths[x])

		if err != nil {
			vu.Errors = append(vu.Errors, vu.Paths[x])
			log.Printf("error during the upload: %v. Error: %v", vu.Paths[x], err)
			returnChan <- err.Error()
		}
		returnChan <- domain.SUCCEEDED.ToString()
	}
	returnChan <- domain.COMPLETED.ToString()
}
