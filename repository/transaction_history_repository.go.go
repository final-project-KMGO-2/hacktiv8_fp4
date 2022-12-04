package repository

import (
	"context"
	"hacktiv8_fp_2/entity"

	"gorm.io/gorm"
)

type TransactionHistoryRepository interface {
	CreateTransactionHistory(ctx context.Context, transactionHistory entity.TransactionHistory) (entity.TransactionHistory, error)
	GetAllTransactionHistory(ctx context.Context) ([]entity.TransactionHistory, error)
	GetTransactionHistoryByUserID(ctx context.Context, userID uint64) ([]entity.TransactionHistory, error)
}

type transactionHistoryConnection struct {
	connection *gorm.DB
}

func NewTransactionHistoryRepository(db *gorm.DB) TransactionHistoryRepository {
	return &transactionHistoryConnection{
		connection: db,
	}
}

func (db *transactionHistoryConnection) CreateTransactionHistory(ctx context.Context, transactionHistory entity.TransactionHistory) (entity.TransactionHistory, error) {
	tx := db.connection.Create(&transactionHistory)
	if tx.Error != nil {
		return entity.TransactionHistory{}, tx.Error
	}
	return transactionHistory, nil
}

func (db *transactionHistoryConnection) GetAllTransactionHistory(ctx context.Context) ([]entity.TransactionHistory, error) {
	var transactionHistoryList []entity.TransactionHistory
	tx := db.connection.Find(&transactionHistoryList)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return transactionHistoryList, nil
}

func (db *transactionHistoryConnection) GetTransactionHistoryByUserID(ctx context.Context, userID uint64) ([]entity.TransactionHistory, error) {
	var transactionHistoryList []entity.TransactionHistory
	tx := db.connection.Where(("user_id = ?"), userID).Find(&transactionHistoryList)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return transactionHistoryList, nil
}
