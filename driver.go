package main

import (
	"os"
	"fmt"
  	"./chordRPC"
)

var (
	nodeAddr string
	peerAddr string
	ftAddr string
)

func main () {
	if len(os.Args) < 3 {
		fmt.Println("=====================================================")
		fmt.Println("USAGE: go run driver.go <nodeAddress> <peerAddress>")
		fmt.Println("EXAMPLE: go run chordRpc.go :14321 :14322")
		fmt.Println("=====================================================")
		fmt.Println("NOTE: If first node in system, use same address")
		fmt.Println("EXAMPLE: go run chordRpc.go :14322 :14322")
		fmt.Println("=====================================================")
		os.Exit(-1)
	}
	nodeAddr = os.Args[1];
	peerAddr = os.Args[2];
	ftAddr = os.Args[3];

	chordRPC.Start(nodeAddr, peerAddr, ftAddr);

}