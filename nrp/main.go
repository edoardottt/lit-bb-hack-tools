package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")

	flag.Parse()

	if *helpPtr {
		help()
	}

	result := sync.Map{}

	input := ScanTargets()
	limiter := make(chan string, 30) // Limits simultaneous requests.
	wg := sync.WaitGroup{}           // Needed to not prematurely exit before all requests have been finished.

	for _, elem := range input {
		limiter <- elem

		wg.Add(1)

		go func(elem string) {
			defer wg.Done()
			defer func() { <-limiter }()

			finalURL := ScanRedirect(elem)
			if finalURL.URL != "" {
				final := finalURL.URL + " " + strconv.Itoa(finalURL.Code)
				if _, ok := result.Load(final); !ok {
					result.Store(final, true)
					fmt.Println(final)
				}
			}
		}(elem)
	}

	wg.Wait()
}

// help shows the usage.
func help() {
	var usage = `Take as input on stdin a list of domains and print on stdout all the unique domains without redirects. 
For example, if two domains (A and B) redirects to the same domain C, the output will be C.
	$> cat urls | nrp`

	fmt.Println()
	fmt.Println(usage)
	fmt.Println()
	os.Exit(0)
}

// ScanTargets return the array of elements
// taken as input on stdin.
func ScanTargets() []string {
	var result []string

	// accept domains on stdin.
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if !IsURL(sc.Text()) {
			continue
		}

		domain := strings.ToLower(sc.Text())
		result = append(result, domain)
	}

	return result
}

// Redirect Struct.
type Redirect struct {
	URL  string
	Code int
}

// ScanRedirect.
func ScanRedirect(input string) Redirect {
	result := []Redirect{{"", 1}}
	nextURL := input

	var i int
	for i < 10 {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}, Transport: tr}

		if len(nextURL) == 0 {
			break
		}

		if nextURL[0] == '/' {
			nextURL = ExtractHost(result[len(result)-1].URL) + nextURL
		}

		resp, err := client.Get(nextURL)
		if err != nil {
			return Redirect{"", 1}
		}

		if resp.StatusCode == 200 {
			output := Redirect{URL: resp.Request.URL.String(), Code: resp.StatusCode}
			result = append(result, output)

			break
		} else {
			nextURL = resp.Header.Get("Location")
			output := Redirect{URL: resp.Request.URL.String(), Code: resp.StatusCode}
			result = append(result, output)
			i += 1
		}

		resp.Body.Close()
	}

	return result[len(result)-1]
}

// IsURL.
func IsURL(input string) bool {
	u, err := url.Parse(input)
	if err != nil {
		return false
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
