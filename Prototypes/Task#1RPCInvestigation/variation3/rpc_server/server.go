/*
 * RPC server accepting requests on HTTP layer using default GOB as the wire format
*/
package main
import (
        "fmt"
        "net/rpc"
        "net/http"
        "log"
        "net"
        "rpc_def"
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
    //Registers rpc server object
    rpc.Register(arithmatic)
    //Register HTTP handler fo RPC
    rpc.HandleHTTP()
    //Listen connections on TCP
    l, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("listen error:", err)
    }
    fmt.Printf("\nAn HTTP RPC server. Listening rpc requests by an HTTP client...")
    //Accepts new HTTP connection on listener and runs the go service routine for rpc request
    http.Serve(l, nil)
}

