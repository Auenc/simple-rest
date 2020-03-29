package repositories

import (
	"github.com/auenc/simple-rest/models"
)

type JobRepository interface {
	Get(id string) (models.Job, error)
	GetAll() ([]models.Job, error)
	Create(job models.Job) error
	Update(id string, job models.Job) error
	Delete(id string) error
}
