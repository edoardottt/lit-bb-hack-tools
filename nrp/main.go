package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := ScanTargets()
	var result []string
	for _, elem := range input {
		finalUrl := ScanRedirect(elem)
		if finalUrl.Url != "" {
			final := finalUrl.Url + " " + strconv.Itoa(finalUrl.Code)
			result = append(result, final)
		}

	}
	for _, elem := range removeDuplicateValues(result) {
		fmt.Println(elem)
	}
}

//ScanTargets return the array of elements
//taken as input on stdin.
func ScanTargets() []string {

	var result []string
	// accept domains on stdin
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if !isUrl(sc.Text()) {
			fmt.Println(sc.Text() + " is not a proper url")
			os.Exit(1)
		}
		domain := strings.ToLower(sc.Text())
		result = append(result, domain)
	}
	return result
}

//Redirect Struct
type Redirect struct {
	Url  string
	Code int
}

//ScanRedirect
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
			nextURL = ExtractHost(result[len(result)-1].Url) + nextURL
		}
		resp, err := client.Get(nextURL)

		if err != nil {
			return Redirect{"", 1}
		}

		if resp.StatusCode == 200 {
			output := Redirect{Url: resp.Request.URL.String(), Code: resp.StatusCode}
			result = append(result, output)
			break
		} else {
			nextURL = resp.Header.Get("Location")
			output := Redirect{Url: resp.Request.URL.String(), Code: resp.StatusCode}
			result = append(result, output)
			i += 1
		}
	}
	return result[len(result)-1]
}

//isUrl >
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

//ExtractHost >
func ExtractHost(input string) string {
	u, err := url.Parse(input)
	if err != nil {
		return ""
	}
	return u.Scheme + u.Host
}

//removeDuplicateValues
func removeDuplicateValues(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
