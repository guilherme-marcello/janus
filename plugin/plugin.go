package plugin

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/guilherme-marcello/janus/requests"
	"github.com/guilherme-marcello/janus/session"
)

type Plugin struct {
	session  session.Session
	name     string
	id       string
	endpoint string
}

func New(session session.Session, name string) Plugin {
	plugin := Plugin {
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
		"Plugin Handle Id: %s\nPlugin handle endpoint: %s",
		_plugin.id,
		_plugin.endpoint,
	)
}
