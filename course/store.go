package course

import (
	"database/sql"
	"fmt"

	"github.com/AhmadMirza2023/krs/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetCourses() ([]types.Course, error) {
	rows, err := s.db.Query("SELECT * FROM courses")
	if err != nil {
		return nil, err
	}
	courses := make([]types.Course, 0)
	for rows.Next() {
		course, err := scanRowIntoCourse(rows)
		if err != nil {
			return nil, err
		}
		courses = append(courses, *course)
	}
	return courses, nil
}

func (s *Store) GetCourseById(id int) (*types.Course, error) {
	rows, err := s.db.Query("SELECT * FROM courses WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	course := new(types.Course)
	for rows.Next() {
		course, err = scanRowIntoCourse(rows)
		if err != nil {
			return nil, err
		}
	}
	if course.Id == 0 {
		return nil, fmt.Errorf("course not found")
	}
	return course, nil
}

func scanRowIntoCourse(rows *sql.Rows) (*types.Course, error) {
	course := new(types.Course)
	err := rows.Scan(
		&course.Id,
		&course.Name,
		&course.Credit,
		&course.Capacity,
		&course.CreatedAt,
		&course.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return course, nil
}
