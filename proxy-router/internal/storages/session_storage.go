package storages

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Session struct {
	Id           string
	UserAddr     string
	ProviderAddr string
}

type User struct {
	Addr   string
	PubKey string
	Url    string
}

type SessionStorage struct {
	db *Storage
}

func NewSessionStorage(storage *Storage) *SessionStorage {
	return &SessionStorage{
		db: storage,
	}
}

func (s *SessionStorage) GetSession(id string) (*Session, bool) {
	id = strings.ToLower(id)
	key := fmt.Sprintf("session:%s", id)

	sessionJson, err := s.db.Get([]byte(key))
	if err != nil {
		return nil, false
	}

	session := &Session{}
	err = json.Unmarshal(sessionJson, session)
	if err != nil {
		return nil, false
	}

	return session, true
}

func (s *SessionStorage) GetUser(addr string) (*User, bool) {
	addr = strings.ToLower(addr)
	key := fmt.Sprintf("user:%s", addr)
	userJson, err := s.db.Get([]byte(key))
	if err != nil {
		return nil, false
	}

	user := &User{}
	err = json.Unmarshal(userJson, user)
	if err != nil {
		return nil, false
	}

	return user, true
}

func (s *SessionStorage) AddSession(session *Session) error {
	sessionId := strings.ToLower(session.Id)
	key := fmt.Sprintf("session:%s", sessionId)
	sessionJson, err := json.Marshal(session)
	if err != nil {
		return err
	}

	err = s.db.Set([]byte(key), sessionJson)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionStorage) AddUser(user *User) error {
	addr := strings.ToLower(user.Addr)
	key := fmt.Sprintf("user:%s", addr)
	userJson, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = s.db.Set([]byte(key), userJson)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionStorage) RemoveSession(id string) error {
	sessionId := strings.ToLower(id)
	key := fmt.Sprintf("session:%s", sessionId)
	err := s.db.Delete([]byte(key))
	if err != nil {
		return err
	}
	return nil
}
