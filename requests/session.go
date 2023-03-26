package requests

import (
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
