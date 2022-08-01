package consoleClient

import (
	"github.com/pterm/pterm"
)

type BasePage struct {
	NamePage string
	NextPage Page
}

func (b *BasePage) Show() {
	pterm.Info.Println(b.NamePage)
}

func (b *BasePage) Next() Page {
	return b.NextPage
}
