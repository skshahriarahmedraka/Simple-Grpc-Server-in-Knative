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
	// "context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"

	"golang.org/x/net/context"

	//"os"
	"client/proto"
	"time"

	"client/logs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// func Panicking(msg string , err error){
// 	if err!= nil {
// 		log.Fatalln(msg ," ", err)
// 	}
// }

// func main () {
// 	conn , err := grpc.Dial(":50051",grpc.WithInsecure(),grpc.WithBlock())

// 	Panicking("error : ",err)
// 	defer conn.Close()

// 	c:=messagepb.NewConversationClient(conn)

// 	name:= "raka"
// 	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
// 	defer cancel()

// 	r ,err:= c.Speaking(ctx,&messagepb.SpeakRequest{MyRequest:name})

// 	Panicking("error ",err)
// 	log.Println("greeting :",r.GetMyResponse())

// }

var (
	serverAddr         = flag.String("server_addr", "127.0.0.1:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "", "")
	insecure           = flag.Bool("insecure", false, "Set to true to skip SSL validation")
	skipVerify         = flag.Bool("skip_verify", false, "Set to true to skip server hostname verification in SSL validation")
)

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	if *serverHostOverride != "" {
		opts = append(opts, grpc.WithAuthority(*serverHostOverride))
	}
	if *insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		cred := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: *skipVerify,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	// conn, err := grpc.Dial("grpc-knative-cluster1.knative-serving.example.com:50051", grpc.WithInsecure(), grpc.WithBlock())
	// logs.Error("Error in connecting", err)
	// defer conn.Close()
	//fmt.Printf("%T",conn)
	ClientConnection(conn)

}

func ClientConnection(conn *grpc.ClientConn) {

	c := messagepb.NewConversationClient(conn)
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Speaking(ctx, &messagepb.SpeakRequest{Client_Request: "raka"})

	logs.Error("Error in getting response", err)
	fmt.Println("server response TO client : ", r.GetServer_Response())
}
