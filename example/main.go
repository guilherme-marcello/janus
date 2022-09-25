package main

import (
	"fmt"

	"github.com/guilherme-marcello/janus"
	"github.com/guilherme-marcello/janus/elements"
	"github.com/guilherme-marcello/janus/plugin"
	"github.com/guilherme-marcello/janus/session"
)

func main() {
	janusClient := janus.Http{
		Endpoint: "https://janus.conf.meetecho.com/janus",
	}

	janusSession := session.New(janusClient)
	janusSession.KeepAlive()

	streamingPlugin := plugin.NewStreamingHandler(janusSession)
	fmt.Println(streamingPlugin)
    
	var mountpointsList []elements.Mountpoint = streamingPlugin.List()

	fmt.Println("<-- Available mountpoints to watch -->")
	for _, mountpoint := range mountpointsList {
		if mountpoint.Enabled {
			fmt.Printf("[%d] %s\n", mountpoint.ID, mountpoint.Description)
		}
	}

	janusSession.Destroy()
}
