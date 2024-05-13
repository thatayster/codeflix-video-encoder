package repositories

import (
	"encoder/domain"
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Job struct {
	ID        string `gorm:"type:uuid;primary_key"`
	VideoID   string `gorm:"column:video_id;type:uuid REFERENCES videos(id);notnull"`
	Status    string
	Video     *Video
	Error     string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func ToJobORM(job *domain.Job) (*Job, error) {
	err := job.Validate()
	if err != nil {
		return nil, err
	}

	// Convert video domain to video ORM before adding it to the job ORM
	videoORM := &Video{
		ID:         job.Video.Id,
		ResourceID: job.Video.ResourceId,
		FilePath:   job.Video.FilePath,
	}
	jobORM := &Job{
		ID:      job.Id,
		VideoID: job.Video.Id,
		Video:   videoORM,
		Status:  string(job.Status),
		Error:   job.Error,
	}
	return jobORM, nil
}

func (job *Job) ToJobDomain() (*domain.Job, error) {
	status, err := domain.OperationStatusFromString(job.Status)
	if err != nil {
		return nil, err
	}
	videoDomain, err := domain.NewVideo(job.Video.ID, job.Video.ResourceID, job.Video.FilePath)
	if err != nil {
		return nil, err
	}
	jobDomain, err := domain.NewJob(job.ID, status, videoDomain)
	if err != nil {
		return nil, err
	}
	return jobDomain, nil
}
