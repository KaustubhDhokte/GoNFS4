/*
 * RPC client that uses json as wire format for marshalling the arguments
*/

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
        fmt.Printf("\nCalling Remote method Add with")
        fmt.Printf("\nOperand 1 = 7")
        fmt.Printf("\nOperand 2 = 8")
 	//Call the remote procedure Add exposed on type Arithmatic
	err = c.Call("Arithmatic.Add", args, &reply)
	if err != nil {
	    log.Fatal("arith error:", err)
	}
	fmt.Printf("\nAnswer returned by the server = %d\n", reply)

}
