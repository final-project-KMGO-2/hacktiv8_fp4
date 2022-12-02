package entity

type Product struct {
	ID         int64  `json:"id" gorm:"primaryKey"`         
	Title      string `json:"title" binding:"required"`      
	Price      uint64  `json:"price" binding:"required,min=0,max=50000000"`      
	Stock      uint  `json:"stock" binding:"required,min=5"`           
	CategoryID int64  `json:"category_id" gorm:"foreignKey"`
}

type ProductCreate struct {
	Title      string `json:"title"`      
	Price      uint64  `json:"price"`      
	Stock      uint  `json:"stock"`      
	CategoryID int64  `json:"category_id" gorm:"foreignKey"`
}

type ProductUpdate struct {
	Title      string `json:"title"`      
	Price      uint64  `json:"price"`      
	Stock      uint  `json:"stock"`      
	CategoryID int64  `json:"category_id" gorm:"foreignKey"`
}



