package random

func RandomPassword(i int) string {
	// This function generates a random password of length i
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	password := make([]byte, i)
	for i := range password {
		password[i] = text[i%len(text)]
	}

	return string(password)
}
