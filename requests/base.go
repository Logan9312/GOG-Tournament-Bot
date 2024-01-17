package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const BASE_URL = "https://api.challonge.com/v1/"

var API_KEY string

func CreateTournament(name string) error {
	body, err := json.Marshal(map[string]any{
		"api_key": API_KEY,
		"tournament": map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return fmt.Errorf("Error creating JSON: %w", err)
	}

	resp, err := http.Post(BASE_URL+"tournaments.json", "application/json", bytes.NewBuffer(body))
	if err != nil {

		return fmt.Errorf("Error creating tournament: %w", err)
	}

	fmt.Println("Response:", resp)

	return nil
}

func AddPlayer() {

}
