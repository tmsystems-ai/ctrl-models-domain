package models

// CheckPassword checks if the password is correct
func (u *User) CheckPassword(password string) bool {
	if u.Password == nil {
		return false
	}
	return *u.Password == password
}
