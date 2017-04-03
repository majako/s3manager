package main

import "net/http"

// IndexViewHandler forwards to "/buckets"
func IndexViewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/buckets", http.StatusPermanentRedirect)
	})
}
