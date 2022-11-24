package service

import (
	"compress/gzip"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/ksean42/test_task/pkg/model"
	"github.com/ksean42/test_task/pkg/storage"
	"os"
	"time"
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

func (u *UserGradeService) Backup() (string, error) {
	entries := u.repo.GetAll()
	if len(*entries) == 0 {
		return "", fmt.Errorf("no data in storage")
	}
	s, err := u.writeCSV(entries)
	return s, err
}

func (u *UserGradeService) writeCSV(entries *[]model.UserGrade) (string, error) {
	date := time.Now().Format("2006-01-02T15:04:05-0700")

	file, err := os.Create("csv.gz")
	if err != nil {
		return "", err
	}

	csvData, err := gocsv.MarshalBytes(entries)
	if err != nil {
		return "", nil
	}
	w := gzip.NewWriter(file)
	defer w.Close()
	_, err = w.Write([]byte(date + "\n"))
	_, err = w.Write(csvData)
	if err != nil {
		return "", err
	}
	return "csv.gz", nil
}
