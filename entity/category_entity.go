package entity

type Category struct {
	ID                uint64   `gorm:"primaryKey" json:"id"`
	Type              string   `json:"type"`
	SoldProductAmount uint64   `json:"sold_product_amount"`
	Product           *Product `json:"product,omitempty"`
	BaseModel
}

type CategoryCreate struct {
	Type string `json:"type" binding:"required"`
}

type CategoryPatch struct {
	ID   uint64 `gorm:"primaryKey" json:"id"`
	Type string `json:"type" binding:"required"`
}
