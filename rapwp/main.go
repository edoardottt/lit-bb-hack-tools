package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")
	payloadPtr := flag.String("p", "", "Input payload.")
	payloadFilePtr := flag.String("pL", "", "Input payload file.")
	oneByOnePtr := flag.Bool("obo", false, "Replace parameters one by one.")
	flag.Parse()
	if *helpPtr {
		help()
	}
	if *payloadPtr != "" && *payloadFilePtr != "" {
		fmt.Println("You can't specify both -p and -pL.")
		os.Exit(0)
	}
	if *payloadPtr == "" && *payloadFilePtr == "" {
		fmt.Println("Payload(s) required.")
		os.Exit(0)
	}
	input := ScanTargets()
	var result []string
	if *payloadPtr != "" {
		if !*oneByOnePtr {
			for _, elem := range input {
				resultString := ReplaceParameters(elem, *payloadPtr)
				if resultString != "" {
					result = append(result, resultString)
				}
			}
		} else {
			for _, elem := range input {
				resultSlice := ReplaceParametersOneByOne(elem, *payloadPtr)
				if len(resultSlice) != 0 {
					result = append(result, resultSlice...)
				}
			}

		}
	}
	if *payloadFilePtr != "" {
		payloads := ReadFileLineByLine(*payloadFilePtr)
		for _, payload := range RemoveDuplicateValues(payloads) {
			if strings.Trim(payload, " ") != "" {
				if !*oneByOnePtr {
					for _, elem := range input {
						resultString := ReplaceParameters(elem, payload)
						if resultString != "" {
							result = append(result, resultString)
						}
					}
				} else {
					for _, elem := range input {
						resultSlice := ReplaceParametersOneByOne(elem, payload)
						if len(resultSlice) != 0 {
							result = append(result, resultSlice...)
						}
					}
				}
			}
		}
	}
	for _, elem := range RemoveDuplicateValues(result) {
		fmt.Println(elem)
	}
}

//help shows the usage
func help() {
	var usage = `Take as input on stdin a list of urls and a payload and print on stdout all the unique urls with ready to use payloads.
	$> cat urls | rapwp -p "<svg onload=alert(1)>"
	$> cat urls | rapwp -pL payloads.txt
	$> cat urls | rapwp -pL payloads.txt -obo`
	fmt.Println()
	fmt.Println(usage)
	fmt.Println()
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

//ReplaceParametersOneByOne >
func ReplaceParametersOneByOne(input string, payload string) []string {
	u, err := url.Parse(input)
	if err != nil {
		return []string{}
	}
	decodedValue, err := url.QueryUnescape(u.RawQuery)
	if err != nil {
		return []string{}
	}
	var queryResult []string
	couples := strings.Split(decodedValue, "&")
	for _, pair1 := range couples {
		var query = ""
		for _, pair := range couples {
			if pair1 == pair {
				values := strings.Split(pair, "=")
				query += values[0] + "=" + url.QueryEscape(payload) + "&"
			} else {
				values := strings.Split(pair, "=")
				query += values[0] + "=" + values[1] + "&"
			}
		}
		queryResult = append(queryResult, u.Scheme+"://"+u.Host+u.Path+"?"+query[:len(query)-1])
	}

	return queryResult
}

//ReadFileLineByLine reads all the lines from input file and returns
//them as a slice of strings
func ReadFileLineByLine(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("failed to open %s", inputFile)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	text = RemoveDuplicateValues(text)
	return text
}
