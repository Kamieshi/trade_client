package priceStorage

import (
	"context"
	"fmt"
	"sync"
	"time"
	"tradeClient/internal/handler"
	"tradeClient/internal/model"

	"github.com/Kamieshi/price_service/protoc"
	"github.com/sirupsen/logrus"
)

type PriceStorage struct {
	PositionHandler *handler.PriceRPC
	mutex           sync.RWMutex
	Prices          map[string]*model.Price
}

func NewPriceStorage(ph *handler.PriceRPC) *PriceStorage {
	return &PriceStorage{
		PositionHandler: ph,
		Prices:          make(map[string]*model.Price),
	}
}

// ListenCompanyChanel G from listen end update price storage
func (p *PriceStorage) ListenCompanyChanel(ctx context.Context) {
	stream, err := p.getPriceStream(ctx)
	if err != nil {
		logrus.WithError(err).Fatalf("service position / ListenCompanyChanel / get stream from handler")
	}
	var bufferCompany *model.Price
	for {
		select {
		case <-ctx.Done():
			return
		default:
			data, err := stream.Recv()
			if err != nil {
				logrus.WithError(err).Error("service price / ListenCompanyChanel / get data from stream")
			}
			timeParse, err := time.Parse("2006-01-02T15:04:05.000TZ-07:00", data.Time)
			if err != nil {
				logrus.WithError(err).Error("service price / ListenCompanyChanel / parse time")
			}

			bufferCompany.CompanyID = data.Company.ID
			bufferCompany.Name = data.Company.Name
			bufferCompany.Ask = data.Ask
			bufferCompany.Bid = data.Bid
			bufferCompany.Time = timeParse

			p.setPrice(bufferCompany)

		}

	}
}

func (p *PriceStorage) getPriceStream(ctx context.Context) (protoc.OwnPriceStream_GetPriceStreamClient, error) {
	stream, err := p.PositionHandler.GetPriceStream(ctx)
	if err != nil {
		return nil, fmt.Errorf("service price / GetStream / get stream from handler : %v", err)
	}
	return stream, err
}

func (p *PriceStorage) setPrice(company *model.Price) {
	p.mutex.Lock()
	if _, ex := p.Prices[company.CompanyID]; !ex {
		p.Prices[company.CompanyID] = &model.Price{
			CompanyID: company.CompanyID,
			Name:      company.Name,
		}
	}
	companyFromMap := p.Prices[company.CompanyID]
	companyFromMap.Time = company.Time
	companyFromMap.Ask = company.Ask
	companyFromMap.Bid = company.Bid
	p.mutex.Unlock()
}

// GetPrice Return actual price
func (p *PriceStorage) GetPrice(companyID string) (model.Price, error) {
	var price *model.Price
	p.mutex.RLock()
	price, exist := p.Prices[companyID]
	p.mutex.RUnlock()
	if !exist {
		return *price, fmt.Errorf("price service / GetPrice / price for company %s not exist", companyID)
	}
	return *price, nil
}

// GetAllPrices Return all actual prices
func (p *PriceStorage) GetAllPrices() []model.Price {
	p.mutex.Lock()
	prices := make([]model.Price, 0, len(p.Prices))
	for _, val := range p.Prices {
		prices = append(prices, *val)
	}
	p.mutex.Unlock()
	return prices
}
