package consoleClient

import (
	"tradeClient/internal/priceStorage"
	"tradeClient/internal/service"

	"github.com/pterm/pterm"
)

type BasePositionPage struct {
	BasePage
	seanceData      *SeanceData
	PositionService *service.PositionService
	PriceStorage    *priceStorage.PriceStorage
}

func NewBasePositionPage(seData *SeanceData, positionService *service.PositionService, prStorage *priceStorage.PriceStorage) *BasePositionPage {
	return &BasePositionPage{
		BasePage:        BasePage{NamePage: "Positions Page"},
		PositionService: positionService,
		PriceStorage:    prStorage,
		seanceData:      seData,
	}
}

func (b *BasePositionPage) Show() {
	options := []string{
		". . .",
		"Open Position",
		"Close Position",
		"Info About Opened User Positions",
	}
	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	pterm.Info.Println(selectedOption)
}
