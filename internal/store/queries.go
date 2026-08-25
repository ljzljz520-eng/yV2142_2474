package store

import (
	"idiomchain/internal/model"
	"time"
)

func (s *Store) Entries(session string, page, size int) ([]model.IdiomEntry, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	rows, e := s.db.Query(`select id,session_id,text,tail,created_at from entries where session_id=? order by rowid limit ? offset ?`, session, size, (page-1)*size)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []model.IdiomEntry{}
	for rows.Next() {
		var v model.IdiomEntry
		var t string
		if e = rows.Scan(&v.ID, &v.SessionID, &v.Text, &v.Tail, &t); e != nil {
			return nil, e
		}
		v.CreatedAt, _ = time.Parse(time.RFC3339Nano, t)
		out = append(out, v)
	}
	return out, rows.Err()
}
func (s *Store) AllEntries(session string) ([]model.IdiomEntry, error) {
	return s.Entries(session, 1, 100000)
}
func (s *Store) Rejections(session string) ([]model.Rejection, error) {
	rows, e := s.db.Query(`select id,session_id,input,reason,created_at from rejections where session_id=? order by rowid`, session)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []model.Rejection{}
	for rows.Next() {
		var v model.Rejection
		var t string
		if e = rows.Scan(&v.ID, &v.SessionID, &v.Input, &v.Reason, &t); e != nil {
			return nil, e
		}
		v.CreatedAt, _ = time.Parse(time.RFC3339Nano, t)
		out = append(out, v)
	}
	return out, rows.Err()
}
func (s *Store) CountEntries(session string) (int, error) {
	var n int
	e := s.db.QueryRow(`select count(*) from entries where session_id=?`, session).Scan(&n)
	return n, e
}
