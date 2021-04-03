package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

var loglevel *int
var version string = "0.2"

func main() {
	WelcomeMsg()
	SetLoglevel()
	container_ids := FecthContainerIds()
	for i, s := range container_ids {
		if s != "" {
			fmt.Printf("Found container %d: %s \n", i, s)
		}
	}
	fmt.Println("Start log output ....")
	StartLogFetcher(container_ids)
}

func SetLoglevel() {
	loglevel = flag.Int("level", 0, "set loglevel\n0 : ALL\n1 : ERROR/EXCEPTION")
	flag.Parse()
	fmt.Printf("set loglevel to: %d \n", *loglevel)
}

func StartLogFetcher(container_ids []string) {
	for {
		for i, s := range container_ids {
			if s != "" {
				out, err := exec.Command("docker", "logs", "--since=2s", s).Output()
				if err != nil {
					log.Fatal(err)
				}
				log := BytesToString(out)
				if log != "" {
					ParseMessageByLevel(i, log)
				}
			}
		}
		time.Sleep(2 * time.Second)
		container_ids = FecthContainerIds()
	}
}

func ParseMessageByLevel(i int, log string) {
	if *loglevel == 0 {
		PrintLog(i, log)
	}
	if *loglevel == 1 && HasError(log) {
		PrintLog(i, log)
	}
}

func PrintLog(i int, log string) {
	fmt.Printf("LOG %d:%s \n", i, log)
	fmt.Print(log)
}

func BytesToString(data []byte) string {
	return string(data[:])
}

func HasError(log string) bool {
	return strings.Contains(strings.ToLower(log), "stacktrace") || strings.Contains(strings.ToLower(log), "error")
}

func WelcomeMsg() {
	fmt.Println("     --- Madlog ---")
	fmt.Printf("      version:%s\n", version)
	fmt.Print(" written by Swen Kalski\n\n\n")
}

func FecthContainerIds() []string {
	var returnIds []string
	out, err := exec.Command("docker", "ps", "-q").Output()
	if err != nil {
		log.Fatal(err)
	}
	container_ids := strings.Split(strings.ReplaceAll(BytesToString(out), "\r\n", "\n"), "\n")
	for _, s := range container_ids {
		if s != "" {
			returnIds = append(returnIds, s)
		}
	}
	return returnIds
}
