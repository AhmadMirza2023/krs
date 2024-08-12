package user

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

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if user.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Nim,
		&user.Semester,
		&user.Major,
		&user.Faculty,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) CreateUser(user types.User) (*types.User, error) {
	result, err := s.db.Exec("INSERT INTO users (email, password, name, nim, semester, major, faculty) VALUES (?,?,?,?,?,?,?)", user.Email, user.Password, user.Name, user.Nim, user.Semester, user.Major, user.Faculty)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.Id = int(id)

	return &user, nil
}
