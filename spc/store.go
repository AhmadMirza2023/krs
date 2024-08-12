package spc

import (
	"database/sql"

	"github.com/AhmadMirza2023/krs/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateSPC(spc types.SPC) (*types.SPC, error) {
	newSPC, err := s.db.Exec("INSERT INTO study_plan_cards (user_id, course_id) VALUES (?,?)", spc.UserId, spc.CourseId)
	if err != nil {
		return nil, err
	}
	id, err := newSPC.LastInsertId()
	if err != nil {
		return nil, err
	}
	spc.Id = int(id)
	return &spc, nil
}

func (s *Store) GetSPCByUserId(userId int) (*types.SPC, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	spc := new(types.SPC)
	for rows.Next() {
		spc, err = scanRowIntoSPC(rows)
		if err != nil {
			return nil, err
		}
	}
	if spc.UserId == 0 {
		return nil, err
	}
	return spc, nil
}

func scanRowIntoSPC(rows *sql.Rows) (*types.SPC, error) {
	spc := new(types.SPC)
	err := rows.Scan(
		&spc.Id,
		&spc.UserId,
		&spc.CourseId,
		&spc.CreatedAt,
		&spc.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return spc, nil
}
