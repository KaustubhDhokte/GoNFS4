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
    fmt.Printf("\nInside Add\n")
    *answer = args.Operand1 + args.Operand2
    fmt.Printf("answer = %d", *answer)
    return nil
}

func main(){

    arithmatic := new(Arithmatic)
    rpc.Register(arithmatic)
    l, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("listen error:", err)
   }
   rpc.Accept(l)
}
