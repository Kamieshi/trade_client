package consoleClient

import (
	"fmt"
	"tradeClient/internal/priceStorage"

	"github.com/pterm/pterm"
)

type BasePricePage struct {
	BasePage
	PriceStorage *priceStorage.PriceStorage
}

func NewBasePricePage(prStorage *priceStorage.PriceStorage) *BasePricePage {
	return &BasePricePage{
		BasePage:     BasePage{NamePage: "Price page"},
		PriceStorage: prStorage,
	}
}

func (b *BasePricePage) Show() {
	prices := b.PriceStorage.GetAllPrices()
	tableData := pterm.TableData{{"id", "Name", "Bid", "Ask", "Time"}}

	for _, price := range prices {
		tableData = append(tableData, []string{price.CompanyID, price.Name, fmt.Sprint(price.Bid), fmt.Sprint(price.Ask), fmt.Sprint(price.Time)})
	}
	pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
	pterm.Println()
}
