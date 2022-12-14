package productdto

type ProductRequest struct {
	Name   string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Desc   string `json:"desc" gorm:"type:text" form:"desc"`
	Price  int    `json:"price" form:"price" gorm:"type: int"`
	Image  string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Stock  int    `json:"stock" form:"stock" gorm:"type: int"`
	UserID int    `json:"user_id" gorm:"type: int"`
}

type UpdateProductRequest struct {
	Name  string `json:"name" form:"name"`
	Desc  string `json:"desc"  form:"desc"`
	Price int    `json:"price" form:"price"`
	Image string `json:"image" form:"image"`
	Users int    `json:"-"`
	Stock int    `json:"stock"`
}
