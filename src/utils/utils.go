package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
)

// DownloadFile is a function for downloading file by its url
func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// Upload file
func Upload(uploadURL, field, filename string, reader io.Reader, holder interface{}) (interface{}, error) {
	client := resty.New()
	resp, err := client.R().
		SetFileReader(field, filename, reader).
		Post(uploadURL)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp.Body(), holder)
	if err != nil {
		return nil, err
	}
	return holder, nil
}
