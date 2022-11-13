package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")
	flag.Parse()
	if *helpPtr {
		help()
	}
	input := ScanTarget()
	if !IsURL(input) {
		fmt.Println("Please enter a valid url.")
		os.Exit(1)
	}
	redirects := ScanRedirects(input)
	fmt.Println()
	for _, elem := range redirects {
		fmt.Println("[>] " + elem.URL + " " + elem.Code)
		fmt.Println()
	}
}

// help shows the usage.
func help() {
	fmt.Println()
	var usage = `Take as input a URL and print on stdout all the redirects.
	$> chainredir http://example.com`
	fmt.Println(usage)
	fmt.Println()
	os.Exit(0)
}

// ScanTarget return the element
// taken as input.
func ScanTarget() string {
	if len(os.Args) < 2 {
		fmt.Println("usage: chainredir <url-here>")
		os.Exit(1)
	}
	input := os.Args[1]
	return input
}

// Redirect Struct.
type Redirect struct {
	URL  string
	Code string
}

// ScanRedirects.
func ScanRedirects(input string) []Redirect {
	result := []Redirect{}
	nextURL := input
	var i int
	for i < 1000 {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}}

		if len(nextURL) == 0 {
			break
		}
		if nextURL[0] == '/' {
			nextURL = ExtractHost(result[len(result)-1].URL) + nextURL
		}
		resp, err := client.Get(nextURL)

		if err != nil {
			panic(err)
		}

		if resp.StatusCode == 200 {
			output := Redirect{URL: resp.Request.URL.String(), Code: resp.Status}
			result = append(result, output)
			break
		} else {
			nextURL = resp.Header.Get("Location")
			output := Redirect{URL: resp.Request.URL.String(), Code: resp.Status}
			result = append(result, output)
			i += 1
		}
	}
	return result
}

// IsURL.
func IsURL(input string) bool {
	u, err := url.Parse(input)
	if err != nil {
		panic(err)
	}
	if u.Scheme != "" && u.Host != "" {
		return true
	}
	return false
}

// ExtractHost.
func ExtractHost(input string) string {
	u, err := url.Parse(input)
	if err != nil {
		return ""
	}
	return u.Scheme + u.Host
}
