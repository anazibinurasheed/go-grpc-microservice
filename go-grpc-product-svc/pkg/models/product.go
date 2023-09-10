package models

type Product struct {
	Id               int64            `json:"id" gorm:"primaryKey"`
	Name             string           `json:"name"`
	Stock            int64            `json:"stock"`
	Price            int64            `json:"price"`
	StockDecreaseLog StockDecreaseLog `gorm:"foreignKey:ProductRefer"`
}
