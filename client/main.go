package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "syahdan.id/coinbit/wallet"
)

func getDataWalletById(client pb.WalletClient, id int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := pb.WalletRequest{WalletId: id}
	wallet, err := client.FindWalletById(ctx, &s)

	if err != nil {
		log.Fatalln("error when get wallet", err.Error())
	}

	data, err := json.Marshal(wallet)
	fmt.Print(string(data))
}

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":1200", opts...)
	if err != nil {
		log.Fatalln("error in dial")
	}

	defer conn.Close()

	client := pb.NewWalletClient(conn)
	getDataWalletById(client, 1)
}
