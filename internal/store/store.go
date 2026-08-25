package store

import (
	"database/sql"
	"fmt"
	"idiomchain/internal/model"
	_ "modernc.org/sqlite"
	"time"
)

type Store struct {
	db   *sql.DB
	path string
}

func Open(path string) (*Store, error) {
	db, e := sql.Open("sqlite", path)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db, path: path}
	if e = s.init(); e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) init() error {
	_, e := s.db.Exec(`create table if not exists sessions(id text primary key,name text not null,created_at text not null); create table if not exists entries(id text primary key,session_id text,text text,tail text,created_at text); create table if not exists rejections(id text primary key,session_id text,input text,reason text,created_at text); create table if not exists settings(key text primary key,value text,updated_at text)`)
	return e
}
func (s *Store) Close() error { return s.db.Close() }
func (s *Store) Path() string { return s.path }
func (s *Store) SaveSession(v model.ChainSession) error {
	_, e := s.db.Exec(`insert or ignore into sessions(id,name,created_at) values(?,?,?)`, v.ID, v.Name, v.CreatedAt.Format(time.RFC3339Nano))
	return e
}
func (s *Store) Session(id string) (model.ChainSession, error) {
	var v model.ChainSession
	var t string
	e := s.db.QueryRow(`select id,name,created_at from sessions where id=?`, id).Scan(&v.ID, &v.Name, &t)
	if e == nil {
		v.CreatedAt, _ = time.Parse(time.RFC3339Nano, t)
	}
	return v, e
}
func (s *Store) AddEntry(v model.IdiomEntry) error {
	if !v.Valid() {
		return fmt.Errorf("invalid entry")
	}
	_, e := s.db.Exec(`insert into entries(id,session_id,text,tail,created_at) values(?,?,?,?,?)`, v.ID, v.SessionID, v.Text, v.Tail, v.CreatedAt.Format(time.RFC3339Nano))
	return e
}
func (s *Store) AddRejection(v model.Rejection) error {
	_, e := s.db.Exec(`insert into rejections(id,session_id,input,reason,created_at) values(?,?,?,?,?)`, v.ID, v.SessionID, v.Input, v.Reason, v.CreatedAt.Format(time.RFC3339Nano))
	return e
}
