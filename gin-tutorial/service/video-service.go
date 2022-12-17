package service

import (
	"github.com/rssarto/gin-tutorial/entity"
	"github.com/rssarto/gin-tutorial/repository"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
}

type videoService struct {
	videoRepository repository.VideoRepository
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videoRepository.Save(video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()
}

func (service *videoService) Update(video entity.Video) {
	service.videoRepository.Update(video)
}

func (service *videoService) Delete(video entity.Video) {
	service.videoRepository.Delete(video)
}

func New(videoRepository repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: videoRepository,
	}
}
