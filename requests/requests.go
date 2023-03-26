package requests

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func LIST_REQUEST() map[string]any {
	return map[string]any{
		"janus": "message", "transaction": uuid.New().String(),
		"body": map[string]string{
			"request": "list",
		},
	}
}

func SendHTTPRequest(method string, endpoint string, payload interface{}) (*http.Response, error) {
	json_data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
