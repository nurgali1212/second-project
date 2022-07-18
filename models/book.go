package models

type Book struct {
	ID     		uint  	 `json:"id" gorm:"primary_key"`
	Title  		string 	 `json:"title"`
	Author 		string   `json:"author"`
	CategoryID  uint     `json:"category_id"`
	Category Category    `json:"category" gorm:"foreignkey:CategoryId"`
}
