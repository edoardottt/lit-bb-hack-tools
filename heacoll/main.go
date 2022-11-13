package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/edoardottt/golazy"
)

func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")

	flag.Parse()

	if *helpPtr {
		help()
	}

	RetrieveHeaders(ScanTargets())
}

// help shows the usage.
func help() {
	var usage = `Take as input on stdin a list of urls and print on stdout all the unique headers found.
	$> cat urls | heacoll`

	fmt.Println()
	fmt.Println(usage)
	fmt.Println()
	os.Exit(0)
}

// ScanTargets return the array of elements
// taken as input on stdin.
func ScanTargets() []string {
	var result []string

	// accept domains on stdin
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		result = append(result, domain)
	}

	return golazy.RemoveDuplicateValues(result)
}

// RetrieveHeaders.
func RetrieveHeaders(input []string) {
	var (
		result = make(map[string][]string)
		mutex  = &sync.Mutex{}
	)

	limiter := make(chan string, 10) // Limits simultaneous requests
	wg := sync.WaitGroup{}           // Needed to not prematurely exit before all requests have been finished

	for i, domain := range input {
		limiter <- domain

		wg.Add(1)

		go func(i int, domain string) {
			defer wg.Done()

			defer func() { <-limiter }()

			resp, err := http.Get(domain)

			mutex.Lock()

			if err == nil {
				for key, elem := range resp.Header {
					_, exists := result[key]
					if !exists {
						result[key] = golazy.RemoveDuplicateValues(elem)
					} else {
						var update = result[key]
						update = append(update, elem...)
						result[key] = golazy.RemoveDuplicateValues(update)
					}
				}

				resp.Body.Close()
			}
			mutex.Unlock()
		}(i, domain)
	}

	wg.Wait()

	for key, elem := range result {
		fmt.Printf("%s : %s\n", key, golazy.RemoveDuplicateValues(elem))
		fmt.Println()
	}
}
