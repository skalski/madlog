package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	WelcomeMsg()

	out, err := exec.Command("docker", "ps", "-q").Output()
	if err != nil {
		log.Fatal(err)
	}
	container_ids := strings.Split(strings.ReplaceAll(BytesToString(out), "\r\n", "\n"), "\n")
	for i, s := range container_ids {
		if s != "" {
			fmt.Printf("Found container %d: %s \n", i, s)
		}
	}
	fmt.Println("Start log output ....")
	StartLogFetcher(container_ids)
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
					fmt.Printf("LOG %d:%s \n", i, s)
					fmt.Print(log)
				}
			}
		}
		time.Sleep(2 * time.Second)
	}
}

func BytesToString(data []byte) string {
	return string(data[:])
}

func WelcomeMsg() {
	fmt.Println("     --- Madlog ---")
	fmt.Print(" written by Swen Kalski\n\n\n")
}
