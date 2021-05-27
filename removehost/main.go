package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	input := ScanTargets()
	var result []string
	for _, elem := range input {
		query := GetQuery(elem)
		if query != "" {
			result = append(result, query)
		}

	}
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
	return result
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

//GetQuery
func GetQuery(input string) string {
	u, err := url.Parse(input)
	if err != nil {
		return ""
	}
	if u.RawQuery != "" && u.Fragment != "" {
		return u.Path + "?" + u.RawQuery + "#" + u.Fragment
	}
	if u.RawQuery != "" {
		return u.Path + "?" + u.RawQuery
	}
	if u.Fragment != "" {
		return u.Path + "#" + u.Fragment
	}
	return u.Path
}
