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
	err = conn.Call("Arithmatic.Add", args, &answer)
	if err != nil {
    		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d\n", answer)
}
