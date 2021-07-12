package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var help = flag.Bool("help", false, "Show help text")
	var version = flag.Bool("version", false, "Show version information")
	var consul = flag.String("consul-url", defaultConsulURL, "Consul URL")
	var datacenter = flag.String("datacenter", "", "Datacenter to query")
	var namespace = flag.String("namespace", "", "Consul namespace to query")
	var _node = flag.String("node", "", "Node name")
	var token = flag.String("token", "", "ACL token")
	var caFile = flag.String("caFile", "", "CA file")
	var insecure = flag.Bool("insecure", false, "Skip SSL verification")

	var node string

	flag.Usage = showUsage
	flag.Parse()

	if *help {
		showUsage()
		os.Exit(0)
	}

	if *version {
		showVersion()
		os.Exit(0)
	}

	consulURL, err := validateConsulURL(*consul)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(NagiosUnknown)
	}

	if *_node == "" {
		node, err = os.Hostname()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Can't get hostname - %s\n", err)
			os.Exit(NagiosUnknown)
		}
	} else {
		node = *_node
	}

	// func requestNodeStatus(u string, node string, token string, ns string, dc string, caf string, insecure bool) (ConsulHealth, error)
	ch, err := requestNodeStatus(consulURL, node, *token, *namespace, *datacenter, *caFile, *insecure)
	if err != nil {
		fmt.Printf("UNKNOWN - Can't fetch node health: %s\n", err)
		os.Exit(NagiosUnknown)
	}

	if len(ch) == 0 {
		fmt.Printf("UNKNOWN - No health data returned for node %s\n", node)
		os.Exit(NagiosUnknown)
	}

	state, msg := parseNodeStatus(ch[0])
	fmt.Println(msg)
	os.Exit(state)
}
