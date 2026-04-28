package main

import (
	"context"
	"log"
	"time"

	inventory "go-micro-pet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Устанавливаем соединение с сервером
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := inventory.NewInventoryClient(conn)

	// Делаем запрос
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	itemID := "apple"
	res, err := client.CheckStock(ctx, &inventory.StockRequest{ItemId: itemID})
	if err != nil {
		log.Fatalf("Ошибка запроса: %v", err)
	}

	if res.IsAvailable {
		log.Printf("Успех! Товар %s в наличии. Заказ создан.", itemID)
	} else {
		log.Printf("Отмена: Товара %s нет на складе.", itemID)
	}
}
