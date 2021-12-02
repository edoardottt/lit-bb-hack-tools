package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

//main
func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")
	flag.Parse()
	if *helpPtr {
		help()
	}

}

//help shows the usage
func help() {
	var usage = `Take as input on stdin a list of urls and print on stdout all the urls with unique parameters.
	$> cat urls | mauu`
	fmt.Println(usage)
	os.Exit(0)
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
	return RemoveDuplicateValues(result)
}

//RemoveDuplicateValues
func RemoveDuplicateValues(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

//GetSchemeHostPath >
func GetSchemeHostPath(input string) string {
	u, err := url.Parse(input)
	if err != nil {
		return ""
	}
	return u.Scheme + u.Host + u.Path
}

//GetRawQuery >
func GetRawQuery(input string) string {
	u, err := url.Parse(input)
	if err != nil {
		return ""
	}
	return u.RawQuery
}

//Mauu >
func Mauu(input []string) map[string]string {
	result := make(map[string]string)
	for _, elem := range input {
		// that url does not exist
		if _, exists := result[GetSchemeHostPath(elem)]; !exists {
			result[GetSchemeHostPath(elem)] = GetRawQuery(elem)
		} else {
			// otherwise check parameters

		}
	}

	return result
}
