package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	host = ""
	port = 4352

	heading float64
	lat     float64
	lon     float64
)

func init() {
	flag.IntVar(&port, "p", port, "Port to connect to NMEA device")
}

func connect(host, port) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return conn
}

func parseSentences(conn) {
	for {
		sentence, err := bufio.NewReaderSize(conn, 80).ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
	}
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Printf("%s [-p port] <NMEA host>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	host = flag.Arg(0)
	fmt.Println("Connecting to %s:%d...", host, port)
	conn := connect(host, port)
	fmt.Println("Success")
	go parseSentences(conn)
}
