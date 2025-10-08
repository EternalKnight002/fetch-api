package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// fetchUsersFromServer fetches data from an external API
func fetchUsersFromServer() ([]map[string]interface{}, error) {
	// Example public API (you can replace this with your own)
	url := "https://jsonplaceholder.typicode.com/users"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned status: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var users []map[string]interface{}
	if err := json.Unmarshal(body, &users); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return users, nil
}

// handleUsers handles GET /users
func handleUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := fetchUsersFromServer()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Root endpoint
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Fetch API Server 🌍")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/users", handleUsers)

	fmt.Println("✅ Server running at http://localhost:8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
