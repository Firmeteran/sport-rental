package service

import (
	"errors"
	"sport-rental/models"
	"sport-rental/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(input models.User) (models.User, error)
	Login(email, password string) (models.User, error)
	AddBalance(userID int, amount float64) error
}

type userService struct {
	userRepo repository.UserRepo
}

// Constructor with dependency injection
func NewUserService(repo repository.UserRepo) UserService {
	return &userService{userRepo: repo}
}

// Register - new user registration
func (s *userService) Register(input models.User) (models.User, error) {
	// Warning password should not be empty
	if input.Password == "" {
		return models.User{}, errors.New("Password cannot be empty.")
	}

	// Hash password before get stored
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	input.Password = string(hashedPassword)

	// Call repo to store the data
	return s.userRepo.Create(input)
}

// Login - email and password verification
func (s *userService) Login(email, password string) (models.User, error) {
	// Search user based on email
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return models.User{}, errors.New("User cannot be found.")
	}

	// Compare password input with the hashed one from database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, errors.New("Wrong password.")
	}

	return user, nil
}

// AddBalance - add deposit balance for user
func (s *userService) AddBalance(userID int, amount float64) error {
	// User validation if it's in the database
	_, err := s.userRepo.GetByID(uint(userID))
	if err != nil {
		return errors.New("User not found.")
	}

	return s.userRepo.AddBalance(userID, amount)
}
