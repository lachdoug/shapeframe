package utils

func IsValidPassword(password string, hash string) (is bool) {
	if hash == HashPassword(password) {
		is = true
	}
	return
}
