package dao

type User struct {
	ID           int64  `json:"id " 			gorm:"column:id"`
	Name         string `json:"name" 			gorm:"column:name"`
	IdentifyName string `json:"identify_name" 	gorm:"column:identify_name"`
	Role         int    `json:"role" 			gorm:"column:role"`
	Phone        string `json:"phone" 			gorm:"column:phone"`
	EMail        string `json:"e_mail"			gorm:"column:e_mail"`
	College      string `json:"college" 		gorm:"column:college"`
	Password     string `json:"password" 		gorm:"column:password"`
}
