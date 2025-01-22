package jobs

import (
	"sync"
	
	"kcassignment/utils"
)
var jobMutex sync.Mutex

func ProcessJob(job *Job) {
	for _, visit := range job.Visits {
		for _, imageURL := range visit.ImageURLs {
			err := ProcessImage(imageURL)
			if err != nil {
				jobMutex.Lock()
				job.Errors = append(job.Errors, ErrorDetail{
					StoreID: visit.StoreID,
					Error:   err.Error(),
				})
				job.Status = "failed"
				jobMutex.Unlock()
				return
			}
		}
	}

	jobMutex.Lock()
	job.Status = "completed"
	jobMutex.Unlock()
}

func ProcessImage(url string) error {
	_, err := utils.DownloadImage(url)
	if err != nil {
		return err
	}

	utils.SleepRandom()
	return nil
}
