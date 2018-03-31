package main

import (
	"flag"
	"fmt"
	"github.com/tevino/tcp-shaker"
	"log"
	"strings"
	"time"
	"github.com/google/cloudprober/probes/external/serverutils"
	"github.com/golang/protobuf/proto"
	"strconv"
)

var address = flag.String("address", "", "Address of target host")
var port = flag.Int("port", -1, "TCP port of target.")
var server = flag.Bool("server", false, "Start probe in server mode.")

func probe(address *string, port *int) (string, error) {
	var payload []string

	c := tcp.NewChecker(true)
	if err := c.InitChecker(); err != nil {
		return "", err
	}

	timeout := time.Second * 1
	startTime := time.Now()
	err := c.CheckAddr(fmt.Sprintf("%s:%d", *address, *port), timeout)
	if err != nil {
		return "", err
	}
	payload = append(payload, fmt.Sprintf("latency_ms %f", float64(time.Since(startTime).Nanoseconds()/1e6)))

	return strings.Join(payload, "\n"), nil
}

func main() {

	flag.Parse()

	if *address == "" {
		log.Fatal("Error. Host address not set.")
	}
	if *port == -1 {
		log.Fatal("Error. TCP port not set.")
	}

	if *server {
		serverutils.Serve(func(request *serverutils.ProbeRequest, reply *serverutils.ProbeReply) {
			address := ""
			portString := ""
			for _, option := range request.Options {
				if *option.Name == "address" {
					address = *option.Value
				}
				if *option.Name == "port" {
					portString = *option.Value
				}
			}
			port, err := strconv.Atoi(portString)
			if err != nil {
				reply.ErrorMessage = proto.String(err.Error())
				return
			}

			payload, err := probe(&address, &port)
			reply.Payload = proto.String(payload)
			if err != nil {
				reply.ErrorMessage = proto.String(err.Error())
			}
		})
	}

	payload, err := probe(address, port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(payload)

}
