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

func (c *CourseController) GetCourseById(ctx *gin.Context) {
	id := ctx.Param("_id")

	course, err := c.service.GetCourseById(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Course not found"})
		return
	}

	ctx.JSON(200, course)
}

func (c *CourseController) UpdateCourseById(ctx *gin.Context) {
    id := ctx.Param("_id")
    existingCourse, err := c.service.GetCourseById(id)
    if err != nil {
        ctx.JSON(404, gin.H{"error": "Course not found"})
        return
    }

    var updatedCourse entity.Course
    if err := ctx.ShouldBindJSON(&updatedCourse); err != nil {
        ctx.JSON(400, gin.H{"error": "Invalid input"})
        return
    }

    updateCourse, err := c.service.UpdateCourseById(id, existingCourse.Rev, updatedCourse)
    if err != nil {
        ctx.JSON(500, gin.H{"error": "Failed to update course"})
        return
    }

    ctx.JSON(200, updateCourse)
}

func (c *CourseController) DeleteCourseById(ctx *gin.Context) {

	id := ctx.Param("_id")
	existingCourse, err := c.service.GetCourseById(id)
    if err != nil {
        ctx.JSON(404, gin.H{"error": "Course not found"})
        return
    }
	c.service.DeleteCourseById(id, existingCourse.Rev)
	ctx.JSON(204, gin.H{"message": "Course deleted successfully"})

}



