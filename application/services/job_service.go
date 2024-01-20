package services

import (
	"encoder/application/repositories"
	"encoder/domain"
	"errors"
	"os"
	"strconv"
)

type JobService struct {
	Job           *domain.Job
	JobRepository repositories.JobRepository
	VideoService  VideoService
}

func (j *JobService) Start() error {
	err := j.changeJobStatus(domain.DOWNLOADING)
	if err != nil {
		return j.failJob(err)
	}
	err = j.VideoService.Download(os.Getenv("INPUT_BUCKET_NAME"))
	if err != nil {
		return j.failJob(err)
	}

	err = j.changeJobStatus(domain.FRAGMENTING)
	if err != nil {
		return j.failJob(err)
	}
	err = j.VideoService.Fragment()
	if err != nil {
		return j.failJob(err)
	}

	err = j.changeJobStatus(domain.ENCODING)
	if err != nil {
		return j.failJob(err)
	}
	err = j.VideoService.Encode()
	if err != nil {
		return j.failJob(err)
	}

	err = j.changeJobStatus(domain.UPLOADING)
	if err != nil {
		return j.failJob(err)
	}
	err = j.performUpload()
	if err != nil {
		return j.failJob(err)
	}

	err = j.changeJobStatus(domain.FINISHNING)
	if err != nil {
		return j.failJob(err)
	}
	err = j.VideoService.Finish()
	if err != nil {
		return j.failJob(err)
	}

	err = j.changeJobStatus(domain.COMPLETED)
	if err != nil {
		return j.failJob(err)
	}

	return nil
}

func (j *JobService) performUpload() error {
	videoUpload := NewVideoUpload()
	videoUpload.OutputBucket = os.Getenv("OUTPUT_BUCKET_NAME")
	videoUpload.VideoPath = os.Getenv("LOCAL_STORAGE_PATH") + "/" + j.VideoService.Video.Id
	concurrency, _ := strconv.Atoi(os.Getenv("CONCURRENCY_UPLOAD"))
	doneUpload := make(chan string)
	go videoUpload.ProcessUpload(concurrency, doneUpload)

	var uploadResult string = <-doneUpload

	if uploadResult != domain.COMPLETED.ToString() {
		return j.failJob(errors.New(uploadResult))
	}
	return nil
}

func (j *JobService) changeJobStatus(status domain.OperationStatus) error {
	var err error

	j.Job.Status = status
	j.Job, err = j.JobRepository.Update(j.Job)

	if err != nil {
		return j.failJob(err)
	}
	return nil
}

func (j *JobService) failJob(err error) error {
	j.Job.Error = err.Error()
	_, update_err := j.JobRepository.Update(j.Job)
	if update_err != nil {
		return update_err
	}
	return err
}
