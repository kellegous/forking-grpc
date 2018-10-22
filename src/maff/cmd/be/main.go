package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"maff/pkg"
)

// MaffServer ...
type MaffServer struct{}

// GetProduct ...
func (s *MaffServer) GetProduct(
	ctx context.Context,
	req *pkg.ProductReq) (*pkg.ProductRes, error) {
	t := time.Now()
	var p int64 = 1
	for _, i := range req.Items {
		p *= i
	}
	return &pkg.ProductRes{
		Product:  p,
		Duration: time.Now().Sub(t).Nanoseconds(),
	}, nil
}

func main() {
	flagAddr := flag.String("addr", ":9090",
		"address for grpc endpoint")

	flag.Parse()

	l, err := net.Listen("tcp", *flagAddr)
	if err != nil {
		log.Panic(err)
	}

	s := grpc.NewServer()
	pkg.RegisterMaffServer(s, &MaffServer{})
	log.Panic(s.Serve(l))
}
