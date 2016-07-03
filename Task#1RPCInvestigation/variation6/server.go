package main
import (
        "fmt"
        "net/rpc"
        "net"
        "rpc_def"
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
    //Register RPC service
    rpc.Register(arithmatic)
    //Resolve TCP address and create listener object
    tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
    //Listen on tcp listener
    listener, err := net.ListenTCP("tcp", tcpAddr)
    //Accept the connections
    conn, err := listener.Accept()
    //Serve the requests on conn
    rpc.ServeConn(conn)

}
