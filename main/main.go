package main

import (
	"flag"

	"github.com/aler9/gomavlib"
	"github.com/koustech/mastermind/telemetry"

	"github.com/koustech/mastermind/server"
	u "github.com/koustech/mastermind/utils"
)

func main() {
	u.InitializeLoggers()
	defer u.SyncLogger()

	var grpcAddress string
	var mavlinkAddress string
	var test bool
	flag.StringVar(&grpcAddress, "a", "0.0.0.0:5678", "Mastermind's gRPC server address")
	flag.StringVar(&mavlinkAddress, "m", "127.0.0.1:14556", "The source MAVLink vehicle's ip address")
	flag.BoolVar(&test, "test", false, "Disables telemetry mode")
	flag.Parse()

	var node *gomavlib.Node

	var sysId uint8 = 0
	var compId uint8 = 0

	if test {
		u.Logger.Info("Running in test mode")
		node = &gomavlib.Node{}
	} else {
		u.Logger.Info("Running in telemetry mode")
		node, sysId, compId = telemetry.ConnectToVehicle(mavlinkAddress)
	}
	defer node.Close()

	if err := server.Run(grpcAddress, node, sysId, compId); err != nil {
		u.Logger.Fatal(err)
	}
}
