package entity

type TransactionHistory struct {
	ID         uint64   `gorm:"primaryKey" json:"id"`
	ProductID  uint64   `gorm:"foreignKey" json:"product_id"`
	Product    *Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product,omitempty"`
	Quantity   uint64   `json:"quantity"`
	TotalPrice uint64   `json:"total_price"`
	UserID     uint64   `gorm:"foreignKey" json:"user_id"`
	User       *User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	BaseModel
}

type TransactionHistoryCreate struct {
	ProductID uint64 `json:"product_id" binding:"required"`
	Quantity  uint64 `json:"quantity" binding:"required"`
	UserID    uint64 `json:"userID"`
}
