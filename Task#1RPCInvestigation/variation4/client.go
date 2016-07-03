package main

import (
    "net/rpc"
    "log"
    "fmt"
    "rpc_def"
)

func main(){

	args := rpc_def.Args{17, 8}
	serverAddress := "192.168.0.7"
	var answer int
	//Connect to RPC server
	client, err := rpc.Dial("tcp", serverAddress + ":1234")
	if err != nil {
        	log.Fatal("dialing:", err)
	}
	//Call the remote procedure
	err = client.Call("Arithmatic.Add", args, &answer)
	if err != nil {
    		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d\n", answer)
}
