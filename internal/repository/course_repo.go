package repository

import (
	"cmd/internal/database"
	"cmd/internal/entity"
	"context"
	"log"

	"github.com/go-kivik/kivik/v3"
	"github.com/google/uuid"
)

type CourseRepo struct{}

func (r *CourseRepo) CreateCourse(course entity.Course) error {
	db := database.GetDB("courses")

	if course.ID == "" {
		course.ID = uuid.New().String()
	}
	_, err := db.Put(context.TODO(), course.ID, course)
	if err != nil {
		log.Println("Failed to create course: ", err)
		return err
	}
	return nil
}

func (r *CourseRepo) GetAllCourses() ([]entity.Course, error) {
	db := database.GetDB("courses")
	rows, err := db.AllDocs(context.TODO(), kivik.Options{"include_docs": true})

	if err != nil {
		log.Println("Failed to retrieve course", err)
		return nil, err
	}

	var courses []entity.Course
	for rows.Next() {
		var course entity.Course
		if err := rows.ScanDoc(&course); err != nil {
			log.Println("Failed to scan course: ", err)
			continue
		}
		courses = append(courses, course)
	}
	return courses, nil
}
