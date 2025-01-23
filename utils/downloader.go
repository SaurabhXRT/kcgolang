package utils

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

func DownloadImage(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return nil, errors.New("failed to download image")
	}
	defer res.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, res.Body)
	if err != nil {
		return nil, errors.New("failed to read image data")
	}

	return &buf, nil
}

func LogImagePerimeter(perimeter int) {
	println("Calculated perimeter:", perimeter)
}