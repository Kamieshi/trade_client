package consoleClient

import (
	"tradeClient/internal/model"
	"tradeClient/internal/priceStorage"
	"tradeClient/internal/service"

	"github.com/pterm/pterm"
)

type SeanceData struct {
	CurrentUser *model.User
}

type MainPage struct {
	BasePage
	seanceData      *SeanceData
	PriceStorage    *priceStorage.PriceStorage
	UserService     *service.UserService
	PositionService *service.PositionService
}

func (m *MainPage) Show() {
	pterm.Info.Println("MainPage")
	options := []string{
		"Users",
		"Positions",
		"Prices",
	}
	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	if selectedOption == options[0] {
		m.NextPage = NewBaseUserPage(m.seanceData, m.UserService)
	}

	if selectedOption == options[1] {
		if m.seanceData.CurrentUser == nil {
			pterm.Error.Println("Not selected current user")
			m.NextPage = m
			return
		}
		m.NextPage = NewBasePositionPage(m.seanceData, m.PositionService, m.PriceStorage)
	}

	if selectedOption == options[2] {
		m.NextPage = NewBasePricePage(m.PriceStorage)
	}
}
