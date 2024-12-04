package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getDay() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter day: ")
	day, _ := reader.ReadString('\n')
	day = strings.TrimSpace(day)
	if _, dayErr := strconv.Atoi(day); dayErr != nil { panic(dayErr) }
	return day
}

func handleErr(e error) {
	if e != nil { panic(e) }
}

func getSession() string {
	data, err := os.ReadFile("session.txt")
	handleErr(err)
	return strings.TrimSpace(string(data))
}

func main() {
	day := getDay()
	url := "https://adventofcode.com/2024/day/" + day + "/input"
	session := getSession()

	req, reqErr := http.NewRequest("GET", url, nil)
	handleErr(reqErr)
	req.AddCookie(&http.Cookie{Name: "session", Value: session})
	var client http.Client
	resp, respErr := client.Do(req)
	handleErr(respErr)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Can not retrieve data from: ", url)
		return
	}
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil { panic(readErr) }
	dataFile := "day" + day + "/input.txt"
	os.WriteFile(dataFile, body, 0666)
}