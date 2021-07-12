package main

import (
	"fmt"
)

func parseNodeStatus(chs ConsulHealthStatus) (int, string) {
	var msg string
	var state int = NagiosUnknown

	// Status strings are defined in api/health.go of the Consul source
	msg = fmt.Sprintf("Node reports %s state: %s", chs.Status, chs.Output)

	switch chs.Status {
	case "passing":
		msg = "OK - " + msg
		state = NagiosOK
	case "warning":
		msg = "WARNING - " + msg
		state = NagiosWarning
	case "critical":
		msg = "CRITICAL - " + msg
		state = NagiosCritical
	case "maintenance":
		msg = "WARNING - " + msg
		state = NagiosWarning
	default:
		msg = fmt.Sprintf("BUG - Can't interpret status %s", chs.Status)
		state = NagiosUnknown
	}

	return state, msg
}
