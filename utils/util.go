package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var MoveCardinal [5]int = [5]int{0, 1, 0, -1, 0}
var MoveDiagonal [5]int = [5]int{-1, 1, 1, -1, -1}

func GetDay(reader *bufio.Reader) string {
	fmt.Print("Enter day: ")
	day, _ := reader.ReadString('\n')
	day = strings.TrimSpace(day)
	
	if day == "" { return strconv.Itoa(time.Now().Day()) }
	if _, dayErr := strconv.Atoi(day); dayErr != nil { panic(dayErr) }
	return day
}

func HandleErr(e error) {
	if e != nil { panic(e) }
}

func getSession() string {
	data, err := os.ReadFile("session.txt")
	HandleErr(err)
	return strings.TrimSpace(string(data))
}

func StrToI(s string) int {
	v, _ := strconv.Atoi(s);
	return v
}

func getInputData(day string) {
	url := "https://adventofcode.com/2024/day/" + day + "/input"
	sessionId := getSession()
	
	req, reqErr := http.NewRequest("GET", url, nil)
	HandleErr(reqErr)
	req.AddCookie(&http.Cookie{Name: "session", Value: sessionId})
	var client http.Client
	resp, respErr := client.Do(req)
	HandleErr(respErr)
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Can not retrieve data from: ", url)
		return
	}
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil { panic(readErr) }
	inputPath := filepath.Join(".", "inputs", "input" + day + ".txt")
	os.WriteFile(inputPath, body, 0666)
}

func SetUp(reader *bufio.Reader) {
	day := GetDay(reader)

	getInputData(day)
	templatePath := filepath.Join(".", "utils", "template.txt")
	template, err := os.ReadFile(templatePath)
	HandleErr(err)
	for i := 0; i < len(template); i++ {
		if template[i] == '#' {
			template[i] = day[0]
			if len(day) > 1 {
				i++
				template[i] = day[1]
			}
		}
	}
	solutionPath := filepath.Join(".", "solutions", "day" + day + ".go")
	os.WriteFile(solutionPath, template, 0666)
}

func IdxInValid(r, c, n int) bool { return r < 0 || c < 0 || r >= n || c >= n }

func GetCoord(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func FuncTimer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}