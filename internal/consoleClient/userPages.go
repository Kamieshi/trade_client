package consoleClient

import (
	"context"
	"fmt"
	"tradeClient/internal/service"

	"github.com/pterm/pterm"
)

type BaseUserPage struct {
	BasePage
	seanceData  *SeanceData
	UserService *service.UserService
}

func NewBaseUserPage(sData *SeanceData, userService *service.UserService) *BaseUserPage {
	return &BaseUserPage{
		UserService: userService,
		seanceData:  sData,
	}
}

func (m *BaseUserPage) Show() {
	options := []string{
		". . .",
		"Create New User",
		"Select user from exist list",
	}
	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	if selectedOption == options[0] {
		m.NextPage = nil
	}
	if selectedOption == options[1] {
		m.NextPage = NewCreateUserPage(*m)
	}
	if selectedOption == options[2] {
		m.NextPage = NewSelectUserPage(*m)
	}
	pterm.Info.Println("MainPage")
}

type CreateNewUserPage struct {
	BaseUserPage
}

func NewCreateUserPage(baseUserPage BaseUserPage) *CreateNewUserPage {
	newPage := &CreateNewUserPage{
		BaseUserPage: baseUserPage,
	}
	newPage.NamePage += " / Create new user"
	return newPage
}

func (c *CreateNewUserPage) Show() {
	username, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("User name").Show()
	balance, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Balance").Show()

	pterm.Println() // Blank line
	pterm.Info.Printfln("You answered: %s", result)
}

type SelectUserPage struct {
	BaseUserPage
}

func NewSelectUserPage(baseUserPage BaseUserPage) *SelectUserPage {
	newPage := &SelectUserPage{
		BaseUserPage: baseUserPage,
	}
	newPage.NamePage += " / Select new user"
	return newPage
}

func (s *SelectUserPage) Show() {
	users, err := s.UserService.GetAll(context.Background())
	if err != nil {
		pterm.Error.Println("Error : ", err.Error())
	}
	selectMenu := make([]string, 0, len(users))
	reverseIndexMap := make(map[string]int)
	for i, user := range users {
		stringView := fmt.Sprintf("Name : %s ,Balsance : %d (id=%s)", user.Name, user.Balance, user.ID)
		reverseIndexMap[stringView] = i
		selectMenu = append(selectMenu, stringView)
	}
	selectedUser, _ := pterm.DefaultInteractiveSelect.WithOptions(selectMenu).Show()
	s.seanceData.CurrentUser = users[reverseIndexMap[selectedUser]]
	pterm.Info.Printfln("Current user : %s", users[reverseIndexMap[selectedUser]])
}

func (s *SelectUserPage) Next() Page {
	return nil
}
