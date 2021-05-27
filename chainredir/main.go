package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	input := ScanTarget()
	if !isUrl(input) {
		fmt.Println("Please enter a valid url.")
		os.Exit(1)
	}
	redirects := ScanRedirects(input)
	fmt.Println()
	for _, elem := range redirects {
		fmt.Println("[>] " + elem.Url + " " + elem.Code)
		fmt.Println()
	}
}

//ScanTarget return the element
//taken as input.
func ScanTarget() string {
	input := os.Args[1]
	return input
}

//Redirect Struct
type Redirect struct {
	Url  string
	Code string
}

//ScanRedirects
func ScanRedirects(input string) []Redirect {
	result := []Redirect{}
	nextURL := input
	var i int
	for i < 1000 {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}}

		resp, err := client.Get(nextURL)

		if err != nil {
			panic(err)
		}

		if resp.StatusCode == 200 {
			output := Redirect{Url: resp.Request.URL.String(), Code: resp.Status}
			result = append(result, output)
			break
		} else {
			nextURL = resp.Header.Get("Location")
			fmt.Println(nextURL)
			output := Redirect{Url: resp.Request.URL.String(), Code: resp.Status}
			result = append(result, output)
			i += 1
		}
	}
	return result
}

//isUrl
func isUrl(input string) bool {
	u, err := url.Parse(input)
	if err != nil {
		panic(err)
	}
	if u.Scheme != "" && u.Host != "" {
		return true
	}
	return false
}
