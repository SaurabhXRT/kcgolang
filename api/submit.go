package api

import (
	"encoding/json"
	"net/http"
	"sync"

	"kcassignment/jobs"
)

var (
	jobIDSeq = 1
	jobMutex sync.Mutex
	//i have used a map store using make for temporarily storing the job details
	jobsMap  = make(map[int]*jobs.Job) 
)

type SubmitRequest struct {
	Count  int          `json:"count"`
	Visits []jobs.Visit `json:"visits"`
}

type SubmitResponse struct {
	JobID int `json:"job_id"`
}

func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	var req SubmitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "invalid request payload"}`, http.StatusBadRequest)
		return
	}

	if req.Count != len(req.Visits) {
		http.Error(w, `{"error": "count mismatch with visits"}`, http.StatusBadRequest)
		return
	}

	jobMutex.Lock()
	jobID := jobIDSeq
	jobIDSeq++
	job := jobs.CreateJob(jobID, req.Visits)
	jobsMap[jobID] = job 
	jobMutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(SubmitResponse{JobID: jobID})


	//once the job is submitted asynchronous function will start to process 
	//the image and if the image is processed it status will be completed or else meanwhile it will 
	//show ongoing i have adedd sleeptime of 30 second 
	go func() {
		jobs.ProcessJob(job)
	}()
}
