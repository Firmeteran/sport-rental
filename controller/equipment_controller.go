package controller

import (
	"net/http"
	"sport-rental/models"
	"sport-rental/service"

	"github.com/labstack/echo/v4"
)

type EquipmentController struct {
	svc service.EquipmentService
}

func NewEquipmentController(s service.EquipmentService) *EquipmentController {
	return &EquipmentController{svc: s}
}

func (h *EquipmentController) Create(c echo.Context) error {
	var eq models.Equipment

	// Bind JSON to struct
	if err := c.Bind(&eq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Input invalid."})
	}

	// Call service
	res, err := h.svc.AddEquipment(eq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}

	// Return success response
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Equipment added successfully.",
		"data":    res,
	})
}

func (h *EquipmentController) GetAll(c echo.Context) error {
	res, err := h.svc.FetchAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to fetch."})
	}

	// Return success response
	return c.JSON(http.StatusCreated, map[string]interface{}{"data": res})
}
