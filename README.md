# Simple REST API

This simple tutorial aims to teach new Go developers how to create a simple RESTful API using the service/repository pattern. Whilst the provided code will work, it's only a _very_ simple example of how to structure a Go RESTful API using multiple packages

  

## Prerequisites

The only prerequisite for this tutorial is to have Go installed.

  

## Setup

Typically Go projects are organised into the following pattern

  

```

{GITHOSTING}/{AUTHOR}/{PROJECT}

```

If this is your first time writing a go program - you will need to create this structure inside your `src` directory within your GOPATH. For example - if you are using github and your github username is doge, you will need to create the following path inside your GOPATH to start the tutorial

```

github.com/doge/simple-rest

```

## Third party libraries

In this tutorial we will be using [Gorilla mux](https://github.com/gorilla/mux) for creating the individual routes that our API will sit behind. Whilst it's not required to create a RESTful service in Go, it does make things such as URL variables easier to access.

  

To install Mux, once inside your project directory, run the following command:

```

go get github.com/gorilla/mux

```

  

## Project structure

The below diagram shows the project structure of the simple API we will be creating.

```

├── endpoints

│ └── JobREST.go

├── models

│ └── Job.go

├── repositories

│ ├── InMemoryJobRepository.go

│ └── JobRepository.go

├── services

│ └── JobService.go

└── main.go

```

### endpoints

Endpoints is a package that is responsible for talking to the outside world. In the tutorial we are creating an RESTful API. However, if in the future you would want to add gRPC or perhaps SOAP, you would be able to add the additional implementations within this package.

  

### models

Models contains the data entities that our API will be managing.

  

### repositories

A repository is a layer of abstraction over the database. The idea being that we can use a Repository interface that will define a behaviour contract that the rest of the system can use to get/create/update/delete the Job models that we have in our system. In this example, we will only be implementing an in-memory database, however, if in the future we wanted to implement a SQL database - it would be as simple as creating a new class in the repositories package that would handle the calls to the database. We could then pass that to our services in the exact same way that we are passing the InMemoryJobRepository class as seen in `main.go`. The advantage of this is that we have a kind of "host swappable" data layer that, if done correctly, means we can change data storage systems without effecting the rest of the system.

  

### services

A service is a level of abstraction that sits in between the API's endpoints and the Repository layer. This is typically where we would implement any business logic required by our API. As this is a simple tutorial, the service itself doesn't do anything other than pass calls off the the repository.

  

## Models

The Job model is our API's representation of a Job listing.

```Go

// models/Job.go

package models

  

type  Job  struct {

ID string  `json:"id"`  //Change ID to id when marshaling to JSON

Name string  `json:"name"`  // Name -> name

Description string  `json:"description"`  //Description -> description

}

```

As mentioned in the comments in the above snippet - the `json:"id"` is a means of us to change the name of the field when we convert this object into JSON. There are two other very commonly used things we can do with this not shown in the example:

  

- `json:"-"` - ignores the field entirely, so it will not show up in the JSON version of the object

- `json:"omitempty"` - Will ignore the field if the field is empty or null

## Repositories

As mentioned above - A Repository is an abstraction layer above the Database. Our aim is to create an interface that the rest of the system can use to be able to preform certain actions without getting into the nitty-gritty of actually storing/fetching data. The first step we need to do is write the interface that the rest of the system will interact with.

### JobRepository interface

```Go

// repositories/JobRepository.Go

package repositories

  

import (

"github.com/auenc/simple-rest/models"

)

  

type  JobRepository  interface {

Get(id string) (models.Job, error)

GetAll() ([]models.Job, error)

Create(job models.Job) error

Update(id string, job models.Job) error

Delete(id string) error

}

```

In Go interfaces are a little different to other likes such as Java. Rather than having to explicitly state that a class implements an interface, instead, we're able to use any class that has the methods described by an interface. In the next section we will implement our in-memory repository that will not reference the `JobRepository` at all, but will have all of the methods specified by the interface. This will mean that, as far as the rest of the system is concerned, they are the same thing.

### InMemoryJobRepository

Our in-memory repository is going to be way our data is stored. This is a very simple implementation (and not a very good one at that). We are just going to simply hold all of our data in an array.

  

If you have never seen a method in Go before, this may look a little weird. If you've used languages such as Java in the passed, you're probably used to seeing methods inside the class definition. The difference here is that we are instead "attaching" the methods to an instance of the class we have created.

```Go

// repositories/InMemoryJobRepository.Go

package repositories

  

import (

"errors"

"strconv"

  

"github.com/auenc/simple-rest/models"

)

  

type  InMemoryJobRepository  struct {

Jobs []models.Job

}

  

func  NewInMemoryJobRepository() *InMemoryJobRepository {

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

return  nil

}

func (r *InMemoryJobRepository) Update(id string, job models.Job) error {

for i, j := range r.Jobs {

if j.ID == id {

job.ID = id

r.Jobs[i] = job

}

return  nil

}

return errors.New("Job not found")

}

func (r *InMemoryJobRepository) Delete(id string) error {

for i, j := range r.Jobs {

if j.ID == id {

// delete from slice

r.Jobs = append(r.Jobs[:i], r.Jobs[i+1:]...)

return  nil

}

}

return errors.New("Job not found")

}

```

## Services

Services are another level of abstraction that sits in-between our Endpoints and our Repository layer. Here we can add any business logic is required of our API. As this is a simple API, our service is essentially just going to pass calls off to our Repository.

```Go

// services/JobService.Go

package services

  

import (

"github.com/auenc/simple-rest/models"

"github.com/auenc/simple-rest/repositories"

)

  

type  JobService  struct {

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

```

## Endpoints

Endpoints are how the outside world interacts with our service. They'll often mirror our service layer, but have a little bit more work to do. They're responsible for taking the requests we get from, in our example REST calls, and creating the actual objects that our system will interact with/store.

### RESTful API

As we're making a RESTful API - we're going to be creating a RESTful endpoint (shocker, I know). However, if in the future we wanted to add additional endpoint type (such as SOAP), it would be as simple as creating an additional class for that implementation

  

There are two important things to note here that we haven't come across in the tutorial so far - `http.ResponseWriter` and `http.Request`.

  

`http.ResponseWriter` is a ResponseWriter interface that essentially is our means of communicating back to the user. In our example we're going to use this in two ways. The first is:

```Go

w.WriteHeader(http.StatusBadRequest)

```

  

here we are setting the status code of a Request. In this case we are telling the user that they have not provided the correct request, so we are unable to complete the action.

  

The second is a little more complicated as we're only ever responding with JSON objects.

```Go

json.NewEncoder(w).Encode(map[string]string{"error": "no id given"})

```

Here we are asking the JSON package to create a new encoder using our `http.ResponseWriter` as the "where we want our JSON object to go" and passing it a map of the error we have detected.

  

```Go

package endpoints

  

import (

"encoding/json"

"net/http"

  

"github.com/auenc/simple-rest/models"

"github.com/auenc/simple-rest/services"

"github.com/gorilla/mux"

)

  

type  JobREST  struct {

JobService *services.JobService

}

  

func  NewJobREST(service *services.JobService) *JobREST {

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

```

  

## Main

Now that we have all of the components created, it's time to bring them all together in our main file.

  

As mentioned earlier in the tutorial - we're using Gorilla/Mux as our router so we can specify URL parameters easier.

  

```Go

package main

  

import (

"log"

"net/http"

  

"github.com/auenc/simple-rest/endpoints"

"github.com/auenc/simple-rest/repositories"

"github.com/auenc/simple-rest/services"

"github.com/gorilla/mux"

)

  

func  main() {

inMemoryRepo := repositories.NewInMemoryJobRepository()

jobService := &services.JobService{JobRepo: inMemoryRepo}

jobAPI := endpoints.NewJobREST(jobService)

  

router := mux.NewRouter()

  

router.HandleFunc("/", jobAPI.Create).Methods("POST")

router.HandleFunc("/", jobAPI.GetAll).Methods("GET")

router.HandleFunc("/{id}", jobAPI.Get).Methods("GET")

router.HandleFunc("/{id}", jobAPI.Update).Methods("PUT")

router.HandleFunc("/{id}", jobAPI.Delete).Methods("DELETE")

  

log.Fatalf("%s\n", http.ListenAndServe(":8080", router))

}

```

  

Now to compile and run this we run the following commands. *Note* if your directory name is not simple-rest, then your program will be named something different

```bash

go build

./simple-rest

```

  

## Conclusion

If you've followed the tutorial correctly (and I haven't messed any of the copy&pasting up) you should now have a working Job listing API using the service/repository pattern. If something is not working for you, please feel free to download the code base on this repository to compare any differences/try it out.