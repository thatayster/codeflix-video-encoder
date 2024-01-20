package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Video struct {
	Id         string `json:"encoded_video_folder" valid:"uuid"`
	ResourceId string `json:"resource_id" valid:"notnull"`
	FilePath   string `json:"file_path" valid:"notnull"`
	Jobs       []*Job `json:"-" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo(id string, resourceId string, filePath string) (*Video, error) {
	if id == "" {
		id = uuid.NewV4().String()
	}
	video := Video{
		Id:         id,
		ResourceId: resourceId,
		FilePath:   filePath,
		Jobs:       []*Job{},
	}
	err := video.Validate()

	if err != nil {
		return nil, err
	}
	return &video, nil
}

func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video)
	if err != nil {
		return err
	}
	return nil
}
