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
)

type SubmitRequest struct {
	Count  int     `json:"count"`
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
	jobMutex.Unlock()

	go jobs.ProcessJob(job)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(SubmitResponse{JobID: jobID})
}
