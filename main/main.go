package main

import (
	"flag"

	"github.com/koustech/mastermind/server"
	"github.com/koustech/mastermind/telemetry"
	u "github.com/koustech/mastermind/utils"
)

func main() {
	u.InitializeLoggers()
	defer u.SyncLogger()

	var grpcAddress string
	var mavlinkAddress string
	flag.StringVar(&grpcAddress, "a", "0.0.0.0:5678", "Mastermind's gRPC server address")
	flag.StringVar(&mavlinkAddress, "m", "127.0.0.1:14556", "The source MAVLink vehicle's ip address")
	flag.Parse()

	node := telemetry.ConnectToVehicle(mavlinkAddress)
	defer node.Close()

	if err := server.Run(grpcAddress, node); err != nil {
		u.Logger.Fatal(err)
	}
}
