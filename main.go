package main

import (
	"context"
	"log"
	"tradeClient/internal/service"

	protocPosition "github.com/Kamieshi/position_service/protoc"
	protocPrice "github.com/Kamieshi/price_service/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	positionConnect, err := grpc.Dial("localhost:5301", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	priceConnect, err := grpc.Dial("localhost:5300", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	priceRPC := protocPrice.NewOwnPriceStreamClient(priceConnect)
	positionRPC := protocPosition.NewPositionsManagerClient(positionConnect)
	clientRPC := protocPosition.NewClientsManagerClient(positionConnect)
	positionStream, err := priceRPC.GetPriceStream(context.Background(), &protocPrice.GetPriceStreamRequest{})
	if err != nil {
		log.Fatalf("Get stream from price service: %v", err)
	}
	companyService := service.NewCompanyService(context.Background(), positionStream)
}
