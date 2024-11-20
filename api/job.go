package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Job struct {
	ID     int        `json:"job_id"`
	Count  int        `json:"count"`
	Visits []Visit    `json:"visits"`
	Status string     `json:"status"` // "ongoing", "completed", "failed"
	Errors []JobError `json:"errors,omitempty"`
}

type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}

type JobError struct {
	StoreID string `json:"store_id"`
	Error   string `json:"error"`
}

var (
	jobs       = make(map[int]*Job) // To store all jobs
	jobCounter = 0
	mu         sync.Mutex // To handle concurrency
)

func SubmitJob(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var job Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil || job.Count != len(job.Visits) {
		http.Error(w, `{"error":"Invalid payload or count mismatch"}`, http.StatusBadRequest)
		return
	}

	// Assign a unique ID to the job
	mu.Lock()
	jobCounter++
	job.ID = jobCounter
	job.Status = "ongoing"
	jobs[job.ID] = &job
	mu.Unlock()

	// Process the job asynchronously
	go processJob(&job)

	// Respond with the job ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"job_id": job.ID})
}

func processJob(job *Job) {
	var wg sync.WaitGroup

	for _, visit := range job.Visits {
		wg.Add(1)

		go func(visit Visit) {
			defer wg.Done()

			for _, imageURL := range visit.ImageURLs {
				// Simulate image processing with a random sleep
				err := processImage(imageURL)
				if err != nil {
					mu.Lock()
					job.Status = "failed"
					job.Errors = append(job.Errors, JobError{StoreID: visit.StoreID, Error: err.Error()})
					mu.Unlock()
					return
				}
			}
		}(visit)
	}

	// Wait for all images to be processed
	wg.Wait()

	mu.Lock()
	if len(job.Errors) == 0 {
		job.Status = "completed"
	}
	mu.Unlock()
}

func processImage(url string) error {
	// Simulate image download and processing
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)+100))

	// Example: Check if URL is valid (you can replace this with real download logic)
	if url == "" {
		return fmt.Errorf("invalid image URL")
	}
	return nil
}
