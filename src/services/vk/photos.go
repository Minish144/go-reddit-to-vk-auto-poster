package vk

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/go-vk-api/vk"
)

// WallUploadServer structer which you get from GetWallUploadServer method
type WallUploadServer struct {
	UploadURL string `json:"upload_url"`
	AlbumID   int    `json:"album_id"`
	UserID    int    `json:"user_id"`
}

// PhotoUploadResult structer which you get from UploadPhoto method
type PhotoUploadResult struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
}

// PhotoPost .
type PhotoPost struct {
	photo *os.File
}

// GetWallUploadServer returns address of server for photos uploading
func (c *Client) GetWallUploadServer() *WallUploadServer {
	var wallUploadServer WallUploadServer
	params := vk.RequestParams{
		"group_id": c.GroupID,
	}
	c.Client.CallMethod(
		"photos.getWallUploadServer",
		params,
		&wallUploadServer,
	)
	return &wallUploadServer
}

//UploadPhoto uploads your photo to the provided VK file storage
func (c *Client) UploadPhoto(url, filePath string) (PhotoUploadResult, error) {
	img, _ := os.Open(filePath)
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("photo", filePath)
	if err != nil {
		return PhotoUploadResult{}, err
	}
	if _, err = io.Copy(fw, img); err != nil {
		return PhotoUploadResult{}, err
	}
	w.Close()

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return PhotoUploadResult{}, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return PhotoUploadResult{}, err
	}

	// resp
	uplRes := PhotoUploadResult{}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&uplRes)
	if err != nil {
		return PhotoUploadResult{}, err
	}
	defer res.Body.Close()

	return PhotoUploadResult{
		Server: uplRes.Server,
		Photo:  uplRes.Photo,
		Hash:   uplRes.Hash,
	}, nil
}
