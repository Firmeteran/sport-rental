package service

import (
	"fmt"
	"os"
	"sport-rental/models"
	"sport-rental/repository"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type TopUpService interface {
	CreateTopUp(userID int, amount float64) (models.TopUp, error)
}

type topUpService struct {
	repo repository.TopUpRepo
}

func NewTopUpService(r repository.TopUpRepo) TopUpService {
	return &topUpService{r}
}

func (s *topUpService) CreateTopUp(userID int, amount float64) (models.TopUp, error) {
	// Take server key from enviroment variable
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")

	// Specify the environment (sandbox vs production)
	enviroment := midtrans.Sandbox
	if os.Getenv("MIDTRANS_IS_PRODUCTION") == "true" {
		enviroment = midtrans.Production
	}

	// Midtrans Snap client init dynamically
	var sc snap.Client
	sc.New(serverKey, enviroment)

	// Create unique Order ID
	orderID := fmt.Sprintf("TUP-%d-%d", userID, time.Now().Unix())

	// Create request to Midtrans
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(amount),
		},
	}

	// Request Snap response
	snapResp, err := sc.CreateTransaction(req)
	if err != nil {
		return models.TopUp{}, fmt.Errorf("Midtrans error: %v", err)
	}

	// Save to database
	topup := models.TopUp{
		UserID:  uint(userID),
		Amount:  amount,
		OrderID: orderID,
		SnapURL: snapResp.RedirectURL,
		Status:  "pending",
	}

	return s.repo.Create(topup)
}
