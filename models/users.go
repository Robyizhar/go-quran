package models

type Users struct {
	Id       int64  `gorm:"primarykey" json:"id"`
	Name     string `gorm:"type:varchar(300)" json:"name"`
	Email    string `gorm:"type:varchar(300)" json:"email"`
	Password string `gorm:"type:varchar(300)" json:"password"`
}
