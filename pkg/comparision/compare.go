package comparision

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

func CheckExpirationTime(exp time.Time) bool {
	return exp.After(time.Now())
}
func ComparePassword(new,old []byte) error {
	err := bcrypt.CompareHashAndPassword(new, old)
	return  err
}