package store

import (
	"idiomchain/internal/model"
	"time"
)

func (s *Store) SetSetting(key, value string, at time.Time) error {
	_, e := s.db.Exec(`insert into settings(key,value,updated_at) values(?,?,?) on conflict(key) do update set value=excluded.value,updated_at=excluded.updated_at`, key, value, at.Format(time.RFC3339Nano))
	return e
}
func (s *Store) GetSetting(key string) (model.Setting, error) {
	var v model.Setting
	var t string
	e := s.db.QueryRow(`select key,value,updated_at from settings where key=?`, key).Scan(&v.Key, &v.Value, &t)
	if e == nil {
		v.UpdatedAt, _ = time.Parse(time.RFC3339Nano, t)
	}
	return v, e
}
func (s *Store) DeleteSetting(key string) error {
	_, e := s.db.Exec(`delete from settings where key=?`, key)
	return e
}
