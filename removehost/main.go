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
	helpPtr := flag.Bool("h", false, "Show usage.")
	flag.Parse()
	if *helpPtr {
		help()
	}
	input := ScanTargets()
	var result []string
	for _, elem := range input {
		query := GetQuery(elem)
		if query != "" {
			result = append(result, query)
		}
	}
	result = RemoveDuplicateValues(result)
	for _, elem := range result {
		fmt.Println(elem)
	}
}

//help shows the usage
func help() {
	var usage = `Take as input on stdin a list of urls and print on stdout all the unique queries without protocol and host.
	$> cat urls | removehost`
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

//GetQuery >
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
