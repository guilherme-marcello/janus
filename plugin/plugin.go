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

func newPluginHandler(session session.Session, name string) (Plugin, error) {
	plugin := Plugin{
		session: session,
		name:    name,
	}
	var err error
	plugin.id, err = plugin.getId()
	if err != nil {
		return plugin, err
	}

	plugin.endpoint = plugin.getEndpoint()
	return plugin, nil
}

func (_plugin Plugin) getId() (string, error) {
	response, err := requests.SendHTTPRequest("POST", _plugin.session.Endpoint, requests.ATTACH_PLUGIN(_plugin.name))
	if err != nil {
		return "", err
	}
	model := &requests.MODEL_ATTACH_PLUGIN{}
	err = json.NewDecoder(response.Body).Decode(&model)
	if err != nil || model.Janus == "error" {
		log.Printf("Failed to attach plugin %s at %s", _plugin.name, _plugin.session.Endpoint)
	}
	return strconv.FormatInt(model.Data.ID, 10), nil
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
