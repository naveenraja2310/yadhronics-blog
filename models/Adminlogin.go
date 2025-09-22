package models

type Admin struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (Admin) TableName() string {
	return "admin"
}
