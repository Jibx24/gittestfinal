package entity
import (
	"gorm.io/gorm"
)
type Member struct {
	gorm.Model
	PhoneNumber	string	`valid:"required,stringlength(10|10)~Phone Number length is not 10 digits."`
	Password	string	`valid:"required"`
	Url	string	`valid:"required,url~Url is invalid"`
}