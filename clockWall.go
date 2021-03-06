package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func handleConn2(conn net.Conn, place string) {
	for {
		hour := make([]byte, 11)
		_, err := conn.Read(hour)
		if err != nil {
			conn.Close()
			log.Print(err)
		} else {
			fmt.Printf("%s : %s", place, hour)
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {

	if len(os.Args[1:]) < 1 {
		fmt.Println("Uso equivocado.")
		os.Exit(1)
	}

	for i := 0; i < len(os.Args[1:]); i++ {
		//
		//os.Args[1:] is a slice of strings.
		clockData := strings.Split(os.Args[i+1], "=")
		if len(clockData) != 2 {
			fmt.Println("Revisa los argumentos.")
		}
		conn, err := net.Dial("tcp", clockData[1])
		if err != nil {
			log.Fatal(err)
		}
		go handleConn2(conn, clockData[0])

	}
	for {
	} 

}