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

func (s *CourseService) GetCourseById(id string) (*entity.Course, error) {
	return s.repo.GetCourseById(id)
}

func (s *CourseService) UpdateCourseById(id string, rev string, updatedCourse entity.Course) (*entity.Course, error) {
	return s.repo.UpdateCourseById(id, rev, updatedCourse)
}

func (s *CourseService) DeleteCourseById(id string, rev string) error {
	return s.repo.DeleteCourseById(id, rev)
}