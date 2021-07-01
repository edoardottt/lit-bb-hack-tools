package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	input := ScanTargets()
	result := GetRobots(input)
	result = removeDuplicateValues(result)
	for _, elem := range result {
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
		domain := strings.ToLower(sc.Text())
		result = append(result, domain)
	}
	return removeDuplicateValues(result)
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

//GetRobots
func GetRobots(input []string) []string {
	var result []string
	var mutex = &sync.Mutex{}

	limiter := make(chan string, 10) // Limits simultaneous requests
	wg := sync.WaitGroup{}           // Needed to not prematurely exit before all requests have been finished

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
						if len(word) > 1 {
							result = append(result, word[1])
						}
					}
				}
			}
			mutex.Unlock()
		}(elem)
	}
	wg.Wait()
	return removeDuplicateValues(result)
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

//RemoveProtocol
func RemoveProtocol(input string) string {
	res := strings.Index(input, "://")
	if res >= 0 {
		return input[res+3:]
	} else {
		return input
	}
}
