package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
	"myapp/backend/internal/middleware"
)

func main() {
    router := mux.NewRouter()
    
	router.Use(middleware.CORS)
    // API routes
    api := router.PathPrefix("/api").Subrouter()
    api.HandleFunc("/hello", helloHandler).Methods("GET")
    
    // Static files
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend/dist")))
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Hello from Go!"})
}