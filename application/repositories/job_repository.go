package repositories

import (
	"encoder/domain"
	"fmt"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type JobRepository interface {
	Insert(job *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
	Update(job *domain.Job) (*domain.Job, error)
}

type JobRepositoryDB struct {
	Db *gorm.DB
}

func NewJobRepositoryDB(db *gorm.DB) *JobRepositoryDB {
	return &JobRepositoryDB{Db: db}
}

func (repo JobRepositoryDB) Insert(job *domain.Job) (*domain.Job, error) {
	if job.Id == "" {
		job.Id = uuid.NewV4().String()
	}

	err := repo.Db.Create(job).Error
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (repo JobRepositoryDB) Find(id string) (*domain.Job, error) {
	var job domain.Job
	repo.Db.Preload("Video").First(&job, "id = ?", id)

	if job.Id == "" {
		return nil, fmt.Errorf("video does not exist")
	}
	return &job, nil
}

func (repo JobRepositoryDB) Update(job *domain.Job) (*domain.Job, error) {
	err := repo.Db.Save(&job).Error

	if err != nil {
		return nil, err
	}
	return job, nil
}