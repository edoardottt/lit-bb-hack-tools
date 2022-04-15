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
		for _, elem := range input {
			resultString := ReplaceParameters(elem, *payloadPtr)
			if resultString != "" {
				result = append(result, resultString)
			}
		}
	}
	if *payloadFilePtr != "" {
		payloads := ReadFileLineByLine(*payloadFilePtr)
		for _, payload := range RemoveDuplicateValues(payloads) {
			if strings.Trim(payload, " ") != "" {
				for _, elem := range input {
					resultString := ReplaceParameters(elem, payload)
					if resultString != "" {
						result = append(result, resultString)
					}
				}
			}
		}
	}
	for _, elem := range RemoveDuplicateValues(result) {
		fmt.Println(elem)
	}

	/*
		if *payloadPtr != "" && *payloadFilePtr != "" {
			fmt.Println("You can't specify both -p and -pL.")
			os.Exit(0)
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
		} else if *payloadFilePtr != "" {
			input := ScanTargets()
		} else {
			fmt.Println("Payload(s) required.")
			os.Exit(0)
		}
	*/
}

//help shows the usage
func help() {
	var usage = `Take as input on stdin a list of urls and a payload and print on stdout all the unique urls with ready to use payloads.
	$> cat urls | rapwp -p "<svg onload=alert(1)>"
	$> cat urls | rapwp -pL payloads.txt`
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
