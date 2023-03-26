package session

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/guilherme-marcello/janus"
	"github.com/guilherme-marcello/janus/requests"
)

type Session struct {
	client   janus.Http
	id       string
	Endpoint string
}

func New(client janus.Http) (Session, error) {
	session := Session{
		client: client,
	}
	var err error
	session.id, err = session.getId()
	if err != nil {
		return Session{}, err
	}
	session.Endpoint = session.getEndpoint()
	return session, nil
}

func (_session Session) getId() (string, error) {
	response, err := requests.SendHTTPRequest("POST", _session.client.Endpoint, requests.CREATE_SESSION())
	if err != nil {
		return "", err
	}
	model := &requests.MODEL_CREATE_SESSION{}
	err = json.NewDecoder(response.Body).Decode(&model)
	if err != nil {
		log.Fatalf("Failed to establish a new session at %s", _session.client.Endpoint)
	}
	return strconv.FormatInt(model.Data.ID, 10), nil
}

func (_session Session) getEndpoint() string {
	return _session.client.Endpoint + "/" + _session.id
}

func (_session Session) Destroy() {
	response, _ := requests.SendHTTPRequest("POST", _session.Endpoint, requests.DESTROY_SESSION())
	model := &requests.MODEL_DESTROY_SESSION{}
	err := json.NewDecoder(response.Body).Decode(&model)
	if err != nil || model.Janus == "error" {
		log.Printf("Failed to destroy session %s at %s", _session.id, _session.client.Endpoint)
	}
}

func (_session Session) KeepAlive() {
	response, _ := requests.SendHTTPRequest("POST", _session.Endpoint, requests.KEEPALIVE_SESSION())
	model := &requests.MODEL_KEEPALIVE_SESSION{}
	err := json.NewDecoder(response.Body).Decode(&model)
	if err != nil || model.Janus != "ack" {
		log.Printf("Failed to keep session %s alive at %s", _session.id, _session.client.Endpoint)
	}
}

func (_session Session) String() string {
	return fmt.Sprintf(
		"Session Id: %s\nSession handle endpoint: %s",
		_session.id,
		_session.Endpoint,
	)
}
