package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	input := ScanTargets()
	var result []string
	if !ScanFlag() {
		for _, elem := range input {
			result = append(result, RemoveProtocol(elem))
		}
	} else {
		for _, elem := range input {
			sub := RemovePort(RemoveProtocol(GetOnlySubs(elem)))
			if sub != "" {
				result = append(result, sub)
			}
		}
	}
	result = RemoveDuplicateValues(result)
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

//RemoveDuplicateValues >
func RemoveDuplicateValues(strSlice []string) []string {
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

//RemoveProtocol >
func RemoveProtocol(input string) string {
	res := strings.Index(input, "://")
	if res >= 0 {
		return input[res+3:]
	} else {
		return input
	}
}

//GetOnlySubs >
func GetOnlySubs(input string) string {
	u, err := url.Parse(input)
	if err != nil {
		return ""
	}
	return u.Host
}

//ScanFlag >
func ScanFlag() bool {
	subsPtr := flag.Bool("subs", false, "Return only subdomains without protocols.")
	flag.Parse()
	return *subsPtr
}

//RemovePort >
func RemovePort(input string) string {
	res := strings.Index(input, ":")
	if res >= 0 {
		return input[:res-1]
	}
	return input
}
