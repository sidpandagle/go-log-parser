package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type LogDetails struct {
	TimeStamp string
	Level     string
	Code      string
	Method    string
	URL       string
	IP        string
}

func main() {
	file, err := os.Open("weblog.log")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var logDetailsList []LogDetails

	for scanner.Scan() {
		line := scanner.Text()
		logDetailsList = append(logDetailsList, getLogDetails(line))
	}

	jsonResult, err := json.Marshal(logDetailsList)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonResult))

	// Error whle scanning
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
}

func getLogDetails(line string) LogDetails {
	var logDetails LogDetails
	words := strings.Fields(line)
	logDetails.TimeStamp = words[0] + " " + words[1]
	logDetails.Level = words[2]
	logDetails.Code = words[3]
	logDetails.Method = words[4]
	logDetails.URL = words[5]
	logDetails.IP = words[7]
	return logDetails
}
