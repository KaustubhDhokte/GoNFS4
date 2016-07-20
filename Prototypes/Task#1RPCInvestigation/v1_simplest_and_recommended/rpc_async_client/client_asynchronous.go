/*
 * A TCP RPC client invoking remote procedure asynchronously
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
        fmt.Printf("\n\nCalling Remote method Add asynchronously with")
        fmt.Printf("\nOperand 1 = 17")
        fmt.Printf("\nOperand 2 = 8")
        goreply := client.Go("Arithmatic.Add", args, &answer, nil)
        reply := <-goreply.Done
        if reply!=nil{
        }
        fmt.Printf("\nAnswer returned by the server = %d\n", answer)
}
