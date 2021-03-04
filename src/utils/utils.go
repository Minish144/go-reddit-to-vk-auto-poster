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
