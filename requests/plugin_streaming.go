package requests

import (
	"github.com/guilherme-marcello/janus/elements"
)

type MODEL_LIST_STREAMING struct {
	Janus       string `json:"janus"`
	SessionID   int64  `json:"session_id"`
	Transaction string `json:"transaction"`
	Sender      int64  `json:"sender"`
	Plugindata  struct {
		Plugin string `json:"plugin"`
		Data   struct {
			Streaming string                `json:"streaming"`
			List      []elements.Mountpoint `json:"list"`
		} `json:"data"`
	} `json:"plugindata"`
}
