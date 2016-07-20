/*
 * A simple HTTP server listening requests from HTTP client using default GOB wire format
*/

package main
import (
        "fmt"
        "net/rpc"
        "net/http"
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
    //Register an RPC object
    rpc.Register(arithmatic)
    //Register an HTTP handler
    rpc.HandleHTTP()
    fmt.Printf("\nAn HTTP server listening requests from an HTTP client ... ")
    //Listen on tcp address and call serve to serve the request
    http.ListenAndServe(":1234", nil)     
}
