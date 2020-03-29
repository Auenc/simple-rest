package repositories

import (
	"errors"
	"strconv"

	"github.com/auenc/simple-rest/models"
)

type InMemoryJobRepository struct {
	Jobs []models.Job
}

func NewInMemoryJobRepository() *InMemoryJobRepository {
	jobs := make([]models.Job, 0)
	return &InMemoryJobRepository{
		Jobs: jobs,
	}
}

func (r *InMemoryJobRepository) Get(id string) (models.Job, error) {
	var job models.Job

	for _, j := range r.Jobs {
		if j.ID == id {
			return j, nil
		}
	}

	return job, errors.New("Job not found")
}
func (r *InMemoryJobRepository) GetAll() ([]models.Job, error) {
	return r.Jobs, nil
}
func (r *InMemoryJobRepository) Create(job models.Job) error {
	job.ID = strconv.Itoa(len(r.Jobs))
	r.Jobs = append(r.Jobs, job)
	return nil
}
func (r *InMemoryJobRepository) Update(id string, job models.Job) error {
	for i, j := range r.Jobs {
		if j.ID == id {
			job.ID = id
			r.Jobs[i] = job
		}
		return nil
	}
	return errors.New("Job not found")
}
func (r *InMemoryJobRepository) Delete(id string) error {
	for i, j := range r.Jobs {
		if j.ID == id {
			// delete from slice
			r.Jobs = append(r.Jobs[:i], r.Jobs[i+1:]...)
			return nil
		}
	}
	return errors.New("Job not found")
}
