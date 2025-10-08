package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// fetchUsersFromServer fetches data from an external API and returns raw JSON-decoded value.
func fetchUsersFromServer() ([]map[string]interface{}, error) {
	url := "https://jsonplaceholder.typicode.com/users"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to GET external URL %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("external server returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read external response body: %w", err)
	}

	var users []map[string]interface{}
	if err := json.Unmarshal(body, &users); err != nil {
		return nil, fmt.Errorf("failed to parse JSON from external server: %w", err)
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
		// Log detailed error on server console and return a short message to client
		log.Printf("ERROR fetching users: %v\n", err)
		http.Error(w, "Failed to fetch users from upstream", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// encode response
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Printf("ERROR encoding users to response: %v\n", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Fetch API Server 🌍")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/users", handleUsers)

	addr := ":8083" // matches the port you opened in the browser
	fmt.Printf("✅ Server running at http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
