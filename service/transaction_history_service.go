package service

import (
	"context"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/repository"

	"github.com/mashingan/smapping"
)

type TransactionHistoryService interface {
	CreateTransactionHistory(ctx context.Context, transactionHistory entity.TransactionHistoryCreate) (entity.TransactionHistory, error)
	GetAllTransactionHistory(ctx context.Context) ([]entity.TransactionHistory, error)
	GetTransactionHistoryByUserID(ctx context.Context, userID uint64) ([]entity.TransactionHistory, error)
}

type transactionHistoryService struct {
	transactionHistoryRepository repository.TransactionHistoryRepository
}

func NewTransactionHistoryService(thr repository.TransactionHistoryRepository) TransactionHistoryService {
	return &transactionHistoryService{
		transactionHistoryRepository: thr,
	}
}

func (s *transactionHistoryService) CreateTransactionHistory(ctx context.Context, transactionHistory entity.TransactionHistoryCreate) (entity.TransactionHistory, error) {
	transactionHistoryEntity := entity.TransactionHistory{}
	err := smapping.FillStruct(&transactionHistoryEntity, smapping.MapFields(&transactionHistory))
	if err != nil {
		return transactionHistoryEntity, err
	}

	res, err := s.transactionHistoryRepository.CreateTransactionHistory(ctx, transactionHistoryEntity)
	if err != nil {
		return transactionHistoryEntity, err
	}
	return res, nil
}

func (s *transactionHistoryService) GetAllTransactionHistory(ctx context.Context) ([]entity.TransactionHistory, error) {
	res, err := s.transactionHistoryRepository.GetAllTransactionHistory(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *transactionHistoryService) GetTransactionHistoryByUserID(ctx context.Context, userID uint64) ([]entity.TransactionHistory, error) {
	res, err := s.transactionHistoryRepository.GetTransactionHistoryByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
