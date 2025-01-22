package utils

import (
	"errors"
	"net/http"
)

func DownloadImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("failed to download image")
	}
	defer resp.Body.Close()
	return nil, nil
}
