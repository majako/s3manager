package s3manager

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
)

// HandleGetObject downloads an object to the client.
func HandleGetObject(s3 S3, forceDownload bool, bucketMap map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bucketGuid := mux.Vars(r)["bucketGuid"]
		bucketName := ""
		if val, ok := bucketMap[bucketGuid]; ok {
			bucketName = val
		} else {
			handleHTTPUnauthorizedError(w, fmt.Errorf("bucket not found"))
			return
		}
		objectName := mux.Vars(r)["objectName"]

		object, err := s3.GetObject(r.Context(), bucketName, objectName, minio.GetObjectOptions{})
		if err != nil {
			handleHTTPError(w, fmt.Errorf("error getting object: %w", err))
			return
		}

		if forceDownload {
			w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", objectName))
			w.Header().Set("Content-Type", "application/octet-stream")
		}
		_, err = io.Copy(w, object)
		if err != nil {
			handleHTTPError(w, fmt.Errorf("error copying object to response writer: %w", err))
			return
		}
	}
}
