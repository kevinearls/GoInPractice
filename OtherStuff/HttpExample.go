package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	//url := "https://golang.org"
	// --insecure      Allow connections to SSL sites without certs (H)
	url := "https://agent-as-daemonset-fred.192.168.64.66.nip.io/api/services"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := http.Client{Timeout: time.Second, Transport: tr}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	res, err := c.Do(req)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(res.Body)

	fmt.Printf("%s\n", body)

}
