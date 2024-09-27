package usecase

import (
	"cmd/internal/entity"
	"cmd/internal/repository"
)

type CourseService struct {
	repo *repository.CourseRepo
}

func NewCourseService(repo *repository.CourseRepo) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) CreateCourse(course entity.Course) error {
	return s.repo.CreateCourse(course)
}

func (s *CourseService) GetAllCourses() ([]entity.Course, error) {
	return s.repo.GetAllCourses()
}