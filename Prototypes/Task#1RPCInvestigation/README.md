RPC Investigtion: RPC server and client usage for all the variants
==================================================================

This document assumes that you already have your go application structure(bin, pkg, src) in place as suggested at
https://golang.org/doc/code.html and your go tool is set to use (GOPATH and other required env variables are set).

In all the RPC variations, the server exposes the method 'Add' and returns the addition of two numbers passed by the client
as arguments. The arguments, server IP and port are hardcoded inside the client code and can easily be make out looking 
at the source code. As server IP and port are hardcoded, server and client should be run on the same machine/VM though they 
are tested and work well on different VMs. 

Below are the steps to execute each variant. Each client should call the remote procedure exposed by the respective server only.

1. Copy packages rpc_server, rpc_client, rpc_def from the same variant to your $GOPATH/src/ folder.
2. Build and create binaries

        go install rpc_def/
        go install rpc_server/
        go install rpc_client/

3. Run the server

        $GOPATH/bin/rpc_server
        
   This will start the server and prints the message on the console. It is recommended not to run the server in the
   background as it prints the result on the console only.
4. Run the client

        $GOPATH/bin/rpc_client
   See the results on both client and server console.
   
   Repeat steps 1 to 4 if you make any changes to the source code.