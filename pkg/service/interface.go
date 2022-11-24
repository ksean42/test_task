package service

import "github.com/ksean42/test_task/pkg/model"

type UserGrade interface {
	Get(id string) (*model.UserGrade, error)
	Set(grade *model.UserGrade) error
	Backup() (string, error)
}
