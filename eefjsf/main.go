package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
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

	input := ScanTargets()

	results := RetrieveContents(golazy.RemoveDuplicateValues(input))
	for _, elem := range results {
		fmt.Println(elem[1 : len(elem)-1])
	}
}

// help shows the usage.
func help() {
	var usage = `Take as input on stdin a list of js file urls and print on stdout all the unique endpoints found.
	$> cat js-urls | eefjsf`

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

	return result
}

// RetrieveContents.
func RetrieveContents(input []string) []string {
	var (
		result = []string{}
		mutex  = &sync.Mutex{}
	)

	r := regexp.MustCompile(`\"\/[a-zA-Z0-9_\/?=&]*\"`)

	limiter := make(chan string, 10) // Limits simultaneous requests.
	wg := sync.WaitGroup{}           // Needed to not prematurely exit before all requests have been finished.

	for i, domain := range input {
		limiter <- domain

		wg.Add(1)

		go func(i int, domain string) {
			defer wg.Done()
			defer func() { <-limiter }()

			resp, err := http.Get(domain)

			mutex.Lock()

			if err == nil {
				body, err := ioutil.ReadAll(resp.Body)
				if err == nil && len(body) != 0 {
					// Convert the body to type string.
					sb := string(body)
					results := r.FindAllString(sb, -1)
					result = append(result, golazy.RemoveDuplicateValues(results)...)
				}

				resp.Body.Close()
			}
			mutex.Unlock()
		}(i, domain)
	}

	wg.Wait()

	return golazy.RemoveDuplicateValues(result)
}
