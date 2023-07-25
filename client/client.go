/**
 * @Author: sk shahriar ahemd raka <ahmed>
 * @Date:   1970-01-01T06:00:00+06:00
 * @Email:  skshahriarahmedraka@gmail.com
 * @Filename: main.go
 * @Last modified by:   ahmed
 * @Last modified time: 2021-08-28T22:58:34+06:00
 */

package main

import (
	"context"
	// "crypto/tls"
	// "flag"
	"fmt"
	// "log"
	// "net"

	//"os"
	"client/proto"
	"time"

	"client/logs"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
)





func main() {
	
	conn, err := grpc.Dial("grpc-knative-cluster4.default.127.0.0.1.sslip.io:80", grpc.WithInsecure(), grpc.WithBlock())
	logs.Error("Error in connecting", err)
	defer conn.Close()
	fmt.Printf("%T\n",conn)
	ClientConnection(conn)

}

func ClientConnection(conn *grpc.ClientConn) {

	c := messagepb.NewConversationClient(conn)
    fmt.Println("🚀ClientConnection  : ", c)
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	r, err := c.Speaking(ctx, &messagepb.SpeakRequest{Client_Request: "raka"})

	logs.Error("Error in getting response", err)
	fmt.Println("server response TO client : ", r.GetServer_Response())
}
