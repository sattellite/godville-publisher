package client

import (
	"errors"
	"io"
	"math/rand/v2"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const proxyList = "https://raw.githubusercontent.com/vakhov/fresh-proxy-list/refs/heads/master/http.txt"

func HTTP(withProxy bool) *http.Client {
	if withProxy {
		proxyURL, err := getProxyURL()
		if err != nil {
			return http.DefaultClient
		}
		return &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL), // use environment proxy settings
			},
		}
	}
	return http.DefaultClient
}

func validateProxy(p string) bool {
	pu, err := url.Parse("http://" + p)
	if err != nil {
		return false
	}

	c := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(pu),
		},
	}

	resp, err := c.Get("https://godville.net/gods/api/admin")
	if err != nil {
		return false
	}
	resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func getProxyURL() (*url.URL, error) {
	resp, err := http.DefaultClient.Get(proxyList)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	list := make([]string, 0, 2000)
	for _, line := range strings.Split(string(body), "\n") {
		// remove control characters
		list = append(list, strings.Trim(line, "\r\n\t "))
	}

	// select a random proxy from the list
	p := randomItem(list)
	if p == "" {
		return nil, errors.New("no proxy available")
	}

	for !validateProxy(p) {
		p = randomItem(list)
	}

	return url.Parse("http://" + p)
}

func randomItem(list []string) string {
	if len(list) == 0 {
		return ""
	}

	return list[rand.Int64N(int64(len(list)-1))]
}
