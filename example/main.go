package main

import (
	"fmt"

	"github.com/guilherme-marcello/janus"
	"github.com/guilherme-marcello/janus/plugin"
	"github.com/guilherme-marcello/janus/session"
)

func main() {
	janusClient := janus.Http{
		Endpoint: "https://janus.conf.meetecho.com/janus",
	}

	janusSession, err := session.New(janusClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	janusSession.KeepAlive()

	streamingPlugin, err := plugin.NewStreamingHandler(janusSession)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(streamingPlugin)

	mountpointsList, err := streamingPlugin.List()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("<-- Available mountpoints to watch -->")
	for _, mountpoint := range mountpointsList {
		if mountpoint.Enabled {
			fmt.Printf("[%d] %s\n", mountpoint.ID, mountpoint.Description)
		}
	}

	janusSession.Destroy()
}
