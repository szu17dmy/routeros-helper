package addrlists

import (
	"bufio"
	"errors"
	"net"
	"net/http"
	"os"
	"strings"
)

const (
	key = "ADDRESS_LIST_URL"
)

func Load() ([]string, error) {
	url := os.Getenv(key)
	if url == "" {
		return nil, errors.New("failed to get the address list url")
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var l []string
	s := bufio.NewScanner(resp.Body)
	for s.Scan() {
		i, err := parse(s.Text())
		if err != nil {
			continue
		}
		l = append(l, i)
	}
	if s.Err() != nil {
		return nil, s.Err()
	}
	l = append(l, "")
	return l, nil
}

func parse(cidr string) (string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		ip = net.ParseIP(cidr)
		if ip == nil {
			return "", err
		} else {
			return ip.String(), nil
		}
	}
	i := ipnet.String()
	if strings.HasSuffix(i, "/32") {
		return ip.String(), nil
	}
	return i, nil
}
