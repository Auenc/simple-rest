package services

import (
	"github.com/auenc/simple-rest/models"
	"github.com/auenc/simple-rest/repositories"
)

type JobService struct {
	JobRepo repositories.JobRepository
}

func (s *JobService) Get(id string) (models.Job, error) {
	return s.JobRepo.Get(id)
}
func (s *JobService) GetAll() ([]models.Job, error) {
	return s.JobRepo.GetAll()
}
func (s *JobService) Create(job models.Job) error {
	return s.JobRepo.Create(job)
}
func (s *JobService) Update(id string, job models.Job) error {
	return s.JobRepo.Update(id, job)
}
func (s *JobService) Delete(id string) error {
	return s.JobRepo.Delete(id)
}
