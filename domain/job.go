package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Job struct {
	Id               string          `json:"job_id" valid:"uuid"`
	Status           OperationStatus `json:"status" valid:"notnull"`
	Video            *Video          `json:"video" valid:"required"`
	Error            string          `valid:"-"`
}

func NewJob(id string, status OperationStatus, video *Video) (*Job, error) {
	if id == "" {
		id = uuid.NewV4().String()
	}
	job := Job{
		Id:               id,
		Status:           status,
		Video:            video,
	}
	err := job.Validate()

	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (job *Job) Validate() error {
	_, err := govalidator.ValidateStruct(job)

	if err != nil {
		return err
	}
	return nil
}
