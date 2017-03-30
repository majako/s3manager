package buckets

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mastertinner/s3-manager/web"
	minio "github.com/minio/minio-go"
)

// DeleteHandler deletes a bucket
func DeleteHandler(s3 *minio.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		err := s3.RemoveBucket(vars["bucketName"])
		if err != nil {
			msg := "error removing bucket"
			web.HandleHTTPError(w, msg, err, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}
