// Package echoHTTP Handlers from echo
package echoHTTP

import (
	"net/http"

	"github.com/Kamieshi/trade_client/internal/model"
	"github.com/Kamieshi/trade_client/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// PositionHandler HTTP echo handler from position
type PositionHandler struct {
	PositionService *service.PositionService
}

// ClosePosition Close position
// @Tags         position
// @Param        positionID path string true "position ID"
// @Param        userID path string true "userID ID"
// @Success      200  {object} model.Position
// @Failure      400  string true "Err"
// @Router       /position/close/{userID}/{positionID} [get]
func (p *PositionHandler) ClosePosition(ctx echo.Context) error {
	userID := ctx.Param("userID")
	positionID := ctx.Param("positionID")
	if userID == "" || positionID == "" {
		logrus.Error("position handler / ClosePosition / Error Parse company id")
		return ctx.String(http.StatusBadRequest, "Invalid position or user IDs")
	}
	position := model.Position{ID: positionID, UserID: userID, Price: &model.Price{}}
	_, err := p.PositionService.ClosePosition(ctx.Request().Context(), &position)
	if err != nil {
		logrus.Error("position handler / ClosePosition / close position form position service")
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, position)
}

// OpenPosition open position
// @Tags         position
// @Param        companyID path string true "company ID"
// @Param 		 position body model.Position true "postition for open"
// @Success      200  {object}  model.Position
// @Failure      400  string true "Err"
// @Router       /position/{companyID}/open [post]
func (p *PositionHandler) OpenPosition(ctx echo.Context) error {
	companyID := ctx.Param("companyID")
	if companyID == "" {
		logrus.Error("position handler / OpenPosition / Error Parse company id")
		return ctx.String(http.StatusBadRequest, "Invalid company ID")
	}
	position := model.Position{
		Price: &model.Price{CompanyID: companyID},
	}
	err := ctx.Bind(&position)
	if err != nil {
		logrus.WithError(err).Error("position handler / OpenPosition / Error bind position")
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	err = p.PositionService.OpenPosition(ctx.Request().Context(), &position)
	if err != nil {
		logrus.WithError(err).Error("position handler / OpenPosition / Open position from service")
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, position)
}

// GetPosition Get position by ID
// @Tags         position
// @Param        positionID path string true "position ID"
// @Success      200  {object}  model.Position
// @Failure      400  string true "Err"
// @Router       /position/{positionID} [get]
func (p *PositionHandler) GetPosition(ctx echo.Context) error {
	positionID := ctx.Param("positionID")
	if positionID == "" {
		logrus.Errorf("Parse error, not fount PositionID")
		return ctx.String(http.StatusBadRequest, "Invalid ID")
	}
	position, err := p.PositionService.GetPositionByID(ctx.Request().Context(), positionID)
	if err != nil {
		logrus.WithError(err).Error("position handler / GetPosition / Get position form service")
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, position)
}

// GetAllUserPosition get all positions for user
// @Tags         position
// @Param        userID path string true "user ID"
// @Success      200  {array}  model.Position
// @Failure      400  string true "Err"
// @Router       /position/user/{userID} [get]
func (p *PositionHandler) GetAllUserPosition(ctx echo.Context) error {
	userID := ctx.Param("userID")
	if userID == "" {
		logrus.Errorf("Parse error, not fount PositionID")
		return ctx.String(http.StatusBadRequest, "Invalid ID")
	}
	positions, err := p.PositionService.GetAllUserPosition(ctx.Request().Context(), userID)
	if err != nil {
		logrus.WithError(err).Error("position handler / GetAllUserPosition / Get positions form service")
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, positions)
}
