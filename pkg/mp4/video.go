package mp4

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Video is an mp4 video
type Video struct {
	// Path is an patho to video file
	Path string

	// Data represents the content of the video file
	Data []byte
}

// NewVideo creates a new structure of Video
func NewVideo(p string) (*Video, error) {
	var v Video

	if !exists(p) {
		return &v, errors.New("File not found")
	}

	if !fileExtension(p) {
		return &v, errors.New("Unsuitable file extension")
	}

	// Read file data
	data, err := ioutil.ReadFile(p)
	if err != nil {
		return &v, err
	}
	v.Data = data

	v.Path = p

	return &v, nil
}

func exists(p string) bool {
	_, err := os.Stat(p)
	return !os.IsNotExist(err)
}

func fileExtension(p string) bool {
	return filepath.Ext(p) == ".mp4"
}
