package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

//main
func main() {
	RetrieveHeaders(ScanTargets())
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

//RetrieveHeaders
func RetrieveHeaders(input []string) {
	result := make(map[string][]string)
	var mutex = &sync.Mutex{}

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
						result[key] = removeDuplicateValues(elem)
					} else {
						var update = result[key]
						update = append(update, elem...)
						result[key] = removeDuplicateValues(update)
					}
				}
				resp.Body.Close()
			}
			mutex.Unlock()
		}(i, domain)
	}
	wg.Wait()
	for key, elem := range result {
		fmt.Printf("%s : %s\n", key, removeDuplicateValues(elem))
		fmt.Println()
	}
}

//removeDuplicateValues
func removeDuplicateValues(strSlice []string) []string {
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
