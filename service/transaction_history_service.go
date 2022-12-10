package service

import (
	"context"
	"errors"
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

	// begin tx
	tx, err := s.transactionHistoryRepository.BeginTx(ctx)
	if err != nil {
		return entity.TransactionHistory{}, err
	}

	userRepositoryTx := repository.NewUserRepository(tx)
	categoryRepositoryTx := repository.NewCategoryRepository(tx)
	productRepositoryTx := repository.NewProductRepo(tx)

	// fetch product
	product, err := productRepositoryTx.GetProductByID(ctx, transactionHistoryEntity.ProductID)
	if err != nil {
		s.transactionHistoryRepository.RollbackTx(ctx, tx)
	}

	// check if product exists
	if (product == entity.Product{}) {
		if err := s.transactionHistoryRepository.CommitTx(ctx, tx); err != nil {
			return entity.TransactionHistory{}, err
		}
		return entity.TransactionHistory{}, errors.New("product does not exist")
	}
	// check if stock is less than or equal to req quantity
	if product.Stock < transactionHistoryEntity.Quantity {
		if err := s.transactionHistoryRepository.CommitTx(ctx, tx); err != nil {
			return entity.TransactionHistory{}, err
		}
		return entity.TransactionHistory{}, errors.New("not enough stock")
	}

	// fetch user
	balance, err := userRepositoryTx.GetUserBalance(ctx, transactionHistoryEntity.UserID)
	if err != nil {
		s.transactionHistoryRepository.RollbackTx(ctx, tx)
	}

	// check user balance
	if balance < (product.Price * transactionHistoryEntity.Quantity) {
		if err := s.transactionHistoryRepository.CommitTx(ctx, tx); err != nil {
			return entity.TransactionHistory{}, err
		}
		return entity.TransactionHistory{}, errors.New("not enough money")
	}

	// reduce stock from product
	err = productRepositoryTx.ReduceProductStock(ctx, transactionHistory.ProductID, transactionHistory.Quantity)
	if err != nil {
		s.transactionHistoryRepository.RollbackTx(ctx, tx)
		return entity.TransactionHistory{}, err
	}

	// reduce user balance
	err = userRepositoryTx.ReduceUserBalance(ctx, transactionHistoryEntity.UserID, product.Price*transactionHistoryEntity.Quantity)
	if err != nil {
		s.transactionHistoryRepository.RollbackTx(ctx, tx)
		return entity.TransactionHistory{}, err
	}

	// increase category sold amount
	err = categoryRepositoryTx.IncreaseSoldProductAmount(ctx, product.CategoryID, transactionHistoryEntity.Quantity)
	if err != nil {
		s.transactionHistoryRepository.RollbackTx(ctx, tx)
		return entity.TransactionHistory{}, err
	}

	transactionHistoryEntity.TotalPrice = product.Price * transactionHistoryEntity.Quantity
	res, err := s.transactionHistoryRepository.CreateTransactionHistory(ctx, transactionHistoryEntity)
	if err != nil {
		s.transactionHistoryRepository.RollbackTx(ctx, tx)
		return transactionHistoryEntity, err
	}

	err = s.transactionHistoryRepository.CommitTx(ctx, tx)
	if err != nil {
		return entity.TransactionHistory{}, err
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
