package controller

import (
	"net/http"
	"sport-rental/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RentalController struct {
	svc service.RentalService
}

func NewRentalController(s service.RentalService) *RentalController {
	return &RentalController{svc: s}
}

func (h *RentalController) CreateRental(c echo.Context) error {
	// Take UserID from JWT or Body
	uid, _ := strconv.Atoi(c.FormValue("user_id"))
	eid, _ := strconv.Atoi(c.FormValue("equipment_id"))

	res, err := h.svc.RentEquipment(uid, eid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Rental attempt is successful.",
		"data":    res,
	})
}
