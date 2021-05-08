package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

//main
func main() {
	RetrieveHeaders(ScanTargets())
}

//ScanInput return the array of elements
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

//getHeaders
func RetrieveHeaders(input []string) {
	var result map[string][]string
	result = make(map[string][]string)
	for _, elem := range input {
		result = GetHeaders(AddHeaders(elem), result)
	}
	for key, elem := range result {
		fmt.Printf("%s : %s\n", key, removeDuplicateValues(elem))
		fmt.Println()
	}
}

//GetRequest performs a GET request and return
//a string (the headers of the response)
func GetHeaders(target string, result map[string][]string) map[string][]string {
	resp, err := http.Get(target)
	if err == nil {
		defer resp.Body.Close()
		for key, elem := range resp.Header {
			_, exists := result[key]
			if !exists {
				result[key] = removeDuplicateValues(elem)
			} else {
				var update = result[key]
				for _, el := range elem {
					update = append(update, el)
				}
				result[key] = removeDuplicateValues(update)
			}
		}
	}
	return result
}

//AddHeaders
func AddHeaders(input string) string {
	if len(input) > 8 {
		if input[:8] != "https://" {
			return "https://" + input
		}
	}
	return input
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
