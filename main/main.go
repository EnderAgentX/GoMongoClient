package main

import (
	//"log"
	//"net"

	"github.com/EnderAgentX/app"
)

func main() {
	// conn, err := net.Dial("tcp", "192.168.0.106:8080")
	// if err != nil {
	// 	log.Fatalf("Connection to the server failed: %v", err)
	// }
	// defer conn.Close()

	app.App()

}
