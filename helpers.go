package models

import "time"

// CheckPassword checks if the password is correct
func (u *User) CheckPassword(password string) bool {
	if u.Password == nil {
		return false
	}
	return *u.Password == password
}

func (o OTPCode) IsValid() bool {
	return time.Now().UTC().Before(o.ExpiresAt.UTC())
}
