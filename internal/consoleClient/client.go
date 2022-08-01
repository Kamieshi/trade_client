package consoleClient

import (
	"tradeClient/internal/priceStorage"
	"tradeClient/internal/service"
)

type ConsoleClient struct {
	Pages      []Page
	activePage int
}

func NewConsoleClient(userService *service.UserService, prStorage *priceStorage.PriceStorage, positionService *service.PositionService) *ConsoleClient {
	consoleClient := &ConsoleClient{
		Pages:      make([]Page, 0, 1),
		activePage: 0,
	}
	consoleClient.AddPage(&MainPage{
		UserService:     userService,
		PriceStorage:    prStorage,
		PositionService: positionService,
		seanceData:      &SeanceData{},
	})
	return consoleClient
}

func (c *ConsoleClient) Run() {
	for {
		c.ShowPage(c.Pages[len(c.Pages)-1])
		if c.Pages[len(c.Pages)-1].Next() != nil {
			c.AddPage(c.Pages[len(c.Pages)-1].Next())
		} else {
			c.ReturnPreviewPage()
		}
	}
}

func (c *ConsoleClient) AddPage(page Page) {
	c.Pages = append(c.Pages, page)
}

func (c *ConsoleClient) ShowPage(page Page) {
	page.Show()
}

func (c *ConsoleClient) ReturnPreviewPage() {
	c.Pages = c.Pages[:len(c.Pages)-1]
}
