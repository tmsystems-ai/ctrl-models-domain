package models

// CheckPassword checks if the password is correct
func (u *User) CheckPassword(password string) bool {
	return u.Password == password
}
