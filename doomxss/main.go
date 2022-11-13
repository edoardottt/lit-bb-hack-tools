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
		fmt.Println("[ " + elem.Sink + " ] " + elem.URL)
	}
}

// help shows the usage.
func help() {
	var usage = `Take as input on stdin a list of html/js file urls and print on stdout all the possible DOM XSS sinks found.
	$> cat urls | doomxss`

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

type Result struct {
	Sink string
	URL  string
}

// RetrieveContents.
func RetrieveContents(input []string) []Result {
	var (
		result = []Result{}
		mutex  = &sync.Mutex{}
	)

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
					results := CheckSinks(sb, domain)
					result = append(result, results...)
				}

				resp.Body.Close()
			}
			mutex.Unlock()
		}(i, domain)
	}

	wg.Wait()

	return result
}

// CheckSinks returns a slice containing all
// the probable sinks in the body.
func CheckSinks(body string, url string) []Result {
	var result []Result

	toCheck := strings.ToLower(body)
	toCheck2 := strings.ReplaceAll(toCheck, " ", "")

	for _, sink := range sinks {
		if strings.Contains(toCheck2, sink) {
			res := Result{Sink: sink, URL: url}
			result = append(result, res)
		}
	}

	return result
}

var sinks = []string{
	"document.url=",
	"document.documenturi=",
	"document.urlencoded=",
	"document.baseuri=",
	"location=",
	"location.href=",
	"location.search=",
	"location.hash=",
	"location.pathname=",
	"document.cookie=",
	"document.referrer=",
	"window.name=",
	"history.pushstate(",
	"history.replacestate(",
	"localstorage.setitem(",
	"localstorage.getitem(",
	"sessionstorage=",
	"document.write(",
	"document.writeIn(",
	"innerHTML=",
	"outerHTML=",
	"eval(",
	"setTimeout(",
	"setInterval(",
	"{{__html", // REACT
}
