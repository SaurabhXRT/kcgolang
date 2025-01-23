package jobs

import "time"

type Job struct {
	ID        int
	Visits    []Visit
	Status    string
	Errors    []ErrorDetail
	CreatedAt time.Time
}

type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}

type ErrorDetail struct {
	StoreID string `json:"store_id"`
	Error   string `json:"error"`
}

//var jobs = make(map[int]*Job)

func CreateJob(jobID int, visits []Visit) *Job {
	return &Job{
		ID:     jobID,
		Status: "ongoing",
		Visits: visits,
		Errors: []ErrorDetail{},
	}
}
