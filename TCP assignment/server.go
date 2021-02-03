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
	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "Exit" {
			fmt.Println("Exiting TCP server!")
			return
		}

		t1 := time.Now()
		myTime1 := t1.Format(time.RFC3339) + "\n"
		fmt.Print("client "+myTime1+" ->: ", string(netData))

		fmt.Print(" server " + myTime1 + " >>")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		c.Write([]byte(text))

	}
}
