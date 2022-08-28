package main

import (
	"fmt"

	"github.com/guilherme-marcello/janus"
	"github.com/guilherme-marcello/janus/session"
)

func main() {
	janusClient := janus.Http{
		Endpoint: "https://janus.conf.meetecho.com/janus",
	}

	janusSession := session.New(janusClient)
	fmt.Println(janusSession)
	janusSession.Destroy()
}
