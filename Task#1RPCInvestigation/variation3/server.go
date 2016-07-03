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
    fmt.Printf("\nInside Add\n")
    *answer = args.Operand1 + args.Operand2
    fmt.Printf("answer = %d", *answer)
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
    //Accepts new HTTP connection on listener and runs the go service routine for rpc request
    http.Serve(l, nil)
}

