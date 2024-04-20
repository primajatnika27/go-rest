package orm

type UsersOrm struct {
	ID       int64  `gorm:"column:i_id"`
	Name     string `gorm:"column:n_name"`
	Surname  string `gorm:"column:n_surname"`
	Username string `gorm:"column:n_username"`
	Password string `gorm:"column:n_password"`
	Email    string `gorm:"column:n_email"`
	Phone    string `gorm:"column:n_phone"`
	Company  string `gorm:"column:n_company"`
	Address  string `gorm:"column:n_address"`
	Active   string `gorm:"column:c_active"`
}

func (c *UsersOrm) TableName() string {
	return "tm_users"
}
