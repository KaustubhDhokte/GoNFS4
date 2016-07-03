package main

import (
    "log"
    "fmt"
    "rpc_def"
    "net/rpc/jsonrpc"
    "net"
)


func main(){

        //Connect to the server
  	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	args := rpc_def.Args{7, 8}
	var reply int
	//Create new client object for jsonrpc
	c := jsonrpc.NewClient(conn)
	//Call the remote procedure Add exposed on type Arithmatic
	err = c.Call("Arithmatic.Add", args, &reply)
	if err != nil {
	    log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d+%d=%d\n", args.Operand1, args.Operand2, reply)

}
