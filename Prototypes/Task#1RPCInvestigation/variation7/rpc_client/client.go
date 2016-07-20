/*
 * A simple json RPC client
*/

package main

import (
    "log"
    "fmt"
    "rpc_def"
    "net/rpc/jsonrpc"
)


func main(){

	args := rpc_def.Args{17, 8}
	serverAddress := "127.0.0.1"
	var answer int
	//Connect to remote server        
	conn, err := jsonrpc.Dial("tcp", serverAddress + ":1234")

	if err != nil {
        	log.Fatal("dialing:", err)
	}
	//Call remote procedure
        fmt.Printf("\nCalling Remote method Add with")
        fmt.Printf("\nOperand 1 = 17")
        fmt.Printf("\nOperand 2 = 8")
	err = conn.Call("Arithmatic.Add", args, &answer)
	if err != nil {
    		log.Fatal("arith error:", err)
	}
	fmt.Printf("\nAnswer returned by the server =  %d\n", answer)
}
