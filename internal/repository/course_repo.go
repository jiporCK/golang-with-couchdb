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

func (r *CourseRepo) GetCourseById(id string) (*entity.Course, error) {
    db := database.GetDB("courses")
    
    row := db.Get(context.TODO(), id)

	if row.Err != nil {
		log.Println("Failed to retrieve course: ", row.Err)
		return nil, row.Err
	}

	var course entity.Course
	if err :=  row.ScanDoc(&course); err != nil {
		log.Println("Failed to scan course document: ", err)
		return nil, err
	}

	return &course, nil
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

func (r *CourseRepo) UpdateCourseById(id string, rev string, updatedCourse entity.Course) (*entity.Course, error) {
    db := database.GetDB("courses")
    
    row := db.Get(context.TODO(), id)
    if row.Err != nil {
        log.Println("Failed to retrieve course: ", row.Err)
        return nil, row.Err
    }

    var existingCourse entity.Course

    if err := row.ScanDoc(&existingCourse); err != nil {
        log.Println("Failed to scan course document: ", err)
        return nil, err
    }

    existingCourse.Name = updatedCourse.Name
    existingCourse.TeacherID = updatedCourse.TeacherID

    existingCourse.Rev = rev

    _, err := db.Put(context.TODO(), id, existingCourse)
    if err != nil {
        log.Println("Failed to update course: ", err)
        return nil, err
    }

    return &existingCourse, nil
}

func (r *CourseRepo) DeleteCourseById(id string, rev string) error {
	db := database.GetDB("courses")

	_, err := db.Delete(context.TODO(), id, rev)
	if err != nil {
		log.Println("Failed to delete course: ", err)
		return err
	}
	return nil
}


