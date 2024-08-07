package s3manager

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/minio/minio-go/v7"
)

// HandleBucketsView renders all buckets on an HTML page.
func HandleBucketsView(s3 S3, templates fs.FS, bucketName string, allowDelete bool) http.HandlerFunc {
	type pageData struct {
		Buckets     []minio.BucketInfo
		AllowDelete bool
	}

	return func(w http.ResponseWriter, r *http.Request) {
		buckets, err := s3.ListBuckets(r.Context())
		if err != nil {
			handleHTTPError(w, fmt.Errorf("error listing buckets: %w", err))
			return
		}

		buckets = filterBuckets(buckets, bucketName)

		data := pageData{
			Buckets:     buckets,
			AllowDelete: allowDelete,
		}

		t, err := template.ParseFS(templates, "layout.html.tmpl", "buckets.html.tmpl")
		if err != nil {
			handleHTTPError(w, fmt.Errorf("error parsing template files: %w", err))
			return
		}
		err = t.ExecuteTemplate(w, "layout", data)
		if err != nil {
			handleHTTPError(w, fmt.Errorf("error executing template: %w", err))
			return
		}
	}
}

func filterBuckets(buckets []minio.BucketInfo, bucketName string) []minio.BucketInfo {
	var filteredBuckets []minio.BucketInfo
	for _, bucket := range buckets {
		if bucket.Name != bucketName {
			continue
		}
		filteredBuckets = append(filteredBuckets, bucket)
	}
	return filteredBuckets
}
