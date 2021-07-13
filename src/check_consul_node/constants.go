package main

const name = "check_consul_node"
const version = "1.0.0"

var userAgent = name + "/" + version

const defaultConsulURL = "http://localhost:8500"

const helpText = `Usage: %s [--ca-file=<ca>] [--consul-url=<uri>] [--datacenter=<dc>] [--help] [--insecure] [--namespace=<ns>] [--node=<node>] [--token=<token>] [--version]

    --ca-file=<ca>      CA file if the CA is not present in the CA store of the system

    --consul-url=<url>  Consul URL to query.
                        Default is %s

    --datacenter=<dc>   Datacenter to query

    --help              This text

    --insecure          Skip SSL verification

    --namespace=<ns>    Consul namespace to query
                        Requires Consul Enterprise

    --node=<name>       Node name to check
                        Default: current hostname
    
    --token=<token>     Consul token to use
                        Required if Consul ACLs are configured

    --version           Show version information

`

const versionText = `%s version %s
Copyright (C) 2021 by Andreas Maus <maus@ypbind.de>
This program comes with ABSOLUTELY NO WARRANTY.

%s is distributed under the Terms of the GNU General
Public License Version 3. (http://www.gnu.org/copyleft/gpl.html)

Build with go version: %s
`

const (
	// NagiosOK - State is OK
	NagiosOK int = iota
	// NagiosWarning - State is in warning state
	NagiosWarning
	// NagiosCritical - State is in critical state
	NagiosCritical
	// NagiosUnknown - Other error
	NagiosUnknown
)

const consulHealthURL = "/v1/health/node/"
