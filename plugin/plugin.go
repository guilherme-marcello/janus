package plugin

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/guilherme-marcello/janus/elements"
	"github.com/guilherme-marcello/janus/requests"
	"github.com/guilherme-marcello/janus/session"
)

type Plugin struct {
	session  session.Session
	name     string
	id       string
	endpoint string
}

func newPluginHandler(session session.Session, name string) Plugin {
	plugin := Plugin{
		session: session,
		name:    name,
	}
	plugin.id = plugin.getId()
	plugin.endpoint = plugin.getEndpoint()
	return plugin
}

func (_plugin Plugin) getId() string {
	response := requests.GetPostResponse(_plugin.session.Endpoint, requests.ATTACH_PLUGIN(_plugin.name))
	model := &requests.MODEL_ATTACH_PLUGIN{}
	err := json.NewDecoder(response.Body).Decode(&model)
	if err != nil || model.Janus == "error" {
		log.Printf("Failed to attach plugin %s at %s", _plugin.name, _plugin.session.Endpoint)
	}
	return strconv.FormatInt(model.Data.ID, 10)
}

func (_plugin Plugin) getEndpoint() string {
	return _plugin.session.Endpoint + "/" + _plugin.id
}

func (_plugin Plugin) String() string {
	return fmt.Sprintf(
		"|%s|\nPlugin Handle Id: %s\nPlugin handle endpoint: %s",
		_plugin.name,
		_plugin.id,
		_plugin.endpoint,
	)
}

type RecordPlay struct {
	Plugin
}

func NewRecordPlayHandler(session session.Session) RecordPlay {
	return RecordPlay{
		newPluginHandler(session, "janus.plugin.recordplay"),
	}
}

func (_recordplay RecordPlay) List() []elements.Recording {
	response := requests.GetPostResponse(_recordplay.endpoint, requests.LIST_REQUEST())
	model := &requests.MODEL_LIST_RECORDPLAY{}
	err := json.NewDecoder(response.Body).Decode(&model)
	if err != nil || model.Janus == "error" {
		log.Printf("Failed list recordings with plugin %s at %s", _recordplay.name, _recordplay.session.Endpoint)
	}
	return model.Plugindata.Data.List
}

type Streaming struct {
	Plugin
}

func NewStreamingHandler(session session.Session) Streaming {
	return Streaming{
		newPluginHandler(session, "janus.plugin.streaming"),
	}
}

func (_streaming Streaming) List() []elements.Mountpoint {
	response := requests.GetPostResponse(_streaming.endpoint, requests.LIST_REQUEST())
	model := &requests.MODEL_LIST_STREAMING{}
	err := json.NewDecoder(response.Body).Decode(&model)
	if err != nil || model.Janus == "error" {
		log.Printf("Failed to list mountpoints with plugin %s at %s", _streaming.name, _streaming.session.Endpoint)
	}
	return model.Plugindata.Data.List
}
