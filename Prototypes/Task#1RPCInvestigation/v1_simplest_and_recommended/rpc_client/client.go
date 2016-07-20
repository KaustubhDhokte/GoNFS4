/*
 * A simple TCP RPC client
*/

package main

import (
    "net/rpc"
    "log"
    "rpc_def"
    "fmt"
)

func main(){

	args := rpc_def.Args{17, 8}
	serverAddress := "127.0.0.1"
	var answer int
	client, err := rpc.Dial("tcp", serverAddress + ":1234")
	if err != nil {
        	log.Fatal("dialing:", err)
	}
        fmt.Printf("\nCalling Remote method Add with")
        fmt.Printf("\nOperand 1 = 17")
        fmt.Printf("\nOperand 2 = 8")
	err = client.Call("Arithmatic.Add", args, &answer)
	if err != nil {
    		log.Fatal("Arith error:", err)
	}
        fmt.Printf("\nAnswer returned by server = %d\n", answer)
}
