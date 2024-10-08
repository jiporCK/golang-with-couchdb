package routes

import (
	"cmd/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(courseController *controller.CourseController) *gin.Engine {

	r := gin.Default()

	courseRouter := r.Group("/courses") 
	{
		courseRouter.POST("", courseController.CreateCourse)
		courseRouter.GET("", courseController.GetAllCourses)
		courseRouter.GET("/:_id", courseController.GetCourseById)
		courseRouter.PUT("/:_id", courseController.UpdateCourseById)
		courseRouter.DELETE("/:_id", courseController.DeleteCourseById)
	}

	return r

}