package controller

import (
	"net/http"
	"sport-rental/service"

	"github.com/labstack/echo/v4"
)

type RentalController struct {
	svc service.RentalService
}

func NewRentalController(s service.RentalService) *RentalController {
	return &RentalController{svc: s}
}

type ReturnRequest struct {
	ID int `json:"rental_id"`
}

func (h *RentalController) CreateRental(c echo.Context) error {
	// Struct for JSON request
	var input struct {
		UserID        int `json:"user_id"`
		EquipmentID   int `json:"equipment_id"`
		DurationHours int `json:"duration_hours"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request format."})
	}

	res, err := h.svc.RentEquipment(input.UserID, input.EquipmentID, input.DurationHours)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Rental attempt is successful.",
		"data":    res,
	})
}

func (h *RentalController) ReturnRental(c echo.Context) error {
	var req ReturnRequest

	// Bind body to JSON
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid JSON format.",
		})
	}

	// Make sure ID is not empty
	if req.ID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "ID is required in JSON body.",
		})
	}

	err := h.svc.ReturnEquipment(req.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "The tool has been successfully returned, and the stock has been updated.",
	})
}
