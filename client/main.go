package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":8080"

func main() {
	con, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer con.Close()

	//client:=pb.NewGreetServiceClient(con)

	// names:=&pb.Namelist{
	// 	Names:[]string{"Aytaditya","Budi","Citra"},
	// }

	//callSayHello(client)
}
