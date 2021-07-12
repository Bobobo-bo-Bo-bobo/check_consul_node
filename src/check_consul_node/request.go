package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func requestNodeStatus(u string, node string, token string, ns string, dc string, caf string, insecure bool) ([]ConsulHealthStatus, error) {
	var result []ConsulHealthStatus
	var head = make(map[string]string)
	var ca []byte
	var err error

	if caf != "" {
		ca, err = ioutil.ReadFile(caf)
		if err != nil {
			return result, err
		}
	}

	if ns != "" {
		head["X-Consul-Namespace"] = ns
	}

	if token != "" {
		head["X-Consul-Token"] = token
	}

	if len(head) == 0 {
		head = nil
	}

	url := fmt.Sprintf("%s%s%s", u, consulHealthURL, node)
	if dc != "" {
		url += "?dc=" + dc
	}

	reply, err := httpRequest(url, "GET", head, nil, insecure, ca)
	if err != nil {
		return result, err
	}

	if reply.StatusCode != http.StatusOK {
		return result, fmt.Errorf("Invalid HTTP status; Expected \"200 OK\" but got \"%s\" instead", reply.Status)
	}

	err = json.Unmarshal(reply.Content, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
