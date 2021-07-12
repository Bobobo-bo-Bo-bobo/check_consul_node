package main

import (
	"fmt"
	"net/url"
	"strings"
)

func validateConsulURL(u string) (string, error) {
	_url, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	if strings.ToLower(_url.Scheme) != "http" && strings.ToLower(_url.Scheme) != "https" {
		return "", fmt.Errorf("Invalid scheme %s, only http and https are supported", _url.Scheme)
	}

	port := _url.Port()
	if port != "" {
		port = ":" + port
	}

	return fmt.Sprintf("%s://%s%s", strings.ToLower(_url.Scheme), strings.ToLower(_url.Hostname()), port), nil
}
