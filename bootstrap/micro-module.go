package bootstrap

import (
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

// MicroModule go micro module helper
type MicroModule struct{}

var microClient client.Client
var publisher = make(map[string]micro.Publisher)

// RegisterClient register micro client for using on functions
func RegisterClient(c client.Client) {
	microClient = c
}

// RegisterPublishers register publisher for using on functions etc. handlers
func RegisterPublishers(topics map[string]string) {
	for k, v := range topics {
		publisher[k] = micro.NewEvent(v, microClient)
	}
}

// Publisher get publisher for publish message to messaging service
func (ctl *MicroModule) Publisher(s string) micro.Publisher {
	if p, ok := publisher[s]; ok {
		return p
	}
	panic(fmt.Sprintf("publisher %s not found", s))
}

// Client get micro client
func (ctl *MicroModule) Client() client.Client {
	return microClient
}
