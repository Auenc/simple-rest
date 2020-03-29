package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/auenc/simple-rest/models"
	"github.com/auenc/simple-rest/services"
	"github.com/gorilla/mux"
)

type JobREST struct {
	JobService *services.JobService
}

func NewJobREST(service *services.JobService) *JobREST {
	return &JobREST{
		JobService: service,
	}
}

func (j *JobREST) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "no id given"})
		return
	}
	job, err := j.JobService.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(job)
}
func (j *JobREST) GetAll(w http.ResponseWriter, r *http.Request) {
	jobs, err := j.JobService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jobs)
}
func (j *JobREST) Create(w http.ResponseWriter, r *http.Request) {
	var job models.Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid requst body"})
		return
	}
	err = j.JobService.Create(job)
	if err != nil {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
func (j *JobREST) Update(w http.ResponseWriter, r *http.Request) {
	var job models.Job
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "no id given"})
		return
	}
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid requst body"})
		return
	}
	err = j.JobService.Update(id, job)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

}
func (j *JobREST) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "no id given"})
		return
	}
	err := j.JobService.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
}
