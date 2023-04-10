package entity

type Translation struct {
	Model
	Key string `json:"key" gorm:"not null"`
	Fr  string `json:"fr" gorm:"not null"`
	En  string `json:"en" gorm:"not null"`
}
