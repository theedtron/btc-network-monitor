package domain

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"  gorm:"unique;not null"`
	Password  string `json:"password"`
	Model
}
