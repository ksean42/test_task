package storage

import (
	"fmt"
	"github.com/ksean42/test_task/pkg/model"
	"sync"
)

type Storage struct {
	sync.RWMutex
	UserGrades map[string]*model.UserGrade
}

func NewStorage() *Storage {
	return &Storage{UserGrades: make(map[string]*model.UserGrade)}
}

func (s *Storage) Get(id string) (*model.UserGrade, error) {
	s.RLock()
	defer s.RUnlock()
	if value, ok := s.UserGrades[id]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("not found")
}

func (s *Storage) Set(grade *model.UserGrade) error {
	s.Lock()
	if _, ok := s.UserGrades[grade.UserId]; ok {
		if grade.PostpaidLimit != 0 {
			s.UserGrades[grade.UserId].PostpaidLimit = grade.PostpaidLimit
		}
		if grade.Spp != 0 {
			s.UserGrades[grade.UserId].Spp = grade.Spp
		}
		if grade.ShippingFee != 0 {
			s.UserGrades[grade.UserId].ShippingFee = grade.ShippingFee
		}
		if grade.ReturnFee != 0 {
			s.UserGrades[grade.UserId].ReturnFee = grade.ReturnFee
		}
	} else {
		s.UserGrades[grade.UserId] = grade
	}
	s.Unlock()
	return nil
}
