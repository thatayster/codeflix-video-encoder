package repositories

import (
	"encoder/domain"
	"time"
)

type Video struct {
	ID         string    `gorm:"type:uuid;primary_key"`
	ResourceID string    `gorm:"type:varchar(255)"`
	FilePath   string    `gorm:"type:varchar(255)"`
	Jobs       []*Job    `gorm:"ForeignKey:VideoID"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func ToVideoORM(video *domain.Video) (*Video, error) {
	err := video.Validate()
	if err != nil {
		return nil, err
	}

	videoORM := &Video{
		ID:         video.Id,
		ResourceID: video.ResourceId,
		FilePath:   video.FilePath,
	}

	// Filling in the jobs
	for index := range video.Jobs {
		jobDomain := video.Jobs[index]
		jobOrm := &Job{
			ID:     jobDomain.Id,
			Status: string(jobDomain.Status),
			Error:  jobDomain.Error,
		}
		videoORM.Jobs = append(videoORM.Jobs, jobOrm)
	}
	return videoORM, nil
}

func (video *Video) ToVideoDomain() (*domain.Video, error) {
	videoDomain, err := domain.NewVideo(video.ID, video.ResourceID, video.FilePath)
	if err != nil {
		return nil, err
	}
	// Filling in the jobs
	for index := range video.Jobs {
		jobOrm := video.Jobs[index]

		status, err := domain.OperationStatusFromString(jobOrm.Status)
		if err != nil {
			return nil, err
		}

		jobDomain, err := domain.NewJob(jobOrm.ID, status, videoDomain)
		if err != nil {
			return nil, err
		}
		videoDomain.Jobs = append(videoDomain.Jobs, jobDomain)
	}
	return videoDomain, nil
}
