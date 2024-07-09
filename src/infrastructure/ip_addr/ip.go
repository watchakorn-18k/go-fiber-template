package ipAddr

import (
	"io"
	"net/http"

	"github.com/goccy/go-json"
)

type IP struct {
	IP string `json:"ip"`
}

func GetIp() (string, error) {
	url := "https://api.ipify.org?format=json"

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var ip IP

	err = json.Unmarshal(body, &ip)
	if err != nil {
		return "", err
	}
	return ip.IP, nil
}
