//Aram Maljanian
package shodan

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type DNSInfo struct {
	HostIP string `json:"hostname_ip"`
}

type HostIP struct {
	Matches []DNSInfo `json:"matches"`
}

func (s *Client) DNSInfo(q string) (string, error) {
	res, err := http.Get(
		fmt.Sprintf("https://api.shodan.io/dns/resolve?hostnames=%s&key=%s", q, s.apiKey),
	)
	fmt.Printf("\n\n")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	//used this idea: https://www.codegrepper.com/code-examples/go/golang+print+http+request+body
	bodyBytes, err := ioutil.ReadAll(res.Body)
	bodyString := string(bodyBytes)

	return bodyString, nil
}
