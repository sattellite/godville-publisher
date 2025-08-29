package client

import (
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const proxyList = "https://raw.githubusercontent.com/vakhov/fresh-proxy-list/refs/heads/master/http.txt"

func HTTP(withProxy bool) *http.Client {
	fmt.Println("With proxy:", withProxy)
	if withProxy {
		proxyURL, err := getProxyURL()
		if err != nil {
			fmt.Println("Failed to get proxy URL:", err)
			return http.DefaultClient
		}
		fmt.Println("Using proxy:", proxyURL)
		return &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL), // use environment proxy settings
			},
		}
	}
	return http.DefaultClient
}

func validateProxy(p string) bool {
	fmt.Println("Validating proxy:", p)
	pu, err := url.Parse("http://" + p)
	if err != nil {
		fmt.Println("Failed to parse proxy URL:", err)
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
		fmt.Println("Proxy error:", err)
		return false
	}
	resp.Body.Close()

	fmt.Println("Proxy status code:", resp.StatusCode)
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
		fmt.Println("Failed to read proxy list:", readErr)
		return nil, readErr
	}

	list := make([]string, 0, 2000)
	for _, line := range strings.Split(string(body), "\n") {
		// remove control characters
		list = append(list, strings.Trim(line, "\r\n\t "))
	}
	fmt.Println("Loaded proxies:", len(list))

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
