package main

import (
	"context"
	"fmt"
	"log"
	"tradeClient/internal/config"
	"tradeClient/internal/handler"
	"tradeClient/internal/handler/echoHTTP"
	prStor "tradeClient/internal/priceStorage"
	"tradeClient/internal/service"

	protocPosition "github.com/Kamieshi/position_service/protoc"
	protocPrice "github.com/Kamieshi/price_service/protoc"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		logrus.WithError(err).Fatalf("Error parse config")
	}
	positionConnect, err := grpc.Dial(conf.PositionServerRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	priceConnect, err := grpc.Dial(conf.PriceServerRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	priceRPC := protocPrice.NewOwnPriceStreamClient(priceConnect)
	positionRPC := protocPosition.NewPositionsManagerClient(positionConnect)
	userRPC := protocPosition.NewClientsManagerClient(positionConnect)

	handlerPrice := handler.PriceRPC{PriceManagerClient: priceRPC}
	handlerPosition := handler.PositionRPC{PositionManagerClient: positionRPC}
	handlerUser := handler.UserRPC{UserManagerClient: userRPC}

	priceStorage := prStor.NewPriceStorage(&handlerPrice)
	go priceStorage.ListenCompanyChanel(context.Background())

	userService := service.UserService{
		UserHandler: &handlerUser,
	}

	positionService := service.PositionService{
		PositionHandler: &handlerPosition,
		PriceStorage:    priceStorage,
	}

	userHandlerHTTP := echoHTTP.UserHandler{UserService: &userService}
	priceHandlerHTTP := echoHTTP.PriceHandler{PriceStorage: priceStorage}
	positionHandlerHTTP := echoHTTP.PositionHandler{PositionService: &positionService}

	e := echo.New()

	userGroup := e.Group("/user")
	userGroup.GET("", userHandlerHTTP.GetAll)

	positionGroup := e.Group("/position")
	positionGroup.GET("", positionHandlerHTTP.GetAll)

	priceGroup := e.Group("/price")
	priceGroup.GET("", priceHandlerHTTP.GetAll)
	logrus.Info(e.Start(fmt.Sprintf(":%s", conf.HttpEchoPort)))
}
