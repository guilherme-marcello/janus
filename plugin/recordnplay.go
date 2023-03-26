package plugin

import (
	"encoding/json"
	"log"

	"github.com/guilherme-marcello/janus/elements"
	"github.com/guilherme-marcello/janus/requests"
	"github.com/guilherme-marcello/janus/session"
)

type RecordPlay struct {
	Plugin
}

func NewRecordPlayHandler(session session.Session) (RecordPlay, error) {
	plugin, err := newPluginHandler(session, "janus.plugin.recordplay")
	if err != nil {
		return RecordPlay{}, err
	}
	return RecordPlay{plugin}, nil
}

func (_recordplay RecordPlay) List() ([]elements.Recording, error) {
	response, err := requests.SendHTTPRequest("POST", _recordplay.endpoint, requests.LIST_REQUEST())
	if err != nil {
		return nil, err
	}
	model := &requests.MODEL_LIST_RECORDPLAY{}
	err = json.NewDecoder(response.Body).Decode(&model)
	if err != nil || model.Janus == "error" {
		log.Printf("Failed list recordings with plugin %s at %s", _recordplay.name, _recordplay.session.Endpoint)
	}
	return model.Plugindata.Data.List, nil
}
