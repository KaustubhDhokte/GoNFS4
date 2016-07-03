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
    fmt.Printf("\nInside Add\n")
    *answer = args.Operand1 + args.Operand2
    fmt.Printf("answer = %d", *answer)
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
    //Accept each connection and serve
    rpc.Accept(inbound)
}
