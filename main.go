package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jbaikge/nmea-relay/decoder"
	"log"
	"net"
	"net/http"
	"os"
)

type Location struct {
	Heading   float64
	Latitude  float64
	Longitude float64
}

var (
	current = &Location{}
	host    = ""
	listen  = ":80"
	port    = 4352
)

func init() {
	flag.IntVar(&port, "p", port, "Port to connect to NMEA device")
	flag.StringVar(&listen, "l", listen, "HTTP port to listen")
}

func connect(host string, port int) net.Conn {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	return conn
}

func httpLocation(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.RemoteAddr, r.RequestURI)
	if err := json.NewEncoder(w).Encode(current); err != nil {
		fmt.Fprint(w, err)
	}
}

func parseSentences(conn net.Conn) {
	r := bufio.NewReaderSize(conn, 128)
	for {
		sentence, err := r.ReadBytes('\n')
		if err != nil {
			log.Print("NMEA provider closed connection")
			log.Print(err)
			os.Exit(3)
		}

		msg, err := decoder.Decode(sentence[:len(sentence)-1])
		if err != nil {
			log.Printf("%s - %s", err, sentence)
		}

		switch msg.Type {
		case "GLL":
			log.Print(string(sentence))
			current.Latitude = msg.Latitude
			current.Longitude = msg.Longitude
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
	log.Printf("Connecting to %s:%d...", host, port)
	conn := connect(host, port)
	defer conn.Close()
	log.Printf("Connected to %s:%d", host, port)

	go parseSentences(conn)

	http.HandleFunc("/location", httpLocation)
	log.Fatal(http.ListenAndServe(listen, nil))
}
