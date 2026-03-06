package controller

import (
	"fmt"
	"net/http"
	"sport-rental/models"
	"sport-rental/service"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.UserService
}

// Constructor
func NewUserController(s service.UserService) *UserController {
	return &UserController{userService: s}
}

// Register - POST /register handler
func (h *UserController) Register(c echo.Context) error {
	var user models.User

	// Bind JSON to struct
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Format data are invalid."})
	}

	// Call service
	res, err := h.userService.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "User registration failed."})
	}

	// Return success response
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "User registration success.",
		"data":    res,
	})
}

// Login - POST /login handler
func (h *UserController) Login(c echo.Context) error {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON to struct
	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Input invalid."})
	}

	// Call service
	user, err := h.userService.Login(loginData.Email, loginData.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
	}

	// Return success response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"data":    user,
	})

}

// Midtrans Notification
func (h *UserController) HandleMTNotifs(c echo.Context) error {
	var notification map[string]interface{}
	if err := c.Bind(&notification); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid notification."})
	}

	transactionStatus := notification["transaction_status"].(string)
	orderID := notification["order_id"].(string)
	grossAmountstr := notification["gross_amount"].(string)

	// Convert gross amount to float64
	var grossAmount float64
	fmt.Sscanf(grossAmountstr, "%f", &grossAmount)

	if transactionStatus == "settlement" {
		// Take user ID from order ID
		parts := strings.Split(orderID, "-")
		if len(parts) > 1 {
			userID, _ := strconv.Atoi(parts[1])

			// Update balance in database
			err := h.userService.AddBalance(userID, grossAmount)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "Failed to update user's balance."})
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Notification has been handled."})
}
