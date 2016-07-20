/*
 * A simple TCP server .... listening to only one request using GOB as wire format
*/

package main
import (
        "fmt"
        "net/rpc"
        "net"
        "rpc_def"
)

type Arithmatic int

func (t *Arithmatic) Add(args rpc_def.Args, answer *int) error {
    fmt.Printf("\nProcedure Add called by the remote client")
    fmt.Printf("\nOperand 1 = %d", args.Operand1)
    fmt.Printf("\nOperand 2 = %d", args.Operand2)
    *answer = args.Operand1 + args.Operand2
    fmt.Printf("\nAnswer = %d", *answer)
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
    fmt.Printf("\nA Single connection TCP RPC server. Will serve single request and terminate. Listening...")
    conn, err := listener.Accept()
    //Serve the requests on conn
    rpc.ServeConn(conn)

}
