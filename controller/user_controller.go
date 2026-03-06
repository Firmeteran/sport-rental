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
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "Invalid notification."})
	}

	// Handle string using type assertion
	transactionStatus, _ := notification["transaction_status"].(string)
	orderID, _ := notification["order_id"].(string)

	// Midtrans notification test
	if strings.HasPrefix(orderID, "payment_notif_test") {
		fmt.Println("DEBUG: Menerima notifikasi tes dari Midtrans.")
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "Test notification success"})
	}

	if transactionStatus == "settlement" {
		var grossAmount float64
		switch v := notification["gross_amount"].(type) {
		case string:
			fmt.Sscanf(v, "%f", &grossAmount)
		case float64:
			grossAmount = v
		}

		// Update balance in database
		parts := strings.Split(orderID, "-")
		if len(parts) >= 2 {
			userID, err := strconv.Atoi(parts[1])
			if err == nil || userID > 0 {
				_ = h.userService.AddBalance(userID, grossAmount)
				fmt.Printf("DEBUG: Berhasil tambah saldo %.2f ke User ID %d\n", grossAmount, userID)
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Notification has been handled."})
}
