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
	payloadPtr := flag.String("p", "", "Input payload.")
	flag.Parse()
	if *helpPtr {
		help()
	}
	if *payloadPtr != "" {
		input := ScanTargets()
		var result []string
		for _, elem := range input {
			resultString := ReplaceParameters(elem, *payloadPtr)
			if resultString != "" {
				result = append(result, resultString)
			}

		}
		for _, elem := range RemoveDuplicateValues(result) {
			fmt.Println(elem)
		}
	} else {
		fmt.Println("Payload required.")
		os.Exit(0)
	}
}

//help shows the usage
func help() {
	var usage = `Take as input on stdin a list of urls and a payload and print on stdout all the unique urls with ready to use payloads.
	$> cat urls | rapwp -p "<svg onload=alert(1)>"`
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

//ReplaceParameters >
func ReplaceParameters(input string, payload string) string {
	u, err := url.Parse(input)
	if err != nil {
		return ""
	}
	decodedValue, err := url.QueryUnescape(u.RawQuery)
	if err != nil {
		return ""
	}
	var queryResult = ""
	couples := strings.Split(decodedValue, "&")
	for _, pair := range couples {
		values := strings.Split(pair, "=")
		queryResult += values[0] + "=" + url.QueryEscape(payload) + "&"
	}
	return u.Scheme + "://" + u.Host + u.Path + "?" + queryResult[:len(queryResult)-1]
}
