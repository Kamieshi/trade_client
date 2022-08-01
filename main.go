package main

import (
	"context"
	"log"
	"tradeClient/internal/config"
	"tradeClient/internal/handler"
	"tradeClient/internal/priceStorage"

	protocPosition "github.com/Kamieshi/position_service/protoc"
	protocPrice "github.com/Kamieshi/price_service/protoc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		logrus.WithError(err).Fatalf("Error parse config")
	}
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
	userRPC := protocPosition.NewClientsManagerClient(positionConnect)
	handlerPrice := handler.PriceRPC{PriceManagerClient: priceRPC}
	handlerPosition := handler.PositionRPC{PositionManagerClient: positionRPC}
	handlerUser := handler.UserRPC{UserManagerClient: userRPC}
	priceService := priceStorage.NewPriceStorage(&handlerPrice)
	go priceService.ListenCompanyChanel(context.Background())

}
