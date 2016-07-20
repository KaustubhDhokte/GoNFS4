/*
 * A single connection TCP json rpc server
*/
package main
import (
        "fmt"
        "net/rpc"
        "net"
        "rpc_def"
        "net/rpc/jsonrpc"
)

type Arithmatic int

func (t *Arithmatic) Add(args rpc_def.Args, answer *int) error {
    fmt.Printf("\nProcedure Add called by the remote client")
    fmt.Printf("\nOperand 1 = %d", args.Operand1)
    fmt.Printf("\nOperand 2 = %d", args.Operand2)
    *answer = args.Operand1 + args.Operand2
    fmt.Printf("\nAnswer = %d\n", *answer)
    return nil
}

func main(){

    arithmatic := new(Arithmatic)
    //Register an RPC 
    rpc.Register(arithmatic)
    //Create TCP listener object
    tcpAddr, err  := net.ResolveTCPAddr("tcp", ":1234")
    //Listen on TCP listener object
    listener, err  := net.ListenTCP("tcp", tcpAddr)
    fmt.Printf("\nA TCP single connection json rpc server\n")
    fmt.Printf("\nWill serve only a single connection. Serve only JSON encoded requests.\n")
    fmt.Printf("\nListening for requests ...")
    // Accept connections
    conn, err := listener.Accept()
    if err != nil {
            fmt.Printf("Error : %s", err)
    }
    //Serve on the TCP connection
    jsonrpc.ServeConn(conn)
}
