package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"kcassignment/jobs"
)

type StatusResponse struct {
	Status string        `json:"status"`
	JobID  int           `json:"job_id"`
	Error  []jobs.ErrorDetail `json:"error,omitempty"`
}

func GetJobStatusHandler(w http.ResponseWriter, r *http.Request) {
	jobIDStr := r.URL.Query().Get("jobid")
	if jobIDStr == "" {
		http.Error(w, `{"error": "jobid parameter is required"}`, http.StatusBadRequest)
		return
	}

	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		http.Error(w, `{"error": "invalid jobid"}`, http.StatusBadRequest)
		return
	}

	jobMutex.Lock()
	job, exists := jobsMap[jobID]
	jobMutex.Unlock()

	if !exists {
		http.Error(w, `{"error": "job not found"}`, http.StatusNotFound)
		return
	}

	response := StatusResponse{
		Status: job.Status,
		JobID:  job.ID,
		Error:  job.Errors,
	}
	json.NewEncoder(w).Encode(response)
}
