package plugin

import (
	"encoding/json"
	"log"

	"github.com/guilherme-marcello/janus/elements"
	"github.com/guilherme-marcello/janus/requests"
	"github.com/guilherme-marcello/janus/session"
)

type Streaming struct {
	Plugin
}

func NewStreamingHandler(session session.Session) (Streaming, error) {
	plugin, err := newPluginHandler(session, "janus.plugin.streaming")
	if err != nil {
		return Streaming{}, err
	}
	return Streaming{plugin}, nil
}

func (_streaming Streaming) List() ([]elements.Mountpoint, error) {
	response, err := requests.SendHTTPRequest("POST", _streaming.endpoint, requests.LIST_REQUEST())
	if err != nil {
		return nil, err
	}
	model := &requests.MODEL_LIST_STREAMING{}
	err = json.NewDecoder(response.Body).Decode(&model)
	if err != nil || model.Janus == "error" {
		log.Printf("Failed to list mountpoints with plugin %s at %s", _streaming.name, _streaming.session.Endpoint)
	}
	return model.Plugindata.Data.List, nil
}
