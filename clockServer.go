// Clock2 is a concurrent TCP server that periodically writes the time.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func handleConn(c net.Conn, tz string) {
	defer c.Close()
	for {
		_, tErr := time.LoadLocation(tz)
		if tErr != nil {
			fmt.Println("Error loading the %s timezone: ", tz)
			log.Print(tErr)
			break
		}
		
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Print(err)
			return 
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {

	if len(os.Args[1:]) < 2 {
		fmt.Println("Uso equivocado.")
		os.Exit(1)
	}
	fmt.Println(os.Args[2])
	listener, err := net.Listen("tcp", "localhost:"+os.Args[2])

	if err != nil {
		log.Fatal(err)
	}

	tz := os.Getenv("TZ")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) 
			continue
		}
		go handleConn(conn, tz) 
	}
}