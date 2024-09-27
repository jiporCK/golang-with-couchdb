package main

import (
	"cmd/internal/controller"
	"cmd/internal/database"
	"cmd/internal/repository"
	"cmd/internal/usecase"
	"cmd/routes"
)


func main() {
	database.InitCouchDB()

	courseRepo := &repository.CourseRepo{}
	courseService := usecase.NewCourseService(courseRepo)
	courseController := controller.NewCourseController(courseService)

	r := routes.InitRoutes(courseController)

	r.Run(":8080")

}