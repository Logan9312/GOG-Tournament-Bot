package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

const BASE_URL = "https://api.challonge.com/v1/"

var API_KEY string

type Tournament struct {
	Name      string `json:"name"`
	URL       string `json:"url,omitempty"`
	SignupCap int    `json:"signup_cap,omitempty"`
	// Include other fields as needed
}

type TournamentCreate struct {
	ApiKey     string     `json:"api_key"`
	Tournament Tournament `json:"tournament"`
}

type TournamentResponse struct {
	Tournament Tournament `json:"tournament"`
}

func formatURL(name string) string {
	// Replace all spaces with underscores
	name = strings.ReplaceAll(name, " ", "_")

	// Use a regular expression to remove unwanted characters
	reg, err := regexp.Compile("[^a-zA-Z0-9_]+")
	if err != nil {
		fmt.Println(err)
	}
	name = reg.ReplaceAllString(name, "")
	fmt.Println(name)
	return name
}

func CreateTournament(name string) (*Tournament, error) {
	body, err := json.Marshal(map[string]any{
		"api_key": API_KEY,
		"tournament": map[string]any{
			"name": name,
			// "url":  formatURL(name), // Returns Unprocessable Entity if the URL is already taken
			"signup_cap": 32,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("Error creating JSON: %w", err)
	}

	resp, err := http.Post(BASE_URL+"tournaments.json", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("Error creating tournament: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error creating tournament: %s", resp.Status)
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %w", err)
	}

	// Unmarshal the JSON response into the TournamentResponse struct
	var tr TournamentResponse
	err = json.Unmarshal(bodyBytes, &tr)
	if err != nil {
		return nil, fmt.Errorf("Error parsing JSON response: %w", err)
	}

	return &tr.Tournament, nil
}

func AddPlayer() {

}
