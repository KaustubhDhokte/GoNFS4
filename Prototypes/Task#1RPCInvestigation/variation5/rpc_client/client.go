/*
 * A simple HTTP RPC client
*/

package main

import (
    "net/rpc"
    "log"
    "fmt"
    "rpc_def"
)

func main(){

	args := rpc_def.Args{17, 8}
	serverAddress := "127.0.0.1"
	var answer int
	//Connect to RPC server listening on HTTP        
	client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if err != nil {
        	log.Fatal("dialing:", err)
	}
	//Call remote procedure
        fmt.Printf("\nCalling Remote method Add with")
        fmt.Printf("\nOperand 1 = 17")
        fmt.Printf("\nOperand 2 = 8")
	err = client.Call("Arithmatic.Add", args, &answer)
	if err != nil {
    		log.Fatal("arith error:", err)
	}
        fmt.Printf("\nAnswer returned by the server = %d", answer)
}
