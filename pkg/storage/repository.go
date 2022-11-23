package storage

import "github.com/ksean42/test_task/pkg/model"

type Repository interface {
	Get(id string) (*model.UserGrade, error)
	Set(grade *model.UserGrade) error
}
