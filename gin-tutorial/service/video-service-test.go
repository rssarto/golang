package service

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rssarto/gin-tutorial/entity"
	"github.com/rssarto/gin-tutorial/repository"
)

const (
	TITLE       = "Video Title"
	DESCRIPTION = "Video Description"
	URL         = "https://www.youtube.com/watch?v=t4MBGpEFHyk&list=PL3eAkoh7fypr8zrkiygiY1e9osoqjoV9w&index=12"
	FIRSTNAME   = "John"
	LASTNAME    = "Doe"
	EMAIL       = "jdoe@gmail.com"
)

var testVideo entity.Video = entity.Video{
	Title:       TITLE,
	Description: DESCRIPTION,
	URL:         URL,
	Author: entity.Person{
		FirstName: FIRSTNAME,
		LastName:  LASTNAME,
		Email:     EMAIL,
	},
}

var _ = Describe("Video Service", func() {

	var (
		videoRepository repository.VideoRepository
		videoService    VideoService
	)

	BeforeSuite(func() {
		videoRepository = repository.NewVideoRepository()
		videoService = New(videoRepository)
	})

	Describe("Fetching all existing videos", func() {
		Context("If there is a video in the database", func() {
			BeforeEach(func() {
				videoList := videoService.FindAll()
				Ω(videoList).ShouldNot(BeEmpty())
			})
		})
	})

})
