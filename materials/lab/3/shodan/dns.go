package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DNSInfo struct {
	HostIP string `json:"hostname_ip"`
}

type HostIP struct {
	Matches []DNSInfo `json:"matches"`
}

func (s *Client) DNSInfo(q string) (*HostIP, error) {
	res, err := http.Get(
		fmt.Sprintf("%s/shodan/dns/resolve?hostnames=%s&key=%s", BaseURL, q, s.apiKey),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret HostIP
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
