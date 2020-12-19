package service

import "go-gin-app/controller"

type VideoService interface {
	Save(controller.Video) controller.Video
	FindAll() []controller.Video
}
