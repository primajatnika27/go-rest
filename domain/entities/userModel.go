package entities

type UserResponseModel struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname; omitempty"`
	Username string `json:"username; omitempty"`
	Password string `json:"password"`
	Email    string `json:"email; omitempty"`
	Phone    string `json:"phone; omitempty"`
	Company  string `json:"company; omitempty"`
	Address  string `json:"address; omitempty"`
	Active   string `json:"active"`
}
