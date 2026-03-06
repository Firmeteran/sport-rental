package controller

import (
	"net/http"
	"sport-rental/service"

	"github.com/labstack/echo/v4"
)

type TopUpController struct {
	svc service.TopUpService
}

func NewTopUpController(s service.TopUpService) *TopUpController {
	return &TopUpController{s}
}

func (h *TopUpController) RequestTopUp(c echo.Context) error {
	type TopUpRequest struct {
		UserID uint    `json:"user_id"`
		Amount float64 `json:"amount"`
	}

	req := new(TopUpRequest)

	// Bind JSON to body
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid JSON format."})
	}

	// Call service
	result, err := h.svc.CreateTopUp(int(req.UserID), req.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":      "success",
		"message":     "Top up request successfully created.",
		"payment_url": result.SnapURL,
		"order_id":    result.OrderID,
	})
}
