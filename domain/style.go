package domain

// Style struct is base of style model
type Style struct {
	ID   uint8  `gorm:"primary_key"`
	Name string `json:"name" gorm:"size:100;not null"`
}
