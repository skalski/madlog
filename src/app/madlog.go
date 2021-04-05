package main

import (
	"flag"
	"fmt"
	"log"
	"madlog/logoutput"
	"os/exec"
	"strings"
	"time"
)

var loglevel *int
var version string = "0.2"

func main() {
	WelcomeMsg()
	SetLoglevel()
	containerIds := FetchContainerIds()
	for i, s := range containerIds {
		if s != "" {
			fmt.Printf("Found container %d: %s \n", i, s)
		}
	}
	fmt.Println("Start log output ....")
	StartLogFetcher(containerIds)
}

func SetLoglevel() {
	loglevel = flag.Int("level", 0, "set loglevel\n0 : ALL\n1 : ERROR/EXCEPTION")
	flag.Parse()
	fmt.Printf("set loglevel to: %d \n", *loglevel)
}

func StartLogFetcher(containerIds []string) {
	for {
		for i, s := range containerIds {
			if s != "" {
				out, err := exec.Command("docker", "logs", "--since=2s", s).Output()
				if err != nil {
					log.Fatal(err)
				}
				logOutput := BytesToString(out)
				if logOutput != "" {
					logoutput.ParseMessageByLevel(i, logOutput, loglevel)
				}
			}
		}
		time.Sleep(2 * time.Second)
		containerIds = FetchContainerIds()
	}
}

func BytesToString(data []byte) string {
	return string(data[:])
}

func WelcomeMsg() {
	fmt.Println("     --- Madlog ---")
	fmt.Printf("      version:%s\n", version)
	fmt.Print(" written by Swen Kalski\n\n\n")
}

func FetchContainerIds() []string {
	var returnIds []string
	out, err := exec.Command("docker", "ps", "-q").Output()
	if err != nil {
		log.Fatal(err)
	}
	containerIds := strings.Split(strings.ReplaceAll(BytesToString(out), "\r\n", "\n"), "\n")
	for _, s := range containerIds {
		if s != "" {
			returnIds = append(returnIds, s)
		}
	}
	return returnIds
}
