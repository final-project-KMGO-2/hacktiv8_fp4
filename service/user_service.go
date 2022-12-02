package service

import (
	"context"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	CreateUser(ctx context.Context, user entity.UserRegister) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	UpdateUserBalance(ctx context.Context, userID uint64, amount uint64) (uint64, error)
	DeleteUser(ctx context.Context, userID uint64) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepository: ur,
	}
}

func (s *userService) CreateUser(ctx context.Context, user entity.UserRegister) (entity.User, error) {
	createdUser := entity.User{}
	err := smapping.FillStruct(&createdUser, smapping.MapFields(&user))
	if err != nil {
		return createdUser, err
	}

	res, err := s.userRepository.CreateUser(ctx, createdUser)
	if err != nil {
		return createdUser, err
	}
	return res, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	return s.userRepository.GetUserByEmail(ctx, email)
}

func (s *userService) UpdateUserBalance(ctx context.Context, userID uint64, amount uint64) (uint64, error) {
	err := s.userRepository.UpdateUserBalance(ctx, userID, amount)
	if err != nil {
		return 0, err
	}

	res, err := s.userRepository.GetUserBalance(ctx, userID)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (s *userService) DeleteUser(ctx context.Context, userID uint64) error {
	err := s.userRepository.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
