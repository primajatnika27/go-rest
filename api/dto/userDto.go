package dto

type UserRequestModel struct {
	ID       string `json:"id"`
	Name     string `json:"name" `
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Company  string `json:"company"`
	Address  string `json:"address"`
	Active   string `json:"active"`
}
