package telemetry

import (
	"github.com/aler9/gomavlib"
	"github.com/aler9/gomavlib/pkg/dialects/ardupilotmega"
	"github.com/koustech/mastermind/utils"
)

func ConnectToVehicle(address string) (node *gomavlib.Node, systemId uint8, componentId uint8) {
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
				return
			}
		}
	}
	getSysDetails := func() (uint8, uint8) {
		var sysId uint8
		var compId uint8
		for evt := range node.Events() {
			if frm, ok := evt.(*gomavlib.EventFrame); ok {
				sysId = frm.SystemID()
				compId = frm.ComponentID()
				break
			}
		}
		return sysId, compId
	}

	awaitHeartbeat()
	sysId, compId := getSysDetails()
	utils.Logger.Infof("received heartbeat from %s sysid: %v cid: %v", address, sysId, compId)
	return node, sysId, compId
}
