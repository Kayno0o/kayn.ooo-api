package entity

var Roles = map[string][]string{
	"ROLE_ADMIN": {"ROLE_ADMIN", "ROLE_USER"},
	"ROLE_USER":  {"ROLE_USER"},
}

type UserRegisterForm struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"`
	Role     string `json:"-" gorm:"default:ROLE_USER;not null"`
}

func (u *User) HasRole(role string) bool {
	for _, r := range Roles[u.Role] {
		if r == role {
			return true
		}
	}
	return false
}
