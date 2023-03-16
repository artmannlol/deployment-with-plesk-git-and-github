package main

import "net/http"

func main() {
	// Create route.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// For this purpose we only print "Hello World".
		_, _ = w.Write([]byte("Hello World"))
	})

	// Start server.
	_ = http.ListenAndServe(":8080", nil)
}
