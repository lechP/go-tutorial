package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var url = "https://jsonplaceholder.typicode.com/users"

// calls example API, parses and prints the response
func main() {

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	// Ensure the response body is closed after reading
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Parse the JSON response
	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return
	}

	// Print the parsed data
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

}
