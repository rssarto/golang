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
				videoService.Save(testVideo)
			})

			It("should return at least one element", func() {
				videoList := videoService.FindAll()
				Ω(videoList).ShouldNot(BeEmpty())
			})

			It("should map the fields correctly", func() {
				firstVideo := videoService.FindAll()[0]

				Ω(firstVideo.Title).Should(Equal(TITLE))
				Ω(firstVideo.Description).Should(Equal(DESCRIPTION))
				Ω(firstVideo.URL).Should(Equal(URL))
				Ω(firstVideo.Author.FirstName).Should(Equal(FIRSTNAME))
				Ω(firstVideo.Author.LastName).Should(Equal(LASTNAME))
				Ω(firstVideo.Author.Email).Should(Equal(EMAIL))
			})

			AfterEach(func() {
				video := videoService.FindAll()[0]
				videoService.Delete(video)
			})
		})

		Context("If there are no videos in the database", func() {
			It("should return an empty list", func() {
				videos := videoService.FindAll()

				Ω(videos).Should(BeEmpty())
			})
		})
	})

	AfterSuite(func() {

	})
})
