package telemetry

import (
	"github.com/aler9/gomavlib"
	"github.com/aler9/gomavlib/pkg/dialects/ardupilotmega"
	"github.com/koustech/mastermind/utils"
)

func GetTelem(vehicleConnected chan<- struct{}, address string) {
	// create a node which
	// - communicates with an UDP endpoint in server mode.
	// - understands ardupilotmega dialect
	// - writes messages with given system id
	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			gomavlib.EndpointUDPServer{Address: address},
		},
		Dialect:     ardupilotmega.Dialect,
		OutVersion:  gomavlib.V2, // change to V1 if you're unable to communicate with the target
		OutSystemID: 10,
	})
	if err != nil {
		panic(err)
	}
	defer node.Close()

	awaitHeartbeat := func() {
		for evt := range node.Events() {
			switch evt.(type) {
			case *gomavlib.EventChannelOpen:
				vehicleConnected <- struct{}{}
				utils.Logger.Infof("received heartbeat from %s", address)
				return
			}
		}
	}

	awaitHeartbeat()

	// print every message we receive
	for evt := range node.Events() {
		if frm, ok := evt.(*gomavlib.EventFrame); ok {
			switch msg := frm.Message().(type) {
			case *ardupilotmega.MessageAttitude:
				utils.Logger.Debugf("%+v\n", msg)
			}
		}
	}
}
