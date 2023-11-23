package main

import (
	"log"
	"net/http"
	"strconv"

	"szu17dmy/routeros-helper/pkg/ip/firewall/addrlists"
)

var lists []string

func handler(w http.ResponseWriter, r *http.Request) {
	begin, end, err := parse(r.URL.Query().Get("page"), r.URL.Query().Get("size"), len(lists))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err, r.URL)
		return
	}
	var buf []byte
	for _, addr := range lists[begin:end] {
		buf = append(buf, []byte(addr+"\n")...)
	}
	_, err = w.Write(buf)
	if err != nil {
		log.Print(err, buf)
		return
	}
}

func parse(page, size string, limit int) (int, int, error) {
	if page == "" && size == "" {
		return 0, 100, nil
	}
	p, err := strconv.Atoi(page)
	if err != nil {
		return 0, 0, err
	}
	s, err := strconv.Atoi(size)
	if err != nil {
		return 0, 0, err
	}
	begin := p * s
	end := (p + 1) * s
	if begin < 0 {
		begin = 0
	}
	if begin > limit {
		begin = limit - 1
	}
	if end < 0 {
		end = 0
	}
	if end > limit {
		end = limit
	}
	return begin, end, nil
}

func init() {
	l, err := addrlists.Load()
	if err != nil {
		log.Fatal(err)
	}
	lists = l
}

func main() {
	http.HandleFunc("/ip/firewall/address-lists", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
