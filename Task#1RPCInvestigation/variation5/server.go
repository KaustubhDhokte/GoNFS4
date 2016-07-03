package main
import (
        "fmt"
        "net/rpc"
        "net/http"
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
    //Register an RPC object
    rpc.Register(arithmatic)
    //Register an HTTP handler
    rpc.HandleHTTP()
    //Listen on tcp address and call serve to serve the request
    http.ListenAndServe(":1234", nil)     
}
