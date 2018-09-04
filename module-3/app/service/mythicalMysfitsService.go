package main

import (
    "net/http"
    "io/ioutil"
)

// For http://localhost:8088
func healthCheckResponse(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Nothing here, used for health check. Try /misfits instead."))
}

// Show mysfits-response.json for http://localhost:8088/misfits
func showMisfits(w http.ResponseWriter, r *http.Request) {
    // Read Misfits data from file and show it
    body, _ := ioutil.ReadFile("mysfits-response.json")

    // Let the web server know it's JSON
    w.Header().Set("Content-Type", "application/json")
    w.Write(body)
}

func main() {
    mux := http.NewServeMux()
    mux.Handle("/", http.HandlerFunc(healthCheckResponse))
    mux.Handle("/misfits", http.HandlerFunc(showMisfits))
    http.ListenAndServe(":8088", mux)
}