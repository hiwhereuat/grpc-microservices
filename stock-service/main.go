package main

import (
	"context"
	"log"
	"net"

	inventory "go-micro-pet/proto"

	"google.golang.org/grpc"
)

type server struct {
	inventory.UnimplementedInventoryServer
}

func (s *server) CheckStock(ctx context.Context, in *inventory.StockRequest) (*inventory.StockResponse, error) {
	log.Printf("Проверка товара: %s", in.ItemId)
	available := in.ItemId == "apple"
	return &inventory.StockResponse{IsAvailable: available}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка порта: %v", err)
	}
	s := grpc.NewServer()
	inventory.RegisterInventoryServer(s, &server{})
	log.Println("Stock Service запущен на порту :50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Ошибка запуска: %v", err)
	}
}
