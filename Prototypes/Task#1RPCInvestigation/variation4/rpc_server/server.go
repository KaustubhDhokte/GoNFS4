/*
 * A TCP RPC server serving multiple connections and using default GOB as wire format
*/

package main
import (
        "fmt"
        "net/rpc"
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
    //Create an object to listen on TCP 
    addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:1234")
    if err != nil {
		log.Fatal(err)
    }
    //Listen on TCP
    inbound, err := net.ListenTCP("tcp", addy)
    if err != nil {
		log.Fatal(err)
    }
    //Register on RPC service
    rpc.Register(arithmatic)
    fmt.Printf("\nA TCP RPC server. Listening for requests on TCP layer ...")
    //Accept each connection and serve
    rpc.Accept(inbound)
}
