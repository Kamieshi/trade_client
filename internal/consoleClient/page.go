package consoleClient

type Page interface {
	Show()
	Next() Page
}
