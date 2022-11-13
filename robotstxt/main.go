package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/edoardottt/golazy"
)

func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")

	flag.Parse()

	if *helpPtr {
		help()
	}

	input := ScanTargets()
	result := GetRobots(input)

	for _, elem := range result {
		fmt.Println(elem)
	}
}

// help shows the usage.
func help() {
	var usage = `Take as input on stdin a list of urls and print on stdout all the unique paths found in the robots.txt file.
	$> cat urls | robotstxt`

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
		domain := strings.ToLower(sc.Text())
		result = append(result, domain)
	}

	return golazy.RemoveDuplicateValues(result)
}

// GetRobots.
func GetRobots(input []string) []string {
	var (
		result = []string{}
		mutex  = &sync.Mutex{}
	)

	limiter := make(chan string, 10) // Limits simultaneous requests.
	wg := sync.WaitGroup{}           // Needed to not prematurely exit before all requests have been finished.

	for _, elem := range input {
		limiter <- elem

		wg.Add(1)

		go func(domain string) {
			defer wg.Done()
			defer func() { <-limiter }()

			robots := GetRequest("https://" + RemoveProtocol(domain) + "/robots.txt")

			mutex.Lock()

			if robots != "" {
				s := strings.Split(robots, "\n")
				for _, line := range s {
					if strings.Contains(line, "Allow") || strings.Contains(line, "Disallow") {
						word := strings.Split(line, " ")
						if len(word) > 1 && !strings.Contains(word[1], "Disallow") && !strings.Contains(word[1], "Allow") {
							result = append(result, word[1])
						}
					}
				}
			}
			mutex.Unlock()
		}(elem)
	}

	wg.Wait()

	return golazy.RemoveDuplicateValues(result)
}

// GetRequest performs a GET request.
func GetRequest(target string) string {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(target)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	// We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	// Convert the body to type string.
	sb := string(body)

	return sb
}

// RemoveProtocol.
func RemoveProtocol(input string) string {
	res := strings.Index(input, "://")
	if res >= 0 {
		return input[res+3:]
	} else {
		return input
	}
}
