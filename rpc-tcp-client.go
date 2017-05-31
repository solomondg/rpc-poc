package main

import (
	"fmt"
	"github.com/solomondg/rpc-poc/remote"
	"log"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var args remote.PiCalcParams
	args.NumIterations = 5000000

	start := time.Now()

	var reply float64
	// e := client.Call("Pi.CalcPi", args, &reply)
	e := client.Call("Pi.CalcPi", args, &reply)

	fmt.Println(reply)
	fmt.Printf("Took %d to compute or whatever\n", time.Since(start).Seconds())
	fmt.Println(e)
}
