package dao

type User struct {
	UID      int64  `json:"uid "gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	RealName string `json:"real_name" gorm:"column:real_name"`
	Role     int    `json:"role" gorm:"column:role"`
	Phone    string `json:"phone" gorm:"column:phone"`
	EMail    string `json:"e_mail" gorm:"column:e_mail"`
	College  string `json:"college" gorm:"column:college"`
}
