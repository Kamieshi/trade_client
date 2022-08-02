package main

import (
	"context"
	"fmt"
	"log"

	protocPosition "github.com/Kamieshi/position_service/protoc"
	protocPrice "github.com/Kamieshi/price_service/protoc"
	_ "github.com/Kamieshi/trade_client/docs"
	"github.com/Kamieshi/trade_client/internal/config"
	"github.com/Kamieshi/trade_client/internal/handler"
	"github.com/Kamieshi/trade_client/internal/handler/echoHTTP"
	prStor "github.com/Kamieshi/trade_client/internal/priceStorage"
	"github.com/Kamieshi/trade_client/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title Swagger Client Trade service

// @contact.url https://github.com/Kamieshi

// @host localhost:8080
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
	userRPC := protocPosition.NewUsersManagerClient(positionConnect)

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
	userGroup.POST("", userHandlerHTTP.CreateUser)
	userGroup.GET("/:userName", userHandlerHTTP.Get)
	userGroup.GET("/:userID/updateBalance/:difference", userHandlerHTTP.UpdateBalance)

	positionGroup := e.Group("/position")
	positionGroup.GET("/user/:userID", positionHandlerHTTP.GetAllUserPosition)
	positionGroup.GET("/:positionID", positionHandlerHTTP.GetPosition)
	positionGroup.POST("/:companyID/open", positionHandlerHTTP.OpenPosition)
	positionGroup.GET("/close/:userID/:positionID", positionHandlerHTTP.ClosePosition)

	priceGroup := e.Group("/price")
	priceGroup.GET("", priceHandlerHTTP.GetAll)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	logrus.Info(e.Start(fmt.Sprintf(":%s", conf.HttpEchoPort)))
}
