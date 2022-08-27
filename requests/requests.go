package requests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type MODEL_CREATE_SESSION struct {
	Janus       string `json:"janus"`
	Transaction string `json:"transaction"`
	Data        struct {
		ID int64 `json:"id"`
	} `json:"data"`
}

func CREATE_SESSION() map[string]string {
	return map[string]string{
		"janus": "create", "transaction": uuid.New().String(),
	}
}

type MODEL_DESTROY_SESSION struct {
	Janus       string `json:"janus"`
	SessionID   int64  `json:"session_id"`
	Transaction string `json:"transaction"`
}

func DESTROY_SESSION() map[string]string {
	return map[string]string{
		"janus": "destroy", "transaction": uuid.New().String(),
	}
}

func GetPostResponse(endpoint string, payload any) *http.Response {
	json_data, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
