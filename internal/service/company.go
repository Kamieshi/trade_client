package service

import (
	"context"
	"sync"
	"time"
	"tradeClient/internal/model"

	"github.com/Kamieshi/price_service/protoc"
	"github.com/sirupsen/logrus"
)

type CompanyService struct {
	Companies map[string]*model.Company
	sync.RWMutex
}

func NewCompanyService(ctx context.Context, stream protoc.OwnPriceStream_GetPriceStreamClient) *CompanyService {
	compService := &CompanyService{
		Companies: make(map[string]*model.Company),
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				req, err := stream.Recv()
				if err != nil {
					logrus.WithError(err).Fatal("company service / NewCompanyService goroutine / err from Recv()")
				}
				parsTime, err := time.Parse("2006-01-02T15:04:05.000TZ-07:00", req.Time)
				compService.Lock()
				compService.Companies[req.Company.ID] = &model.Company{
					ID:   req.Company.ID,
					Name: req.Company.Name,
					Ask:  req.Ask,
					Bid:  req.Bid,
					Time: parsTime,
				}
				compService.Unlock()
			}
		}
	}()
	return compService
}
