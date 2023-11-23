package addrlists

import (
	"bufio"
	"errors"
	"net/http"
	"os"
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
		l = append(l, s.Text())
	}
	if s.Err() != nil {
		return nil, s.Err()
	}
	l = append(l, "")
	return l, nil
}
