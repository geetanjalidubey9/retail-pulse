package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func GetJobStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the job ID from query parameters
	jobIDStr := r.URL.Query().Get("jobid")
	if jobIDStr == "" {
		http.Error(w, `{"error":"jobid is required"}`, http.StatusBadRequest)
		return
	}

	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		http.Error(w, `{"error":"invalid jobid"}`, http.StatusBadRequest)
		return
	}

	// Check if the job exists
	mu.Lock()
	job, exists := jobs[jobID]
	mu.Unlock()

	if !exists {
		http.Error(w, `{"error":"job not found"}`, http.StatusBadRequest)
		return
	}

	// Respond with job status
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": job.Status,
		"job_id": job.ID,
		"errors": job.Errors,
	})
}
