package main

import (
	"errors"
	"sync"
)

type Account struct {
	Id         int    `json:"id" gorm:"primary_key"`
	Number     string `json:"number"`
	MoneyCount int    `json:"money_count"`
}

type Storage interface {
	Insert(a *Account)
	Get(id int) (Account, error)
	Update(id int, a *Account)
	Delete(id int)
}

type MemoryStorage struct {
	counter int
	data    map[int]Account
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:    make(map[int]Account),
		counter: 1,
	}
}

func (s *MemoryStorage) Insert(a *Account) {
	s.Lock()

	a.Id = s.counter
	s.data[a.Id] = *a

	s.counter++

	s.Unlock()
}

func (s *MemoryStorage) Get(id int) (Account, error) {
	s.Lock()
	defer s.Unlock()

	account, ok := s.data[id]
	if !ok {
		return account, errors.New("Account not found")
	}

	return account, nil
}

func (s *MemoryStorage) Update(id int, a *Account) {
	s.Lock()
	s.data[id] = *a
	s.Unlock()
}

func (s *MemoryStorage) Delete(id int) {
	s.Lock()
	delete(s.data, id)
	s.Unlock()
}
