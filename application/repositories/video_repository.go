package repositories

import (
	"encoder/domain"
	"fmt"

	"github.com/jinzhu/gorm"
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
	videoOrm, err := ToVideoORM(video)
	if err != nil {
		return nil, err
	}
	err = repo.Db.Create(videoOrm).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (repo VideoRepositoryDB) Find(id string) (*domain.Video, error) {
	var videoOrm Video
	repo.Db.Preload("Jobs").First(&videoOrm, "id = ?", id)

	if videoOrm.ID == "" {
		return nil, fmt.Errorf("video does not exist")
	}

	videoDomain, err := videoOrm.ToVideoDomain()
	if err != nil {
		return nil, err
	}
	return videoDomain, nil
}
