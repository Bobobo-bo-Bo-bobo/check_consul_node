package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func httpRequest(u string, method string, header map[string]string, reader io.Reader, insecureSsl bool, ca []byte) (HTTPResult, error) {
	var result HTTPResult
	var transp *http.Transport

	if insecureSsl {
		transp = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	} else {
		transp = &http.Transport{
			TLSClientConfig: &tls.Config{},
		}
		if ca != nil {
			_ca := x509.NewCertPool()
			ok := _ca.AppendCertsFromPEM(ca)
			if !ok {
				return result, fmt.Errorf("Can't add CA to CA chain")
			}
			transp.TLSClientConfig.RootCAs = _ca
		}
	}

	client := &http.Client{
		Transport: transp,
	}

	result.URL = u

	request, err := http.NewRequest(method, u, reader)
	if err != nil {
		return result, err
	}

	defer func() {
		if request.Body != nil {
			ioutil.ReadAll(request.Body)
			request.Body.Close()
		}
	}()

	// add required headers
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	// set User-Agent
	request.Header.Set("User-Agent", userAgent)

	// close connection after response and prevent re-use of TCP connection because some implementations (e.g. HP iLO4)
	// don't like connection reuse and respond with EoF for the next connections
	request.Close = true

	// add supplied additional headers
	if header != nil {
		for key, value := range header {
			request.Header.Add(key, value)
		}
	}

	response, err := client.Do(request)
	if err != nil {
		return result, err
	}

	defer func() {
		ioutil.ReadAll(response.Body)
		response.Body.Close()
	}()

	result.Status = response.Status
	result.StatusCode = response.StatusCode
	result.Header = response.Header
	result.Content, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	return result, nil
}
