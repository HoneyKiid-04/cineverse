package service

import (
	"cineverse/internal/model"
	"cineverse/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Request structs
type UpdateProfileInput struct {
	Username  *string `json:"username"`
	Email     *string `json:"email"`
	Bio       *string `json:"bio"`
	AvatarURL *string `json:"avatar_url"`
}

type ChangePasswordInput struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=6"`
}

// Response structs
type ProfileResponse struct {
	User *model.User `json:"user"`
}

type UserProfileService struct {
	userRepo *repository.UserRepository
}

func NewUserProfileService(userRepo *repository.UserRepository) *UserProfileService {
	return &UserProfileService{
		userRepo: userRepo,
	}
}

func (s *UserProfileService) GetProfile(userID uint) (*ProfileResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &ProfileResponse{User: user}, nil
}

func (s *UserProfileService) UpdateProfile(userID uint, input UpdateProfileInput) (*ProfileResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if input.Username != nil {
		// Check if new username is already taken
		if *input.Username != user.Username {
			if _, err := s.userRepo.FindByUsername(*input.Username); err == nil {
				return nil, errors.New("username already taken")
			}
			user.Username = *input.Username
		}
	}

	if input.Email != nil {
		// Check if new email is already taken
		if *input.Email != user.Email {
			if _, err := s.userRepo.FindByEmail(*input.Email); err == nil {
				return nil, errors.New("email already taken")
			}
			user.Email = *input.Email
		}
	}

	if input.Bio != nil {
		user.Bio = *input.Bio
	}

	if input.AvatarURL != nil {
		user.AvatarURL = *input.AvatarURL
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return &ProfileResponse{User: user}, nil
}

func (s *UserProfileService) ChangePassword(userID uint, input ChangePasswordInput) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	// Verify current password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.CurrentPassword)); err != nil {
		return errors.New("current password is incorrect")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.userRepo.Update(user)
}

func (s *UserProfileService) DeleteProfile(userID uint) error {
	return s.userRepo.Delete(userID)
}
