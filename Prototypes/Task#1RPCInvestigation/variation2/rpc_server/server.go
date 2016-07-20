/*
 * An HTTP server serving json encoded requests.
*/

package main
import (
        "fmt"
        "net/rpc"
        "net/rpc/jsonrpc"
        "log"
        "net"
        "rpc_def"
)


type Arithmatic int

func (t *Arithmatic) Add(args rpc_def.Args, answer *int) error {
    fmt.Printf("\n\nProcedure Add called by the remote client")
    fmt.Printf("\nOperand 1 = %d", args.Operand1)
    fmt.Printf("\nOperand 2 = %d", args.Operand2)
    *answer = args.Operand1 + args.Operand2
    fmt.Printf("\nAnswer = %d\n", *answer)
    return nil
}

func main(){

    arith := new(Arithmatic)

    // Create new server object
    server := rpc.NewServer()
    //Publish methods on arith
    server.Register(arith)
    //Register HTTP handler for RPC messages
    server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
    //Listen on tcp network for connections
    l, e := net.Listen("tcp", ":1234")
    if e != nil {
        log.Fatal("listen error:", e)
    }
    fmt.Printf("\nAn HTTP RPC server.")
    fmt.Printf("\nAccepts JSON encoded requests.")
    for 
	{
        //Accept connections on network
        conn, err := l.Accept()
        if err != nil {
            log.Fatal(err)
    }
    //runs rpc server on single connection
    go server.ServeCodec(jsonrpc.NewServerCodec(conn))
  }
}
