package service

import (
	"fmt"
	"github.com/ksean42/test_task/pkg/model"
	"github.com/ksean42/test_task/pkg/storage"
)

type UserGradeService struct {
	repo storage.Repository
}

func NewUserGradeService(repo storage.Repository) *UserGradeService {
	return &UserGradeService{repo: repo}
}

func (u *UserGradeService) Get(id string) (*model.UserGrade, error) {
	return u.repo.Get(id)
}
func (u *UserGradeService) Set(grade *model.UserGrade) error {
	if grade.UserId == "" || grade.PostpaidLimit < 0 || grade.Spp < 0 ||
		grade.ShippingFee < 0 || grade.ReturnFee < 0 {
		return fmt.Errorf("invalid request")
	}
	return u.repo.Set(grade)
}
