/*
 * A simple and the recommended RPC server implementation on TCP layer using built-in GOB as a wire format
*/

package main
import (
        "fmt"
        "net/rpc"
        "net"
        "rpc_def"
        "log"
)

type Arithmatic int

func (t *Arithmatic) Add(args rpc_def.Args, answer *int) error {
    fmt.Printf("\nProcedure Add called by the remote client\n")
    fmt.Printf("\n Operand 1 = %d", args.Operand1)
    fmt.Printf("\n Operand 2 = %d", args.Operand2)
    *answer = args.Operand1 + args.Operand2
    fmt.Printf("\nAnswer = %d\n", *answer)
    return nil
}

func main(){

    arithmatic := new(Arithmatic)
    rpc.Register(arithmatic)
    l, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("listen error:", err)
   }
   fmt.Printf("\nThe simple RPC server\n Listening for new connections...")
   rpc.Accept(l)
}
