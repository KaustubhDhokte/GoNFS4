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
    fmt.Printf("\nInside Add\n")
    *answer = args.Operand1 + args.Operand2
    fmt.Printf("answer = %d", *answer)
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
    // Accept connections
    conn, err := listener.Accept()
    if err != nil {
            fmt.Printf("Error : %s", err)
    }
    //Serve on the TCP connection
    jsonrpc.ServeConn(conn)
}
