package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var Reset = "\033[0m"
var Red = "\033[31m"

//main
func main() {
	helpPtr := flag.Bool("h", false, "Show usage.")
	flag.Parse()
	if *helpPtr {
		help()
	}
	TestMethods(ScanTargets())
}

//help shows the usage
func help() {
	var usage = `Take as input on stdin a list of urls and print on stdout all the status codes and body sizes for HTTP methods.
	$> cat urls | tahm`
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

//TestMethods >
func TestMethods(input []string) {
	for _, elem := range input {
		fmt.Println("= " + Red + elem + Reset + " =")
		GetStatus, lenSb, err := GetRequest(elem)
		if err == nil {
			fmt.Println(" GET " + "Status: " + GetStatus + " | Size: " + strconv.Itoa(lenSb))
		}
		PostStatus, lenSb, err := PostRequest(elem)
		if err == nil {
			fmt.Println(" POST " + "Status: " + PostStatus + " | Size: " + strconv.Itoa(lenSb))
		}
		PutStatus, lenSb, err := PutRequest(elem)
		if err == nil {
			fmt.Println(" PUT " + "Status: " + PutStatus + " | Size: " + strconv.Itoa(lenSb))
		}
		DeleteStatus, lenSb, err := DeleteRequest(elem)
		if err == nil {
			fmt.Println(" DELETE " + "Status: " + DeleteStatus + " | Size: " + strconv.Itoa(lenSb))
		}
		HeadStatus, lenSb, err := HeadRequest(elem)
		if err == nil {
			fmt.Println(" HEAD " + "Status: " + HeadStatus + " | Size: " + strconv.Itoa(lenSb))
		}
		fmt.Println("---------------------------")
		fmt.Println()
	}
}

//GetRequest performs a GET request
func GetRequest(target string) (string, int, error) {
	resp, err := http.Get(target)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	//Convert the body to type string
	sb := string(body)
	return resp.Status, len(sb), nil
}

//PostRequest performs a POST request
func PostRequest(target string) (string, int, error) {
	postBody, _ := json.Marshal("{data}")
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(target, "application/json", responseBody)
	//Handle Error
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
	return resp.Status, len(sb), nil
}

//HeadRequest performs a HEAD request
func HeadRequest(target string) (string, int, error) {
	resp, err := http.Head(target)
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
	return resp.Status, len(sb), nil
}

//PutRequest performs a PUT request
func PutRequest(target string) (string, int, error) {
	// initialize http client
	client := &http.Client{}

	// marshal User to json
	json, _ := json.Marshal("{data}")

	// set the HTTP method, url, and request body
	req, err := http.NewRequest(http.MethodPut, target, bytes.NewBuffer(json))
	if err != nil {
		return "", 0, err
	}

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	sb := string(body)

	return resp.Status, len(sb), nil
}

//DeleteRequest performs a DELETE request
func DeleteRequest(target string) (string, int, error) {

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("DELETE", target, nil)
	if err != nil {
		return "", 0, err
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	// Read Response Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	sb := string(body)

	return resp.Status, len(sb), nil
}

//RemoveDuplicateValues >
func RemoveDuplicateValues(strSlice []string) []string {
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
