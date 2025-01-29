package model

type Product struct {
	ProductId   int     `gorm:"primaryKey;type:int" json:"productId"`
	Name        string  `gorm:"type:varchar(255)" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Type        string  `gorm:"type:varchar(100)" json:"type"`
	CommentNum  int     `gorm:"type:int" json:"commentNum"`
	Price       float64 `gorm:"type:decimal(10,2)" json:"price"`
	IsAddedCart bool    `gorm:"type:tinyint;column:is_addedCart" json:"isAddedCart"`
	Cover       string  `gorm:"varchar(255)" json:"cover"`
	PublishTime string  `gorm:"type:varchar(50)" json:"publishTime"`
	Link        string  `gorm:"type:varchar(255)" json:"link"`
}

type Cart struct {
	UserId    int     `gorm:"type:int;primaryKey" json:"userId"`
	ProductId int     `gorm:"type:int" json:"productId"`
	Name      string  `gorm:"type:varchar(255)" json:"name"`
	Type      string  `gorm:"type:varchar(100)" json:"type"`
	Price     float64 `gorm:"type:decimal(10,2)" json:"price"`
	Cover     string  `gorm:"type:varchar(255)" json:"cover"`
	Link      string  `gorm:"type:varchar(255)" json:"link"`
	Num       int     `gorm:"type:int" json:"num"`
}
