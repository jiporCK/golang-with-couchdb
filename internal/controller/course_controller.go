package controller

import (
	"cmd/internal/entity"
	"cmd/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)


type CourseController struct {
	service *usecase.CourseService
}

func NewCourseController(s *usecase.CourseService) *CourseController {
	return &CourseController{service: s}
}

func (c *CourseController) CreateCourse(ctx *gin.Context) {
	var course entity.Course
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateCourse(course); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Course created successfully"})
}

func (c *CourseController) GetAllCourses(ctx *gin.Context) {
	courses, err := c.service.GetAllCourses()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}
	ctx.JSON(http.StatusOK, courses)
}

