package repositories

import (
	"encoder/domain"
	"fmt"

	"github.com/jinzhu/gorm"
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
	jobOrm, err := ToJobORM(job)
	if err != nil {
		return nil, err
	}

	err = repo.Db.Create(jobOrm).Error
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (repo JobRepositoryDB) Find(id string) (*domain.Job, error) {
	var jobOrm Job
	repo.Db.Preload("Video").First(&jobOrm, "id = ?", id)

	if jobOrm.ID == "" {
		return nil, fmt.Errorf("video does not exist")
	}

	jobDomain, err := jobOrm.ToJobDomain()
	if err != nil {
		return nil, err
	}
	return jobDomain, nil
}

func (repo JobRepositoryDB) Update(job *domain.Job) (*domain.Job, error) {
	jobOrm, err := ToJobORM(job)
	if err != nil {
		return nil, err
	}
	err = repo.Db.Save(&jobOrm).Error
	if err != nil {
		return nil, err
	}
	return job, nil
}
