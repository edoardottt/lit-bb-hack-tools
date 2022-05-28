package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

/*
TODO:

- Daily api calls check
*/

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"

//main >
func main() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
	}
	helpPtr := flag.Bool("h", false, "Show usage.")
	keyPtr := flag.String("k", "", "API key (if not provided read it from config file).")
	outputPtr := flag.String("o", "", "Output file.")
	flag.Parse()
	if *helpPtr {
		help()
	}
	var apikey string
	if *keyPtr != "" {
		apikey = *keyPtr
	} else {
		apikey = ReadApiKey()
		if apikey == "" {
			fmt.Println(Red + "[ ERR! ] API key cannot be blank." + Reset)
			os.Exit(1)
		}
	}
	if *outputPtr != "" {
		f, err := os.OpenFile(*outputPtr, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Can't create output file.")
			os.Exit(1)
		}
		err = f.Truncate(0)
		if err != nil {
			fmt.Println("Can't create output file.")
			os.Exit(1)
		}
		f.Close()
	}
	input := ScanTargets()
	for _, elem := range RemoveDuplicateValues(input) {
		resp, _, err := KnoxssApi(elem, apikey)
		if err != nil {
			if *outputPtr != "" {
				AppendOutputToTxt("[ ERR! ] "+elem, *outputPtr)
				AppendOutputToTxt(err.Error(), *outputPtr)
			}
			fmt.Println(Red + "[ ERR! ] " + Reset + elem)
			fmt.Println(err.Error())
			continue
		}
		result, err := ReadResult(resp)
		if err != nil {
			if *outputPtr != "" {
				AppendOutputToTxt("[ ERR! ] "+elem, *outputPtr)
			}
			fmt.Println(Red + "[ ERR! ] " + Reset + elem)
			fmt.Println(err.Error())
		} else if result.Xss == "true" { // XSS
			if *outputPtr != "" {
				AppendOutputToTxt("[ XSS! ] "+elem, *outputPtr)
			}
			fmt.Println(Green + "[ XSS! ] " + Reset + result.PoC)
		} else if result.Xss == "none" && result.Error != "" {
			if *outputPtr != "" {
				AppendOutputToTxt("[ ERR! ] "+elem, *outputPtr)
				AppendOutputToTxt(result.Error, *outputPtr)
			}
			fmt.Println(Red + "[ ERR! ] " + Reset + elem)
			fmt.Println(result.Error)
		} else {
			if *outputPtr != "" {
				AppendOutputToTxt("[ SAFE ] "+elem, *outputPtr)
			}
			fmt.Println(Yellow + "[ SAFE ] " + Reset + result.Target)
		}
	}
}

//help shows the usage
func help() {
	var usage = `Take as input on stdin a list of urls and print on stdout the results from Knoxss.me API.
	$> cat urls | knoxssme
	$> cat urls | knoxssme -h exampleapikeywbfkwfiuwlahlflvug
	$> cat urls | knoxssme -o output.txt`
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
		if strings.TrimSpace(sc.Text()) != "" {
			domain := strings.ToLower(sc.Text())
			result = append(result, domain)
		}
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

//FilterAnd replaces all occurrences of & with %26
func FilterAnd(input string) string {
	return strings.ReplaceAll(input, "&", "%26")
}

//KnoxssApi performs a POST request to knoxss api
func KnoxssApi(url string, apikey string) (string, int, error) {
	postBody := FilterAnd("target=" + url)
	responseBody := bytes.NewBuffer([]byte(postBody))
	client := &http.Client{
		Timeout: time.Second * 1000,
	}
	var target = "https://knoxss.me/api/v3"
	req, err := http.NewRequest("POST", target, responseBody)
	if err != nil {
		return "", 0, err
	}
	req.Header.Set("X-API-KEY", apikey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	sb := string(body)
	return sb, resp.StatusCode, nil
}

type Result struct {
	Xss         string `json:"XSS"`
	PoC         string `json:"PoC"`
	Target      string `json:"Target"`
	PostData    string `json:"POST Data"`
	Error       string `json:"Error"`
	ApiCall     string `json:"API Call"`
	TimeElapsed string `json:"Time Elapsed"`
	Timestamp   string `json:"Timestamp"`
}

//ReadResult
func ReadResult(input string) (Result, error) {
	result := Result{}
	var err error
	if strings.Contains(input, "{") && strings.Contains(input, "XSS") { // if json response
		err = json.Unmarshal([]byte(input), &result)
	} else if strings.Contains(input, "Error Code: <b>HTTP 504</b>") { //Sucuri page
		err = errors.New("knoxss.me replied with HTTP 504: Backend or gateway connection timeout")
	} else if strings.Contains(input, "Incorrect API key") {
		fmt.Println(Red + "[ ERROR ] " + Reset + "Incorrect API key.")
		os.Exit(1)
	}
	return result, err
}

//ReadApiKey
func ReadApiKey() string {
	filename := ""
	if runtime.GOOS == "windows" {
		// Don't use colors
		fmt.Println("[ ERROR ] Use -k option to insert Api key.")
		os.Exit(1)
	} else { // linux
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(Red + "[ ERROR ] " + Reset + "Cannot read API Key from config file.")
			os.Exit(1)
		}
		filename = home + "/.config/knoxss/knoxss.key"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(Red + "[ ERROR ] " + Reset + "failed to open " + filename)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var key string
	if scanner.Scan() {
		key = scanner.Text()
	}
	file.Close()
	return key
}

//AppendOutputToTxt >
func AppendOutputToTxt(output string, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	if _, err := file.WriteString(output + "\n"); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
