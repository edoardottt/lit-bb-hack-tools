package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")
	flag.Parse()
	if *helpPtr {
		help()
	}
	input := ScanTargets()
	result := GetMetrics(input)
	for _, elem := range result {
		fmt.Println(elem)
	}
}

//help shows the usage
func help() {
	var usage = `Take as input on stdin a list of urls and print on stdout all the unique paths and url found in the /metrics endpoint.
	$> cat urls | kubemetrics`
	fmt.Println()
	fmt.Println(usage)
	fmt.Println()
	os.Exit(0)
}

//ScanTargets return the array of elements
//taken as input on stdin.
func ScanTargets() []string {

	var result []string

	// accept domains on stdin
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		if !HasProtocol(domain) {
			fmt.Println(domain + " has no protocol!")
			os.Exit(1)
		}
		result = append(result, RemovePath(domain))
	}
	return RemoveDuplicateValues(result)
}

//RemoveDuplicateValues >
func RemoveDuplicateValues(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//GetMetrics >
func GetMetrics(input []string) []string {
	var result []string
	var mutex = &sync.Mutex{}
	pathRe := `path\=\".*\"`
	urlRe := `url\=\".*\"`

	limiter := make(chan string, 10) // Limits simultaneous requests
	wg := sync.WaitGroup{}           // Needed to not prematurely exit before all requests have been finished

	for _, elem := range input {
		limiter <- elem
		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			defer func() { <-limiter }()
			metrics := GetRequest(domain + "/metrics")
			mutex.Lock()
			if metrics != "" {
				if matched, err := regexp.Match(pathRe, []byte(metrics)); err == nil && matched {
					re := regexp.MustCompile(pathRe)
					matches := re.FindAllString(metrics, -1)
					for _, match := range matches {
						elem := strings.ReplaceAll(strings.ReplaceAll(string(strings.Split(match, ",")[0]), "path=\"", ""), "\"", "")
						if elem != "/" && strings.Trim(elem, " ") != "" {
							result = append(result, elem)
						}
					}
				}
				if matched, err := regexp.Match(urlRe, []byte(metrics)); err == nil && matched {
					re := regexp.MustCompile(urlRe)
					matches := re.FindAllString(metrics, -1)
					for _, match := range matches {
						elem := strings.ReplaceAll(strings.ReplaceAll(string(strings.Split(match, ",")[0]), "url=\"", ""), "\"", "")
						if elem != "/" && strings.Trim(elem, " ") != "" {
							result = append(result, elem)
						}
					}
				}
			}
			mutex.Unlock()
		}(elem)
	}
	wg.Wait()
	return RemoveDuplicateValues(result)
}

//GetRequest performs a GET request
func GetRequest(target string) string {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(target)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	//Convert the body to type string
	sb := string(body)
	return sb
}

//HasProtocol >
func HasProtocol(input string) bool {
	return strings.Contains(input, "://")
}

//RemovePath >
func RemovePath(input string) string {
	u, err := url.Parse(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return u.Scheme + "://" + u.Host
}
