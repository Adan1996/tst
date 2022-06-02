package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	pb "syahdan.id/coinbit/wallet"
)

type dataWalletServer struct {
	pb.UnimplementedWalletServer
	mu      sync.Mutex
	wallets []*pb.WalletResponse
}

func (d *dataWalletServer) FindWalletById(ctx context.Context, wallet *pb.WalletRequest) (*pb.WalletResponse, error) {
	for _, v := range d.wallets {
		if v.WalletId == wallet.WalletId {
			return v, nil
		}
	}
	return nil, nil
}

func (d *dataWalletServer) loadData() {
	data, err := ioutil.ReadFile("data/wallets.json")

	if err != nil {
		log.Fatalln("error in read file", err.Error())
	}

	if err := json.Unmarshal(data, &d.wallets); err != nil {
		log.Fatalln("error in unmarshall data json", err.Error())
	}
}

func newServer() *dataWalletServer {
	s := dataWalletServer{}
	s.loadData()
	return &s
}

func main() {
	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalln("error in listen", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterWalletServer(grpcServer, newServer())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("error when server grpc serve grpc", err.Error())
	}
}
