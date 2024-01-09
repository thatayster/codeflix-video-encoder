package services

import (
	"encoder/domain"
	"encoder/framework/utils"
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

type JobWorkerResult struct {
	Job domain.Job
	Message amqp.Delivery
	Error error
}

var Mutex = &sync.Mutex{}

func returnJobResult(job domain.Job, message amqp.Delivery, err error) JobWorkerResult {
	result := JobWorkerResult{
		Job: job,
		Message: message,
		Error: err,
	}
	return result
}

func JobWorker(messageChannel chan amqp.Delivery, returnChan chan JobWorkerResult, jobService JobService, job domain.Job, workerId int) {
	for message := range messageChannel {
		log.Println("Received message")

		err := utils.IsJson(string(message.Body))
		if err != nil {
			returnChan <- returnJobResult(domain.Job{}, message, err)
			continue
		}

		Mutex.Lock()
		err = json.Unmarshal(message.Body, &jobService.VideoService.Video)
		jobService.VideoService.Video.Id = uuid.NewV4().String()
		Mutex.Unlock()

		if err != nil {
			returnChan <- returnJobResult(domain.Job{}, message, err)
			continue
		}

		log.Printf("Received a valid message body")

		err = jobService.VideoService.Video.Validate()
		if err != nil {
			returnChan <- returnJobResult(domain.Job{}, message, err)
			continue
		}

		log.Println("Video successfully validated.")

		Mutex.Lock()
		err = jobService.VideoService.InsertVideo()
		Mutex.Unlock()

		if err != nil {
			returnChan <- returnJobResult(domain.Job{}, message, err)
			continue
		}

		log.Println("Video successfully stored in the database.")

		job.Video = jobService.VideoService.Video
		job.OutputBucketPath = os.Getenv("OUTPUT_BUCKET_NAME")
		job.Id = uuid.NewV4().String()
		job.Status = "STARTING"
		job.CreatedAt = time.Now()

		Mutex.Lock()
		_, err = jobService.JobRepository.Insert(&job)
		Mutex.Unlock()
		
		if err != nil {
			returnChan <- returnJobResult(domain.Job{}, message, err)
			continue
		}

		log.Println("Job successfully stored in the database.")

		jobService.Job = &job
		err = jobService.Start()
		if err != nil {
			returnChan <- returnJobResult(domain.Job{}, message, err)
			continue
		}
		returnChan <- returnJobResult(job, message, nil)
	}	
}
