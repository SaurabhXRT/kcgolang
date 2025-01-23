package jobs

import (
	"errors"
	"image"
	"sync"

	_ "image/jpeg" 
	_ "image/png" 

	"kcassignment/utils"
)

var jobProcessingMutex sync.Mutex

func ProcessJob(job *Job) {
	for _, visit := range job.Visits {
		for _, imageURL := range visit.ImageURLs {
			err := ProcessImage(imageURL)
			if err != nil {
				jobProcessingMutex.Lock()
				job.Errors = append(job.Errors, ErrorDetail{
					StoreID: visit.StoreID,
					Error:   err.Error(),
				})
				job.Status = "failed"
				jobProcessingMutex.Unlock()
				return
			}
		}
	}

	jobProcessingMutex.Lock()
	job.Status = "completed"
	jobProcessingMutex.Unlock()
}

func ProcessImage(url string) error {
	
	imageData, err := utils.DownloadImage(url)
	if err != nil {
		return errors.New(err.Error())
	}
	img, _, err := image.Decode(imageData)
	if err != nil {
		return errors.New("failed to decode image: " + err.Error())
	}
	bounds := img.Bounds()
	perimeter := 2 * (bounds.Dx() + bounds.Dy())

	utils.LogImagePerimeter(perimeter)

	utils.SleepRandom()
	return nil
}
