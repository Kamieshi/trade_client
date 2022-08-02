package echoHTTP

import (
	"net/http"
	"strconv"

	"github.com/Kamieshi/trade_client/internal/model"
	"github.com/Kamieshi/trade_client/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// UserHandler HTTP echo handler from users
type UserHandler struct {
	UserService *service.UserService
}

// Get get user by name
// @Tags         user
// @Param        userName path string true "Username"
// @Success      200  {object}  model.User
// @Failure      404  string true "Not found User"
// @Router       /user/{userName} [get]
// Get get user by user Name
func (c *UserHandler) Get(ctx echo.Context) error {
	userName := ctx.Param("userName")
	user, err := c.UserService.GetByName(ctx.Request().Context(), userName)
	if err != nil {
		logrus.WithError(err).Errorf("User handler echo / Get / Get user {%s} from service ", userName)
		return ctx.String(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, user)
}

// GetAll get all users
// @Tags         user
// @Success      200  {object}  []model.User
// @Failure      400  string true "Err"
// @Router       /user [get]
// GetAll get user by user Name
func (c *UserHandler) GetAll(ctx echo.Context) error {
	users, err := c.UserService.GetAll(ctx.Request().Context())
	if err != nil {
		logrus.WithError(err).Error("User handler echo / GetAll / Get users from service ")
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, users)
}

// UpdateBalance update balance
// @Tags         user
// @Param        difference path string true "difference"
// @Param        user body model.User true "user"
// @Success      200  {string} string
// @Failure      404  string true "Not found User"
// @Router       /user/updateBalance/{difference} [post]
// UpdateBalance get user by user Name
func (c *UserHandler) UpdateBalance(ctx echo.Context) error {
	user := new(model.User)
	different, err := strconv.ParseInt(ctx.Param("difference"), 10, 64)
	if err != nil {
		logrus.WithError(err).Errorf("User handler echo / UpdateBalance / Parse %s to int", ctx.Param("difference"))
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	err = ctx.Bind(&user)
	if err != nil {
		logrus.WithError(err).Error("user handler echo / UpdateBalance / Bind user")
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	err = c.UserService.UpdateBalance(ctx.Request().Context(), user, different)
	if err != nil {
		logrus.WithError(err).Error("user handler echo / UpdateBalance / Update balance from service")
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.String(http.StatusAccepted, "")
}

// CreateUser create new user
// @Tags         user
// @Param        user body model.User true "user"
// @Success      200  {object} model.User
// @Failure      400  string true "bad Request"
// @Router       /user [post]
// CreateUser get user by user Name
func (c *UserHandler) CreateUser(ctx echo.Context) error {
	user := model.User{}
	err := ctx.Bind(&user)
	if err != nil {
		if err != nil {
			logrus.WithError(err).Error("user handler echo / CreateUser / Bind user")
			return ctx.String(http.StatusBadRequest, err.Error())
		}
	}
	err = c.UserService.CreateUser(ctx.Request().Context(), &user)
	if err != nil {
		logrus.WithError(err).Error("user handler echo / CreateUser / Create user from service")
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, user)
}
