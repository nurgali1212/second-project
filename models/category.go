package models

type Category struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Ganre 		string `json:"ganre"`
}