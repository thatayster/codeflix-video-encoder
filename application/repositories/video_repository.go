package repositories

import (
	"encoder/domain"
	"fmt"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type VideoRepository interface {
	Insert(video *domain.Video) (*domain.Video, error)
	Find(id string) (*domain.Video, error)
}

type VideoRepositoryDB struct {
	Db *gorm.DB
}

func NewVideoRepositoryDB(db *gorm.DB) *VideoRepositoryDB {
	return &VideoRepositoryDB{Db: db}
}

func (repo VideoRepositoryDB) Insert(video *domain.Video) (*domain.Video, error) {
	if video.Id == "" {
		video.Id = uuid.NewV4().String()
	}

	err := repo.Db.Create(video).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (repo VideoRepositoryDB) Find(id string) (*domain.Video, error) {
	var video domain.Video
	repo.Db.Preload("Jobs").First(&video, "id = ?", id)

	if video.Id == "" {
		return nil, fmt.Errorf("video does not exist")
	}
	return &video, nil
}