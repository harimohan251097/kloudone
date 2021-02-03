package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)

		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"

		fmt.Print(" client " + myTime + " >>")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print(" server " + myTime + " ->: " + message)
		if strings.TrimSpace(string(text)) == "Exit" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
