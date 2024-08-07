package s3manager

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
)

// HandleDeleteObject deletes an object.
func HandleDeleteObject(s3 S3, bucketMap map[string]string) http.HandlerFunc {
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

		err := s3.RemoveObject(r.Context(), bucketName, objectName, minio.RemoveObjectOptions{})
		if err != nil {
			handleHTTPError(w, fmt.Errorf("error removing object: %w", err))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
