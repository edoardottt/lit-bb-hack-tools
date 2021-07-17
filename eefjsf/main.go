package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

func main() {
	input := ScanTargets()
	results := RetrieveContents(removeDuplicateValues(input))
	for _, elem := range results {
		fmt.Println(elem[1 : len(elem)-1])
	}
}

//ScanTargets return the array of elements
//taken as input on stdin.
func ScanTargets() []string {

	var result []string
	// accept domains on stdin
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		result = append(result, domain)
	}
	return result
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

//getHeaders
func RetrieveContents(input []string) []string {
	var result []string
	var mutex = &sync.Mutex{}
	r, _ := regexp.Compile(`\"\/[a-zA-Z0-9_\/?=&]*\"`)

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
				body, err := ioutil.ReadAll(resp.Body)
				if err == nil && len(body) != 0 {
					//Convert the body to type string
					sb := string(body)
					results := r.FindAllString(sb, -1)
					result = append(result, removeDuplicateValues(results)...)
				}

				resp.Body.Close()
			}
			mutex.Unlock()
		}(i, domain)
	}
	wg.Wait()
	return removeDuplicateValues(result)
}
