package main

import (

    "log"

    pb "github.com/jeremyko/my_proto"
)

func main() {

    myProto := &pb.MyProto{Msg:"some text", Number:999}

    log.Printf("msg: %s", myProto.GetMsg())
    log.Printf("number: %d", myProto.GetNumber())

}
