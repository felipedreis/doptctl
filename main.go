package main

import (
	"fmt"
	//"context"
	"flag"
	//"io"
	"log"
	//"time"

	"google.golang.org/grpc"
)
var (
	serverAddr = flag.String("addr", "localhost:8080", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	var opts []grpc.DialOption


	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

}
