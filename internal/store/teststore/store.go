package teststore

import (
	"awesomeProject/internal/app/model"
	"awesomeProject/internal/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() store.Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}
	return s.userRepository
}
