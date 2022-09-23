package requests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/guilherme-marcello/janus/elements"
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

type MODEL_KEEPALIVE_SESSION struct {
	Janus       string `json:"janus"`
	SessionID   int64  `json:"session_id"`
	Transaction string `json:"transaction"`
}

func KEEPALIVE_SESSION() map[string]string {
	return map[string]string{
		"janus": "keepalive", "transaction": uuid.New().String(),
	}
}

type MODEL_ATTACH_PLUGIN struct {
	Janus       string `json:"janus"`
	SessionID   int64  `json:"session_id"`
	Transaction string `json:"transaction"`
	Data        struct {
		ID int64 `json:"id"`
	} `json:"data"`
}

func ATTACH_PLUGIN(pluginName string) map[string]string {
	return map[string]string{
		"janus": "attach", "transaction": uuid.New().String(),
		"plugin": pluginName,
	}
}

type MODEL_LIST_RECORDPLAY struct {
	Janus       string `json:"janus"`
	SessionID   int64  `json:"session_id"`
	Transaction string `json:"transaction"`
	Sender      int64  `json:"sender"`
	Plugindata  struct {
		Plugin string `json:"plugin"`
		Data   struct {
			Recordplay string               `json:"recordplay"`
			List       []elements.Recording `json:"list"`
		} `json:"data"`
	} `json:"plugindata"`
}

func LIST_RECORDPLAY() map[string]any {
	return map[string]any{
		"janus": "message", "transaction": uuid.New().String(),
		"body": map[string]string{
			"request": "list",
		},
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
