package assetserver

import (
	"net/http"
	"sync"
)
import "github.com/gabriel-vasile/mimetype"

var (
	cache = map[string]string{}
	mutex sync.Mutex
)

func GetMimetype(filename string, data []byte) string {
	mutex.Lock()
	defer mutex.Unlock()

	result := cache[filename]
	if result != "" {
		return result
	}

	detect := mimetype.Detect(data)
	if detect == nil {
		result = http.DetectContentType(data)
	} else {
		result = detect.String()
	}

	if result == "" {
		result = "application/octet-stream"
	}

	cache[filename] = result
	return result
}
