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
        goreply := client.Go("Arithmatic.Add", args, &answer, nil)
        reply := <-goreply.Done
        if reply!=nil{
        }
        fmt.Printf("Arith = %d\n", answer)
}
