package telemetry

import (
	"github.com/aler9/gomavlib"
	"github.com/aler9/gomavlib/pkg/dialects/ardupilotmega"
	"github.com/koustech/mastermind/utils"
)

func ConnectToVehicle(address string) *gomavlib.Node {
	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			gomavlib.EndpointUDPServer{Address: address},
		},
		Dialect:     ardupilotmega.Dialect,
		OutVersion:  gomavlib.V2, // change to V1 if you're unable to communicate with the target
		OutSystemID: 10,
	})
	if err != nil {
		utils.Logger.Fatal(err)
	}
	awaitHeartbeat := func() {
		for evt := range node.Events() {
			switch evt.(type) {
			case *gomavlib.EventChannelOpen:
				utils.Logger.Infof("received heartbeat from %s", address)
				return
			}
		}
	}

	awaitHeartbeat()
	return node
}

// func GetTelem(bus evbus.Bus, node *gomavlib.Node) {
// 	vfrHud := ardupilotmega.MessageVfrHud{}
// 	attitude := ardupilotmega.MessageAttitude{}
// 	globalPositionInt := ardupilotmega.MessageGlobalPositionInt{}

// 	updateFlag := true
// 	// print every message we receive
// 	for evt := range node.Events() {
// 		updateFlag = true

// 		if frm, ok := evt.(*gomavlib.EventFrame); ok {
// 			switch msg := frm.Message().(type) {
// 			case *ardupilotmega.MessageAttitude:
// 				attitude = *msg
// 			case *ardupilotmega.MessageVfrHud:
// 				vfrHud = *msg
// 			case *ardupilotmega.MessageGlobalPositionInt:
// 				globalPositionInt = *msg
// 			default:
// 				updateFlag = false
// 			}

// 			if updateFlag {
// 				bus.Publish(&NewTelemetryEvent)
// 			}
// 		}
// 	}
// }
