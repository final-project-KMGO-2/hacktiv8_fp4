package entity

type Product struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	Title      string    `json:"title"`
	Price      uint64    `json:"price"`
	Stock      uint64    `json:"stock"`
	CategoryID uint64    `gorm:"foreignKey" json:"category_id"`
	Category   *Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category,omitempty"`
	BaseModel
}

type ProductCreate struct {
	Title      string `json:"title" binding:"required"`
	Price      uint64 `json:"price" binding:"required,min=0,max=50000000"`
	Stock      uint64 `json:"stock" binding:"required,min=5"`
	CategoryID uint64 `json:"category_id" binding:"required"`
}

type ProductUpdate struct {
	ID         uint64 `json:"id"`
	Title      string `json:"title" binding:"required"`
	Price      uint64 `json:"price" binding:"required,min=0,max=50000000"`
	Stock      uint64 `json:"stock" binding:"required,min=5"`
	CategoryID uint64 `json:"category_id" binding:"required"`
}
